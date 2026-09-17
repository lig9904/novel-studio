package agents

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/errs"
	"github.com/chenhongyang/novel-studio/internal/store"
	"github.com/chenhongyang/novel-studio/internal/tools"
	"github.com/voocel/agentcore"
)

const PlannerOnlyRecoveryContractV1 = "planner-only-recovery.v1"

type PlannerOnlyRecoveryContract struct {
	Version                    string `json:"version"`
	ExecutionID                string `json:"execution_id"`
	GenerationID               string `json:"generation_id"`
	Chapter                    int    `json:"chapter"`
	PlanningContextDigest      string `json:"planning_context_digest"`
	PlannerProtocol            string `json:"planner_protocol"`
	SimulationID               string `json:"simulation_id"`
	SimulationDigest           string `json:"simulation_digest"`
	SimulationCheckpointSeq    int64  `json:"simulation_checkpoint_seq"`
	SimulationCheckpointDigest string `json:"simulation_checkpoint_digest"`
	EvidenceDigest             string `json:"evidence_digest"`
	PartialDigest              string `json:"partial_digest"`
}

type PlannerOnlyTraceEvent struct {
	Sequence      int      `json:"sequence"`
	ExecutionID   string   `json:"execution_id"`
	Stage         string   `json:"stage"`
	Tool          string   `json:"tool,omitempty"`
	ArgumentKeys  []string `json:"argument_keys,omitempty"`
	ArgumentsHash string   `json:"arguments_digest,omitempty"`
	BeforePartial string   `json:"before_partial_digest,omitempty"`
	AfterPartial  string   `json:"after_partial_digest,omitempty"`
	StateDigest   string   `json:"state_digest,omitempty"`
	Result        string   `json:"result"`
	Error         string   `json:"error,omitempty"`
}

type PlannerOnlyTraceRecorder func(PlannerOnlyTraceEvent) error

type PlannerOnlyRecoveryRequest struct {
	Contract       PlannerOnlyRecoveryContract
	Trace          PlannerOnlyTraceRecorder
	MaxTraceEvents int
}

type plannerOnlyTraceEmitter struct {
	executionID string
	record      PlannerOnlyTraceRecorder
	max         int
	sequence    int
	mu          sync.Mutex
	cancel      context.CancelFunc
	fatal       error
}

type PlannerOnlyTraceError struct {
	Code string
}

func (e *PlannerOnlyTraceError) Error() string {
	if e == nil || e.Code == "" {
		return "planner-only trace failed"
	}
	return "planner-only trace failed: " + e.Code
}

func newPlannerOnlyTraceEmitter(request *PlannerOnlyRecoveryRequest) *plannerOnlyTraceEmitter {
	if request == nil || request.Trace == nil {
		return nil
	}
	limit := request.MaxTraceEvents
	if limit <= 0 {
		limit = 128
	}
	if limit > 1024 {
		limit = 1024
	}
	return &plannerOnlyTraceEmitter{executionID: request.Contract.ExecutionID, record: request.Trace, max: limit}
}

func (e *plannerOnlyTraceEmitter) emit(event PlannerOnlyTraceEvent) error {
	if e == nil || e.record == nil {
		return nil
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.fatal != nil {
		return e.fatal
	}
	if e.sequence >= e.max {
		return e.failLocked("TRACE_LIMIT_EXCEEDED")
	}
	e.sequence++
	event.Sequence = e.sequence
	event.ExecutionID = e.executionID
	event.Stage = boundedPlannerOnlyTraceText(event.Stage, 64)
	event.Tool = boundedPlannerOnlyTraceText(event.Tool, 64)
	event.ArgumentKeys = boundedPlannerOnlyTraceKeys(event.ArgumentKeys)
	event.Error = plannerOnlyTraceErrorCode(event.Error)
	switch event.Result {
	case "STARTED", "PASS", "FAIL", "NOT_EXECUTED":
	default:
		event.Result = "UNKNOWN"
	}
	raw, err := json.Marshal(event)
	if err != nil || len(raw) > 4096 {
		return e.failLocked("TRACE_EVENT_INVALID")
	}
	if err := e.record(event); err != nil {
		return e.failLocked("TRACE_SINK_FAILED")
	}
	return nil
}

func (e *plannerOnlyTraceEmitter) failLocked(code string) error {
	e.fatal = &PlannerOnlyTraceError{Code: code}
	if e.cancel != nil {
		e.cancel()
	}
	return e.fatal
}

func (e *plannerOnlyTraceEmitter) terminalError() error {
	if e == nil {
		return nil
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.fatal
}

func boundedPlannerOnlyTraceText(value string, limit int) string {
	value = strings.TrimSpace(value)
	runes := []rune(value)
	if len(runes) > limit {
		value = string(runes[:limit])
	}
	return value
}

func boundedPlannerOnlyTraceKeys(values []string) []string {
	if len(values) > 32 {
		values = values[:32]
	}
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, boundedPlannerOnlyTraceText(value, 64))
	}
	return out
}

