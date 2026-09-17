package agents

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/chenhongyang/novel-studio/assets"
	"github.com/chenhongyang/novel-studio/internal/bootstrap"
	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/rag"
	"github.com/chenhongyang/novel-studio/internal/store"
	"github.com/chenhongyang/novel-studio/internal/tools"
	"github.com/voocel/agentcore"
)

type plannerOnlyBlockingModel struct {
	started chan struct{}
	release chan struct{}
	calls   atomic.Int32
}

func (m *plannerOnlyBlockingModel) Generate(context.Context, []agentcore.Message, []agentcore.ToolSpec, ...agentcore.CallOption) (*agentcore.LLMResponse, error) {
	if m.calls.Add(1) == 1 {
		close(m.started)
	}
	<-m.release
	return nil, errors.New("synthetic blocked Planner released")
}

func (m *plannerOnlyBlockingModel) GenerateStream(ctx context.Context, messages []agentcore.Message, specs []agentcore.ToolSpec, opts ...agentcore.CallOption) (<-chan agentcore.StreamEvent, error) {
	response, err := m.Generate(ctx, messages, specs, opts...)
	if err != nil {
		return nil, err
	}
	stream := make(chan agentcore.StreamEvent, 1)
	stream <- agentcore.StreamEvent{Type: agentcore.StreamEventDone, Message: response.Message}
	close(stream)
	return stream, nil
}

func (*plannerOnlyBlockingModel) SupportsTools() bool { return true }

type plannerOnlyFixture struct {
	store         *store.Store
	cfg           bootstrap.Config
	bundle        assets.Bundle
	boundary      ProjectedArcBoundary
	generationID  string
	contextDigest string
	protocol      string
	contract      PlannerOnlyRecoveryContract
}