func plannerOnlyTraceErrorCode(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	if value == "" {
		return ""
	}
	switch {
	case strings.Contains(value, "context canceled"):
		return "CONTEXT_CANCELED"
	case strings.Contains(value, "deadline") || strings.Contains(value, "timeout"):
		return "DEADLINE_EXCEEDED"
	case strings.Contains(value, "grounding"):
		return "GROUNDING_REJECTED"
	case strings.Contains(value, "precondition") || strings.Contains(value, "前置"):
		return "PRECONDITION_FAILED"
	default:
		return "EXECUTION_FAILED"
	}
}

type plannerOnlyValidatedState struct {
	simulation         *domain.ChapterWorldSimulation
	checkpoint         *domain.Checkpoint
	characterEvidence  *domain.CharacterAgentEvidenceBundle
	activationEvidence *domain.CharacterActivationChapterEvidence
	contract           PlannerOnlyRecoveryContract
}

func PreparePlannerOnlyRecoveryContract(
	st *store.Store,
	executionID string,
	generationID string,
	chapter int,
	planningContextDigest string,
	plannerProtocol string,
) (PlannerOnlyRecoveryContract, error) {
	seed := PlannerOnlyRecoveryContract{
		Version:               PlannerOnlyRecoveryContractV1,
		ExecutionID:           strings.TrimSpace(executionID),
		GenerationID:          strings.TrimSpace(generationID),
		Chapter:               chapter,
		PlanningContextDigest: strings.TrimSpace(planningContextDigest),
		PlannerProtocol:       strings.TrimSpace(plannerProtocol),
	}
	owner := plannerOnlyLockOwner("planner-only-prepare", seed.ExecutionID)
	if err := st.Runtime.AcquirePipelineExecution(domain.PipelineExecutionLock{
		Mode: domain.PipelineExecutionProjectAll, TargetChapter: chapter, Owner: owner,
		ExpiresAt: time.Now().UTC().Add(10 * time.Minute),
	}); err != nil {
		return PlannerOnlyRecoveryContract{}, fmt.Errorf("acquire planner-only preparation lock: %w", err)
	}
	state, err := inspectPlannerOnlyRecoveryState(st, seed, false)
	releaseErr := st.Runtime.ReleasePipelineExecution(owner)
	if releaseErr != nil {
		releaseErr = fmt.Errorf("release planner-only preparation lock: %w", releaseErr)
	}
	if err := errors.Join(err, releaseErr); err != nil {
		return PlannerOnlyRecoveryContract{}, err
	}
	return state.contract, nil
}

func inspectPlannerOnlyRecoveryState(
	st *store.Store,
	expected PlannerOnlyRecoveryContract,
	compare bool,
) (*plannerOnlyValidatedState, error) {
	if st == nil || expected.Version != PlannerOnlyRecoveryContractV1 ||
		expected.Chapter <= 0 || !boundedPlannerOnlyID(expected.ExecutionID) ||
		strings.TrimSpace(expected.GenerationID) == "" ||
		strings.TrimSpace(expected.PlanningContextDigest) == "" ||
		strings.TrimSpace(expected.PlannerProtocol) == "" {
		return nil, fmt.Errorf("invalid planner-only recovery contract: %w", errs.ErrToolArgs)
	}
	contextToken, err := domain.ProjectedPlanningContextSourceTokenV2(expected.PlanningContextDigest)
	if err != nil {
		return nil, fmt.Errorf("planner-only planning context digest: %w", err)
	}
	cp, err := tools.CurrentChapterWorldSimulationCheckpoint(st, expected.Chapter)
	if err != nil {
		return nil, fmt.Errorf("planner-only simulation checkpoint: %w", err)
	}
	if cp == nil {
		return nil, fmt.Errorf("planner-only requires an existing simulation checkpoint: %w", errs.ErrToolPrecondition)
	}
	_, ready, gaps := tools.ChapterWorldSimulationStatus(st, expected.Chapter)
	if !ready {
		return nil, fmt.Errorf("planner-only simulation is not ready: %s: %w", strings.Join(gaps, "; "), errs.ErrToolPrecondition)
	}
	simulation, err := st.LoadChapterWorldSimulation(expected.Chapter)
	if err != nil || simulation == nil {
		return nil, fmt.Errorf("planner-only load simulation: %w", errorsOrPrecondition(err))
	}
	if simulation.Chapter != expected.Chapter || simulation.GenerationID != expected.GenerationID {
		return nil, fmt.Errorf("planner-only simulation generation/chapter mismatch: %w", errs.ErrToolPrecondition)
	}
	if !exactProjectAllSourceToken(simulation.Sources, contextToken) {
		return nil, fmt.Errorf("planner-only simulation lacks exact projected context %s: %w", contextToken, errs.ErrToolPrecondition)
	}
	simulationDigest := plannerOnlyDigest(*simulation)
	if simulationDigest == "" {
		return nil, fmt.Errorf("planner-only simulation digest failed: %w", errs.ErrToolPrecondition)
	}

	var characterEvidence *domain.CharacterAgentEvidenceBundle
	var activationEvidence *domain.CharacterActivationChapterEvidence
	evidenceDigest := ""
	if simulation.CharacterActivation != nil {
		activationEvidence, err = st.LoadCharacterActivationChapterEvidence(simulation.GenerationID, expected.Chapter)
		if err == nil && activationEvidence == nil {
			err = fmt.Errorf("whole-chapter activation evidence is missing")
		}
		if err == nil {
			err = domain.ValidateCharacterActivationSimulation(*simulation, *activationEvidence)
		}
		if err == nil {
			evidenceDigest = plannerOnlyDigest(*activationEvidence)
		}
	} else {
		characterEvidence, err = loadCharacterAgentEvidence(st, *simulation)
		if err == nil && characterEvidence == nil {
			err = fmt.Errorf("character-agent evidence is missing")
		}
		if err == nil {
			evidenceDigest = plannerOnlyDigest(characterEvidence)
		}
	}
	if err != nil || evidenceDigest == "" {
		return nil, fmt.Errorf("planner-only character/readiness evidence: %w", errorsOrPrecondition(err))
	}
	partialDigest, err := tools.CurrentChapterPlanPartialBindingDigest(st, expected.Chapter, simulation)
	if err != nil {
		return nil, fmt.Errorf("planner-only partial binding: %w", err)
	}
	if formal, err := st.Drafts.LoadChapterPlan(expected.Chapter); err != nil {
		return nil, err
	} else if formal != nil {
		return nil, fmt.Errorf("planner-only recovery requires an unfinalized partial, but a formal plan exists: %w", errs.ErrToolPrecondition)
	}
	actual := expected
	actual.SimulationID = simulation.SimulationID
	actual.SimulationDigest = simulationDigest
	actual.SimulationCheckpointSeq = cp.Seq
	actual.SimulationCheckpointDigest = cp.Digest
	actual.EvidenceDigest = evidenceDigest
	actual.PartialDigest = partialDigest
	if compare && actual != expected {
		return nil, fmt.Errorf("planner-only recovery contract no longer matches frozen sources: %w", errs.ErrToolPrecondition)
	}
	return &plannerOnlyValidatedState{
		simulation: simulation, checkpoint: cp, characterEvidence: characterEvidence,
		activationEvidence: activationEvidence, contract: actual,
	}, nil
}