func newPlannerOnlyFixture(t *testing.T) plannerOnlyFixture {
	t.Helper()
	st, cfg, boundary := activationV3RuntimeFixture(t)
	const generationID = "pg2_planner_only_fixture"
	boundary.Volume, boundary.Arc = 1, 1
	boundary.Title, boundary.Goal = "受控恢复弧", "只复用冻结世界结果"
	contextState := domain.ProjectedPlanningContextV2{
		Version: domain.ProjectedPlanningContextV2Version, GenerationID: generationID,
		NextChapter: 1, ThroughChapter: 0,
		StateRoot:       "sha256:" + strings.Repeat("a", 64),
		CumulativeState: []domain.ProjectedPlanningStateFactV2{}, RecentTransitions: []domain.ProjectedPlanningTransitionV2{},
		OpenObligations: []domain.ProjectedPlanningObligationV2{},
	}
	var err error
	contextState.ContextDigest, err = domain.ComputeProjectedPlanningContextV2Digest(contextState)
	if err != nil {
		t.Fatal(err)
	}
	if err := domain.ValidateProjectedPlanningContextV2(contextState); err != nil {
		t.Fatal(err)
	}
	raw, _ := json.MarshalIndent(contextState, "", "  ")
	if err := os.MkdirAll(filepath.Join(st.Dir(), "meta"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(st.Dir(), "meta", "project_all_state.json"), append(raw, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
	contextToken, err := domain.ProjectedPlanningContextSourceTokenV2(contextState.ContextDigest)
	if err != nil {
		t.Fatal(err)
	}
	activationModel := &activationV3RuntimeModel{}
	activationModels := &bootstrap.ModelSet{Default: bootstrap.NewSwappableModel("test", "planner-only-setup", activationModel)}
	const setupOwner = "planner-only-fixture-setup"
	if err := st.Runtime.AcquirePipelineExecution(domain.PipelineExecutionLock{
		Mode: domain.PipelineExecutionProjectAll, TargetChapter: 1, Owner: setupOwner,
		ExpiresAt: time.Now().UTC().Add(time.Hour),
	}); err != nil {
		t.Fatal(err)
	}
	proof, err := runCharacterActivationChapter(context.Background(), cfg, st, activationModels, generationID, 1, boundary, contextState, nil, 4)
	if err != nil {
		t.Fatal(err)
	}
	bundle := assets.Load(cfg.Style)
	resolvedStyle := bundle.ResolveStyle(cfg.Style)
	contextTool := tools.NewContextTool(st, bundle.References, resolvedStyle.ID).WithConfiguredStyle(resolvedStyle.Body)
	worldContextArgs, _ := json.Marshal(map[string]any{"chapter": 1, "profile": "world_simulation"})
	if _, err := contextTool.Execute(context.Background(), worldContextArgs); err != nil {
		t.Fatalf("world simulation context: %v", err)
	}
	if _, _, err := tools.PublishCharacterActivationSimulation(context.Background(), st, generationID, 1, []string{contextToken, domain.PlanGroundingPolicyV1}); err != nil {
		t.Fatal(err)
	}
	if err := domain.ValidateCharacterActivationChapterEvidence(*proof); err != nil {
		t.Fatal(err)
	}
	chunk := rag.NormalizeChunk(domain.RAGChunk{
		ID: "craft:planner-only", SourcePath: "meta/writing-techniques/planner-only.md",
		SourceKind: "craft_technique", Facet: "methodology", Summary: "场景因果方法", Text: "用既有选择和后果组织场景。",
	})
	if err := st.RAG.SaveIndexState(domain.RAGIndexState{
		SchemaVersion: domain.CurrentRAGIndexSchemaVersion,
		Config:        domain.RAGIndexConfig{Collection: "local_keyword"}, Chunks: []domain.RAGChunk{chunk}, UpdatedAt: "fixture",
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := tools.EnsureProjectAllCraftReceiptForCurrentContext(st, 1); err != nil {
		t.Fatalf("craft receipt: %v", err)
	}
	contextArgs, _ := json.Marshal(map[string]any{"chapter": 1, "profile": "planning"})
	if _, err := contextTool.Execute(context.Background(), contextArgs); err != nil {
		t.Fatalf("planning context: %v", err)
	}
	structureArgs, _ := json.Marshal(map[string]any{
		"chapter": 1, "title": "受控恢复", "goal": "甲保留既有选择并组织计划",
		"conflict": "只允许使用冻结结果", "hook": "既有后果继续约束下一步",
	})
	if _, err := tools.NewPlanStructureTool(st).Execute(context.Background(), structureArgs); err != nil {
		t.Fatalf("plan structure: %v", err)
	}
	if err := st.Runtime.ReleasePipelineExecution(setupOwner); err != nil {
		t.Fatal(err)
	}
	protocol := ProjectAllPlanningProtocolWithActivationProducer(
		bundle.Prompts.Planner, domain.CharacterAgentDecisionProtocolV2Version,
		boundary.MaxCharacterActivationCycles, boundary.CharacterActivationPolicy,
		cfg.CharacterAgents.FrozenActivationProducer,
	)
	contract, err := PreparePlannerOnlyRecoveryContract(st, "planner-only-test", generationID, 1, contextState.ContextDigest, protocol)
	if err != nil {
		t.Fatalf("prepare planner-only contract: %v", err)
	}
	return plannerOnlyFixture{
		store: st, cfg: cfg, bundle: bundle, boundary: boundary, generationID: generationID,
		contextDigest: contextState.ContextDigest, protocol: protocol, contract: contract,
	}
}

func rewritePlannerOnlySimulation(t *testing.T, fixture plannerOnlyFixture, mutate func(*domain.ChapterWorldSimulation)) {
	t.Helper()
	simulation, err := fixture.store.LoadChapterWorldSimulation(1)
	if err != nil || simulation == nil {
		t.Fatalf("load simulation for mutation: simulation=%#v err=%v", simulation, err)
	}
	mutate(simulation)
	if err := fixture.store.SaveChapterWorldSimulation(*simulation); err != nil {
		t.Fatal(err)
	}
	if _, err := fixture.store.Checkpoints.AppendArtifactLatest(
		domain.ChapterScope(1), "chapter_world_simulation", "meta/chapter_simulations/001.json",
	); err != nil {
		t.Fatal(err)
	}
}

func TestPlannerOnlyRecoveryEntersPlannerWithoutUpstreamFallback(t *testing.T) {
	fixture := newPlannerOnlyFixture(t)
	probe := &outlineAllOperationCaptureModel{response: agentcore.Message{
		Role: agentcore.RoleAssistant, StopReason: agentcore.StopReasonStop,
		Content: []agentcore.ContentBlock{agentcore.TextBlock("synthetic planner stopped without finalizing")},
		Usage:   &agentcore.Usage{Input: 10, Output: 2},
	}}
	models := &bootstrap.ModelSet{Default: bootstrap.NewSwappableModel("test", "planner-only-probe", probe)}
	roles := []string{}
	trace := []PlannerOnlyTraceEvent{}
	request := PlannerOnlyRecoveryRequest{
		Contract: fixture.contract,
		Trace: func(event PlannerOnlyTraceEvent) error {
			trace = append(trace, event)
			return nil
		},
	}
	_, err := runProjectedChapterPlanning(
		context.Background(), fixture.cfg, fixture.bundle, fixture.store.Dir(), 1,
		fixture.contextDigest, domain.CharacterAgentDecisionProtocolV2Version, fixture.boundary,
		&request, models, ProjectedPlanningAccounting{StartCall: func(_ string, role string) error {
			roles = append(roles, role)
			return nil
		}, RecordUsage: func(string, agentcore.AgentMessage) {}},
	)
	if err == nil || (!strings.Contains(err.Error(), "did not finalize") && !strings.Contains(err.Error(), "stop guard")) {
		t.Fatalf("synthetic Planner stop did not reach the production Planner path: %v", err)
	}
	if probe.calls == 0 || len(roles) != probe.calls {
		t.Fatalf("unexpected business stages: planner calls=%d roles=%v", probe.calls, roles)
	}
	for _, role := range roles {
		if role != "project_all_planner" {
			t.Fatalf("planner-only invoked upstream role %q", role)
		}
	}
	for _, event := range trace {
		if strings.Contains(event.Stage, "world") || strings.Contains(event.Stage, "character") || strings.Contains(event.Stage, "readiness") {
			t.Fatalf("planner-only trace entered an upstream stage: %+v", event)
		}
	}
	if len(trace) < 3 || trace[0].Stage != "preflight" || trace[1].Stage != "preflight" {
		t.Fatalf("planner-only trace did not retain preflight boundaries: %+v", trace)
	}
	if _, err := inspectPlannerOnlyRecoveryState(fixture.store, fixture.contract, true); err != nil {
		t.Fatalf("synthetic Planner attempt changed frozen sources or the historical partial: %v", err)
	}
}

func TestPlannerOnlyRecoveryFailsClosedBeforeAnyModelCall(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*testing.T, plannerOnlyFixture, *PlannerOnlyRecoveryContract)
	}{
		{name: "wrong generation", mutate: func(_ *testing.T, _ plannerOnlyFixture, contract *PlannerOnlyRecoveryContract) {
			contract.GenerationID = "pg2_wrong"
		}},
		{name: "wrong chapter", mutate: func(_ *testing.T, _ plannerOnlyFixture, contract *PlannerOnlyRecoveryContract) {
			contract.Chapter = 2
		}},
		{name: "wrong context", mutate: func(_ *testing.T, _ plannerOnlyFixture, contract *PlannerOnlyRecoveryContract) {
			contract.PlanningContextDigest = "sha256:" + strings.Repeat("c", 64)
		}},
		{name: "wrong checkpoint digest", mutate: func(_ *testing.T, _ plannerOnlyFixture, contract *PlannerOnlyRecoveryContract) {
			contract.SimulationCheckpointDigest = "sha256:" + strings.Repeat("f", 64)
		}},
		{name: "wrong partial digest", mutate: func(_ *testing.T, _ plannerOnlyFixture, contract *PlannerOnlyRecoveryContract) {
			contract.PartialDigest = "sha256:" + strings.Repeat("e", 64)
		}},
		{name: "wrong protocol", mutate: func(_ *testing.T, _ plannerOnlyFixture, contract *PlannerOnlyRecoveryContract) {
			contract.PlannerProtocol = "sha256:" + strings.Repeat("d", 64)
		}},
		{name: "missing simulation", mutate: func(t *testing.T, fixture plannerOnlyFixture, _ *PlannerOnlyRecoveryContract) {
			if err := os.Remove(filepath.Join(fixture.store.Dir(), "meta", "chapter_simulations", "001.json")); err != nil {
				t.Fatal(err)
			}
		}},
		{name: "simulation not ready", mutate: func(t *testing.T, fixture plannerOnlyFixture, _ *PlannerOnlyRecoveryContract) {
			rewritePlannerOnlySimulation(t, fixture, func(simulation *domain.ChapterWorldSimulation) {
				simulation.ProtagonistProjection.ChosenDecision = ""
			})
		}},
		{name: "stale projected context source", mutate: func(t *testing.T, fixture plannerOnlyFixture, _ *PlannerOnlyRecoveryContract) {
			rewritePlannerOnlySimulation(t, fixture, func(simulation *domain.ChapterWorldSimulation) {
				kept := simulation.Sources[:0]
				for _, source := range simulation.Sources {
					if !strings.HasPrefix(source, domain.ProjectedPlanningContextSourcePrefix) {
						kept = append(kept, source)
					}
				}
				simulation.Sources = kept
			})
		}},
		{name: "tampered activation evidence", mutate: func(t *testing.T, fixture plannerOnlyFixture, _ *PlannerOnlyRecoveryContract) {
			path := filepath.Join(fixture.store.Dir(), "meta", "character_agents", "activation_sessions", fixture.generationID, "000001", "chapter_evidence.json")
			raw, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var evidence domain.CharacterActivationChapterEvidence
			if err := json.Unmarshal(raw, &evidence); err != nil {
				t.Fatal(err)
			}
			evidence.Session.Phase = "tampered"
			raw, _ = json.MarshalIndent(evidence, "", "  ")
			if err := os.WriteFile(path, append(raw, '\n'), 0o644); err != nil {
				t.Fatal(err)
			}
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fixture := newPlannerOnlyFixture(t)
			contract := fixture.contract
			tc.mutate(t, fixture, &contract)
			probe := &outlineAllOperationCaptureModel{}
			models := &bootstrap.ModelSet{Default: bootstrap.NewSwappableModel("test", "must-not-run", probe)}
			beforePartial, _ := fixture.store.Drafts.LoadChapterPlanPartial(1)
			beforeDigest, _ := domain.DeterministicPlanningHash(beforePartial)
			_, err := runProjectedChapterPlanning(
				context.Background(), fixture.cfg, fixture.bundle, fixture.store.Dir(), 1,
				fixture.contextDigest, domain.CharacterAgentDecisionProtocolV2Version, fixture.boundary,
				&PlannerOnlyRecoveryRequest{Contract: contract}, models,
			)
			if err == nil || probe.calls != 0 {
				t.Fatalf("invalid planner-only source reached a model: calls=%d err=%v", probe.calls, err)
			}
			afterPartial, _ := fixture.store.Drafts.LoadChapterPlanPartial(1)
			afterDigest, _ := domain.DeterministicPlanningHash(afterPartial)
			if beforeDigest != afterDigest {
				t.Fatal("failed preflight rewrote the historical partial")
			}
		})
	}
}

func TestPlannerOnlyRecoveryCancellationAndTraceLimitFailClosed(t *testing.T) {
	for _, mode := range []string{"canceled", "trace-limit"} {
		t.Run(mode, func(t *testing.T) {
			fixture := newPlannerOnlyFixture(t)
			probe := &outlineAllOperationCaptureModel{}
			models := &bootstrap.ModelSet{Default: bootstrap.NewSwappableModel("test", "must-not-run", probe)}
			ctx := context.Background()
			request := PlannerOnlyRecoveryRequest{Contract: fixture.contract}
			if mode == "canceled" {
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				cancel()
			} else {
				request.MaxTraceEvents = 1
				request.Trace = func(PlannerOnlyTraceEvent) error { return nil }
			}
			_, err := runProjectedChapterPlanning(
				ctx, fixture.cfg, fixture.bundle, fixture.store.Dir(), 1,
				fixture.contextDigest, domain.CharacterAgentDecisionProtocolV2Version, fixture.boundary,
				&request, models,
			)
			if err == nil || probe.calls != 0 {
				t.Fatalf("%s reached a model: calls=%d err=%v", mode, probe.calls, err)
			}
		})
	}
}

func TestPlannerOnlyRecoverySameContractIsMutuallyExclusive(t *testing.T) {
	fixture := newPlannerOnlyFixture(t)
	blocking := &plannerOnlyBlockingModel{started: make(chan struct{}), release: make(chan struct{})}
	firstModels := &bootstrap.ModelSet{Default: bootstrap.NewSwappableModel("test", "blocking-planner", blocking)}
	firstDone := make(chan error, 1)
	go func() {
		_, err := runProjectedChapterPlanning(
			context.Background(), fixture.cfg, fixture.bundle, fixture.store.Dir(), 1,
			fixture.contextDigest, domain.CharacterAgentDecisionProtocolV2Version, fixture.boundary,
			&PlannerOnlyRecoveryRequest{Contract: fixture.contract}, firstModels,
		)
		firstDone <- err
	}()
	<-blocking.started

	secondProbe := &outlineAllOperationCaptureModel{}
	secondModels := &bootstrap.ModelSet{Default: bootstrap.NewSwappableModel("test", "must-not-run", secondProbe)}
	_, secondErr := runProjectedChapterPlanning(
		context.Background(), fixture.cfg, fixture.bundle, fixture.store.Dir(), 1,
		fixture.contextDigest, domain.CharacterAgentDecisionProtocolV2Version, fixture.boundary,
		&PlannerOnlyRecoveryRequest{Contract: fixture.contract}, secondModels,
	)
	if secondErr == nil || !strings.Contains(secondErr.Error(), "acquire project-all execution lock") || secondProbe.calls != 0 {
		close(blocking.release)
		t.Fatalf("same recovery contract reentered: second_calls=%d err=%v", secondProbe.calls, secondErr)
	}
	close(blocking.release)
	if firstErr := <-firstDone; firstErr == nil {
		t.Fatal("blocking synthetic Planner unexpectedly succeeded")
	}
}

func TestPlannerOnlyTraceFailureAfterToolStopsFurtherModelCalls(t *testing.T) {
	fixture := newPlannerOnlyFixture(t)
	structureArgs, _ := json.Marshal(map[string]any{
		"chapter": 1, "title": "受控恢复", "goal": "甲继续整理冻结结果",
		"conflict": "不得补跑世界", "hook": "既有后果仍待处理",
	})
	probe := &outlineAllOperationCaptureModel{responses: []agentcore.Message{{
		Role: agentcore.RoleAssistant, StopReason: agentcore.StopReasonToolUse,
		Content: []agentcore.ContentBlock{agentcore.ToolCallBlock(agentcore.ToolCall{
			ID: "trace-failure-tool", Name: "plan_structure", Args: structureArgs,
		})}, Usage: &agentcore.Usage{Input: 10, Output: 2},
	}}}
	models := &bootstrap.ModelSet{Default: bootstrap.NewSwappableModel("test", "trace-failure", probe)}
	request := PlannerOnlyRecoveryRequest{
		Contract: fixture.contract,
		Trace: func(event PlannerOnlyTraceEvent) error {
			if event.Stage == "tool_end" {
				return errors.New("synthetic trace sink unavailable with secret text")
			}
			return nil
		},
	}
	_, err := runProjectedChapterPlanning(
		context.Background(), fixture.cfg, fixture.bundle, fixture.store.Dir(), 1,
		fixture.contextDigest, domain.CharacterAgentDecisionProtocolV2Version, fixture.boundary,
		&request, models,
	)
	var traceErr *PlannerOnlyTraceError
	if !errors.As(err, &traceErr) || traceErr.Code != "TRACE_SINK_FAILED" || probe.calls != 1 {
		t.Fatalf("trace failure did not stop the run after one model call: calls=%d err=%v", probe.calls, err)
	}
}

func TestPlannerOnlyTraceIsSanitizedAndByteBounded(t *testing.T) {
	var captured PlannerOnlyTraceEvent
	emitter := newPlannerOnlyTraceEmitter(&PlannerOnlyRecoveryRequest{
		Contract: PlannerOnlyRecoveryContract{ExecutionID: "trace-sanitize"},
		Trace: func(event PlannerOnlyTraceEvent) error {
			captured = event
			return nil
		},
	})
	keys := make([]string, 80)
	for index := range keys {
		keys[index] = strings.Repeat("k", 100)
	}
	secret := "/Users/private/story: grounding quote must not leave trace"
	if err := emitter.emit(PlannerOnlyTraceEvent{Stage: strings.Repeat("s", 100), Tool: strings.Repeat("t", 100), ArgumentKeys: keys, Result: "FAIL", Error: secret}); err != nil {
		t.Fatal(err)
	}
	raw, _ := json.Marshal(captured)
	if len(raw) > 4096 || len(captured.ArgumentKeys) != 32 || len([]rune(captured.Stage)) > 64 || len([]rune(captured.Tool)) > 64 || captured.Error != "GROUNDING_REJECTED" || strings.Contains(string(raw), "Users/private") || strings.Contains(string(raw), "grounding quote") {
		t.Fatalf("trace was not sanitized and bounded: bytes=%d event=%+v", len(raw), captured)
	}
}