func boundedPlannerOnlyID(value string) bool {
	value = strings.TrimSpace(value)
	if value == "" || len(value) > 128 {
		return false
	}
	for _, r := range value {
		if !(r == '-' || r == '_' || r == '.' || r >= '0' && r <= '9' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}

func plannerOnlyLockOwner(prefix, executionID string) string {
	return prefix + "-" + executionID + "-" + rand.Text()
}

func errorsOrPrecondition(err error) error {
	if err != nil {
		return err
	}
	return errs.ErrToolPrecondition
}

func plannerOnlyDigest(value any) string {
	digest, err := domain.DeterministicPlanningHash(value)
	if err != nil {
		return ""
	}
	return "sha256:" + digest
}

type plannerOnlyTracedTool struct {
	tool  agentcore.Tool
	store *store.Store
	emit  *plannerOnlyTraceEmitter
}

type plannerOnlyGuardedModel struct {
	base  agentcore.ChatModel
	trace *plannerOnlyTraceEmitter
}

func (m *plannerOnlyGuardedModel) Generate(ctx context.Context, messages []agentcore.Message, specs []agentcore.ToolSpec, opts ...agentcore.CallOption) (*agentcore.LLMResponse, error) {
	if err := m.trace.terminalError(); err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return m.base.Generate(ctx, messages, specs, opts...)
}

func (m *plannerOnlyGuardedModel) GenerateStream(ctx context.Context, messages []agentcore.Message, specs []agentcore.ToolSpec, opts ...agentcore.CallOption) (<-chan agentcore.StreamEvent, error) {
	if err := m.trace.terminalError(); err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return m.base.GenerateStream(ctx, messages, specs, opts...)
}

func (m *plannerOnlyGuardedModel) SupportsTools() bool { return m.base.SupportsTools() }
func (m *plannerOnlyGuardedModel) ProviderName() string {
	if named, ok := m.base.(agentcore.ProviderNamer); ok {
		return named.ProviderName()
	}
	return ""
}
func (m *plannerOnlyGuardedModel) ModelName() string {
	if named, ok := m.base.(agentcore.ModelNamer); ok {
		return named.ModelName()
	}
	return ""
}

func (t *plannerOnlyTracedTool) Name() string           { return t.tool.Name() }
func (t *plannerOnlyTracedTool) Description() string    { return t.tool.Description() }
func (t *plannerOnlyTracedTool) Schema() map[string]any { return t.tool.Schema() }
func (t *plannerOnlyTracedTool) Execute(ctx context.Context, args json.RawMessage) (result json.RawMessage, returnErr error) {
	keys := []string{}
	var decoded map[string]any
	if json.Unmarshal(args, &decoded) == nil {
		for key := range decoded {
			keys = append(keys, key)
		}
		slices.Sort(keys)
	}
	before := plannerOnlyPartialDigest(t.store, decoded)
	if err := t.emit.emit(PlannerOnlyTraceEvent{Stage: "tool_start", Tool: t.Name(), ArgumentKeys: keys, ArgumentsHash: plannerOnlyDigest(json.RawMessage(args)), BeforePartial: before, Result: "STARTED"}); err != nil {
		return nil, err
	}
	defer func() {
		outcome := "PASS"
		errorText := ""
		if returnErr != nil {
			outcome = "FAIL"
			errorText = returnErr.Error()
		}
		if err := t.emit.emit(PlannerOnlyTraceEvent{Stage: "tool_end", Tool: t.Name(), AfterPartial: plannerOnlyPartialDigest(t.store, decoded), Result: outcome, Error: errorText}); err != nil && returnErr == nil {
			returnErr = err
		}
	}()
	return t.tool.Execute(ctx, args)
}

func plannerOnlyPartialDigest(st *store.Store, args map[string]any) string {
	if st == nil {
		return ""
	}
	chapter := 0
	if value, ok := args["chapter"].(float64); ok {
		chapter = int(value)
	}
	if chapter <= 0 {
		return ""
	}
	partial, err := st.Drafts.LoadChapterPlanPartial(chapter)
	if err != nil || partial == nil {
		return ""
	}
	return plannerOnlyDigest(partial)
}
