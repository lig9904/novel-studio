package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

const CharacterReadinessReviewPolicy = "chapter-readiness:arbitrated-events.v1"
const CharacterReadinessReviewPolicyV2 = "chapter-readiness:soft-event-outcome.v2"
const CharacterReadinessReviewedVersion = "character-chapter-readiness.v2"
const CharacterReadinessReviewedVersionV3 = "character-chapter-readiness.v3"
const CharacterReadinessPolicyBindingVersionV1 = "character-readiness-policy-binding.v1"

const (
	CharacterSoftEventOccurred        = "OCCURRED"
	CharacterSoftEventRejected        = "REJECTED_WITH_CONSEQUENCE"
	CharacterSoftEventSuperseded      = "SUPERSEDED_BY_ACTUAL_CHOICE"
	CharacterSoftEventPending         = "DEFERRED"
	CharacterSoftEventHardUnsatisfied = "HARD_CONTRACT_UNSATISFIED"
)

// This is a frozen host planning context, not a character observation. The
// session binds its digest before the first cycle, so a later outline edit
// cannot silently change what an already paid readiness review was judging.
type CharacterReadinessContext struct {
	Version                 string                             `json:"version"`
	GenerationID            string                             `json:"generation_id"`
	Chapter                 int                                `json:"chapter"`
	POVCharacter            string                             `json:"pov_character"`
	ArcLastChapter          int                                `json:"arc_last_chapter"`
	BookLastChapter         int                                `json:"book_last_chapter"`
	TargetWords             int                                `json:"target_words,omitempty"`
	ProjectionContextDigest string                             `json:"projection_context_digest,omitempty"`
	SoftOutline             OutlineEntry                       `json:"soft_outline"`
	HardContracts           []string                           `json:"hard_contracts"`
	Obligations             []ProjectedPlanningObligationV2    `json:"obligations,omitempty"`
	PolicyBinding           *CharacterReadinessPolicyBindingV1 `json:"policy_binding,omitempty"`
	Digest                  string                             `json:"digest"`
}

// CharacterReadinessPolicyBindingV1 freezes the exact execution semantics
// that selected the readiness contract. ProducerPolicies is an inventory, not
// character-visible knowledge; the producer digest remains the executable root.
type CharacterReadinessPolicyBindingV1 struct {
	Version          string   `json:"version"`
	ActivationPolicy string   `json:"activation_policy"`
	ProducerDigest   string   `json:"producer_digest"`
	ProducerPolicies []string `json:"producer_policies"`
	ReadinessPolicy  string   `json:"readiness_policy"`
	SoftEventPolicy  string   `json:"soft_event_policy"`
	Digest           string   `json:"digest"`
}

func FinalizeCharacterReadinessPolicyBindingV1(value CharacterReadinessPolicyBindingV1) (CharacterReadinessPolicyBindingV1, error) {
	if value.Version == "" {
		value.Version = CharacterReadinessPolicyBindingVersionV1
	}
	if value.Version != CharacterReadinessPolicyBindingVersionV1 || !IsCharacterActivationPolicy(value.ActivationPolicy) || value.ReadinessPolicy != CharacterReadinessReviewPolicyV2 || value.SoftEventPolicy != CharacterSoftEventReadinessPolicyV1 {
		return value, fmt.Errorf("character readiness policy binding has invalid semantics")
	}
	if err := validatePlanningV2Digest("character readiness producer", value.ProducerDigest); err != nil {
		return value, err
	}
	value.ProducerPolicies = normalizeV2Strings(value.ProducerPolicies)
	if len(value.ProducerPolicies) == 0 || !planningV2ContainsExactString(value.ProducerPolicies, value.ActivationPolicy) || !planningV2ContainsExactString(value.ProducerPolicies, value.SoftEventPolicy) {
		return value, fmt.Errorf("character readiness policy binding lacks its activation/soft-event inventory")
	}
	value.Digest = ""
	var err error
	value.Digest, err = characterAgentDigest(value)
	return value, err
}

// ValidateCharacterReadinessContextPolicySources prevents a later producer
// inventory from silently upgrading or downgrading an already frozen chapter.
func ValidateCharacterReadinessContextPolicySources(context CharacterReadinessContext, sources []string) error {
	checked, err := FinalizeCharacterReadinessContext(context)
	if err != nil {
		return err
	}
	if checked.Digest != context.Digest {
		return fmt.Errorf("character readiness context digest mismatch")
	}
	if context.Version != CharacterReadinessReviewPolicyV2 {
		if HasCharacterSoftEventReadinessPolicyV1(sources) {
			return fmt.Errorf("soft-event producer cannot execute against a legacy frozen readiness context")
		}
		return nil
	}
	binding := context.PolicyBinding
	if binding == nil {
		return fmt.Errorf("soft-event readiness context lacks its frozen producer policy binding")
	}
	for _, policy := range binding.ProducerPolicies {
		if !planningV2ContainsExactString(sources, policy) {
			return fmt.Errorf("activation stimulus differs from frozen readiness producer policies")
		}
	}
	wantProducer := "character-agent-protocol:" + binding.ProducerDigest
	producerCount := 0
	for _, source := range sources {
		if strings.HasPrefix(source, "character-agent-protocol:") {
			producerCount++
			if source != wantProducer {
				return fmt.Errorf("activation stimulus producer differs from frozen readiness producer")
			}
		}
	}
	if producerCount != 1 || !planningV2ContainsExactString(sources, binding.ActivationPolicy) || !planningV2ContainsExactString(sources, binding.SoftEventPolicy) {
		return fmt.Errorf("activation stimulus lacks the exact frozen readiness execution binding")
	}
	return nil
}

func FinalizeCharacterReadinessContext(value CharacterReadinessContext) (CharacterReadinessContext, error) {
	if value.Version == "" {
		value.Version = CharacterReadinessReviewPolicy
	}
	if (value.Version != CharacterReadinessReviewPolicy && value.Version != CharacterReadinessReviewPolicyV2) || !strings.HasPrefix(value.GenerationID, PlanningGenerationIDPrefix) || value.Chapter < 1 || value.ArcLastChapter < value.Chapter || value.BookLastChapter < value.ArcLastChapter || value.SoftOutline.Chapter != value.Chapter || strings.TrimSpace(value.POVCharacter) == "" || value.TargetWords < 0 {
		return value, fmt.Errorf("chapter readiness context has invalid identity/boundaries")
	}
	if value.Version == CharacterReadinessReviewPolicyV2 {
		if value.PolicyBinding == nil {
			return value, fmt.Errorf("soft-event readiness context requires its frozen policy binding")
		}
		originalDigest := value.PolicyBinding.Digest
		binding, err := FinalizeCharacterReadinessPolicyBindingV1(*value.PolicyBinding)
		if err != nil {
			return value, err
		}
		if originalDigest != binding.Digest {
			return value, fmt.Errorf("character readiness policy binding digest mismatch")
		}
		value.PolicyBinding = &binding
	} else if value.PolicyBinding != nil {
		return value, fmt.Errorf("legacy readiness context cannot claim a producer policy binding")
	}
	if value.ProjectionContextDigest != "" {
		if err := validatePlanningV2Digest("readiness projection context", value.ProjectionContextDigest); err != nil {
			return value, err
		}
	}
	value.HardContracts = normalizeV2Strings(value.HardContracts)
	seen := map[string]bool{}
	for _, obligation := range value.Obligations {
		if strings.TrimSpace(obligation.ID) == "" || seen[obligation.ID] || strings.TrimSpace(obligation.Contract) == "" {
			return value, fmt.Errorf("readiness context has invalid/duplicate obligations")
		}
		seen[obligation.ID] = true
	}
	value.Digest = ""
	var err error
	value.Digest, err = characterAgentDigest(value)
	return value, err
}

type CharacterReadinessAction struct {
	AgentID         string `json:"agent_id"`
	Character       string `json:"character"`
	ProposalDigest  string `json:"proposal_digest"`
	Decision        string `json:"decision"`
	DecisionReason  string `json:"decision_reason,omitempty"`
	IntendedAction  string `json:"intended_action"`
	Outcome         string `json:"outcome"`
	CompletionState string `json:"completion_state"`
	ImmediateResult string `json:"immediate_result"`
	StateAfter      string `json:"state_after"`
}

type CharacterReadinessCycleView struct {
	Index                 int                        `json:"index"`
	CycleDigest           string                     `json:"cycle_digest"`
	ArbitrationDigest     string                     `json:"arbitration_digest"`
	BeforePhysicalRoot    string                     `json:"before_physical_root,omitempty"`
	AfterPhysicalRoot     string                     `json:"after_physical_root,omitempty"`
	StoryTime             *StoryTimeChapterSchedule  `json:"story_time"`
	Actions               []CharacterReadinessAction `json:"actions"`
	HardContractStatus    string                     `json:"hard_contract_status"`
	HardContractConflicts []string                   `json:"hard_contract_conflicts,omitempty"`
}

type CharacterReadinessActorView struct {
	OperationalObservations []CharacterOperationalObservationV1 `json:"operational_observations,omitempty"`
	AgentID                 string                              `json:"agent_id"`
	Character               string                              `json:"character"`
	Location                string                              `json:"location"`
	Resources               []CharacterResourceViewV2           `json:"resources"`
	TaskProgress            []CharacterTaskProgressV2           `json:"task_progress"`
	ReceivedFacts           []CharacterReceivedFactV2           `json:"received_facts"`
}

// Only the final state is included, once. Individual cycles keep exact intent
// and result text but do not repeat every actor's full physical/memory history.
type CharacterReadinessTrace struct {
	Cycles            []CharacterReadinessCycleView `json:"cycles"`
	FinalPhysicalRoot string                        `json:"final_physical_root"`
	Resources         []WorldResourceBalanceV2      `json:"world_resources"`
	Actors            []CharacterReadinessActorView `json:"actors"`
}

func BuildCharacterReadinessTrace(cycles []CharacterActivationCycle) (CharacterReadinessTrace, error) {
	var trace CharacterReadinessTrace
	if len(cycles) == 0 {
		return trace, fmt.Errorf("readiness requires actual cycle evidence")
	}
	for i, cycle := range cycles {
		if err := ValidateCharacterActivationCycle(cycle); err != nil {
			return trace, err
		}
		if cycle.Index != i+1 || (i > 0 && (cycle.PreviousDigest != cycles[i-1].Digest || cycle.GenerationID != cycles[0].GenerationID || cycle.Chapter != cycles[0].Chapter || cycle.BeforePhysicalRoot != cycles[i-1].AfterPhysicalRoot || cycle.StartDay != cycles[i-1].EndDay)) {
			return trace, fmt.Errorf("readiness cycle chain is incomplete")
		}
		receipt := cycle.Evidence.Arbitrations[len(cycle.Evidence.Arbitrations)-1]
		view := CharacterReadinessCycleView{Index: cycle.Index, CycleDigest: cycle.Digest, ArbitrationDigest: receipt.Digest, BeforePhysicalRoot: cycle.BeforePhysicalRoot, AfterPhysicalRoot: cycle.AfterPhysicalRoot, StoryTime: receipt.StoryTime, HardContractStatus: receipt.HardContractStatus, HardContractConflicts: receipt.HardContractConflicts}
		proposals := map[string]CharacterDecisionProposal{}
		for _, proposal := range LatestCharacterCycleProposals(cycle.Evidence) {
			proposals[proposal.Digest] = proposal
		}
		for _, resolution := range receipt.Resolutions {
			proposal, ok := proposals[resolution.ProposalDigest]
			if !ok || proposal.AgentID != resolution.AgentID {
				return trace, fmt.Errorf("readiness action lacks its exact proposal reason")
			}
			view.Actions = append(view.Actions, CharacterReadinessAction{resolution.AgentID, resolution.Character, resolution.ProposalDigest, resolution.Decision, proposal.DecisionReason, resolution.IntendedAction, resolution.Outcome, resolution.CompletionState, resolution.ImmediateResult, resolution.StateAfter})
		}
		trace.Cycles = append(trace.Cycles, view)
	}
	last := cycles[len(cycles)-1]
	receipt := last.Evidence.Arbitrations[len(last.Evidence.Arbitrations)-1]
	state, err := ApplyArbitrationPhysicalStateV2(receipt, last.Evidence.Stimulus, LatestCharacterCycleProposals(last.Evidence)...)
	if err != nil {
		return trace, err
	}
	trace.FinalPhysicalRoot, trace.Resources = last.AfterPhysicalRoot, state.Resources
	for _, actor := range state.Actors {
		resources, err := BuildCharacterResourceViewsV2(state, actor.AgentID)
		if err != nil {
			return trace, err
		}
		operations, err := selectCharacterOperationalObservationsV1(actor.OperationalObservations)
		if err != nil {
			return trace, err
		}
		trace.Actors = append(trace.Actors, CharacterReadinessActorView{AgentID: actor.AgentID, Character: actor.Character, Location: actor.Location, Resources: resources, TaskProgress: actor.TaskProgress, ReceivedFacts: actor.ReceivedFacts, OperationalObservations: operations})
	}
	return trace, nil
}

type CharacterReadinessRequirement struct {
	ID       string `json:"id"`
	Contract string `json:"contract"`
	DueNow   bool   `json:"due_now"`
}

type CharacterReadinessContractCheck struct {
	ContractID   string   `json:"contract_id"`
	Status       string   `json:"status"` // satisfied / preserved / pending / impossible
	EvidenceRefs []string `json:"evidence_refs"`
}

type CharacterReadinessVerdict struct {
	Decision       string                            `json:"decision"`
	Reason         string                            `json:"reason"`
	EvidenceRefs   []string                          `json:"evidence_refs"`
	ContractChecks []CharacterReadinessContractCheck `json:"contract_checks"`
	SoftEvent      *CharacterReadinessSoftEvent      `json:"soft_event,omitempty"`
}

type CharacterReadinessSoftEvent struct {
	Outcome          string   `json:"outcome"`
	ActorRef         string   `json:"actor_ref,omitempty"`
	ProposalRef      string   `json:"proposal_ref,omitempty"`
	CharacterReason  string   `json:"character_reason,omitempty"`
	WorldConsequence string   `json:"world_consequence,omitempty"`
	EvidenceRefs     []string `json:"evidence_refs"`
}

type CharacterReadinessReviewInput struct {
	Policy          string                          `json:"policy"`
	ReviewProtocol  string                          `json:"review_protocol"`
	SessionDigest   string                          `json:"session_digest"`
	Context         CharacterReadinessContext       `json:"chapter_context"`
	Trace           CharacterReadinessTrace         `json:"actual_events"`
	Requirements    []CharacterReadinessRequirement `json:"required_checks"`
	RemainingCycles int                             `json:"remaining_cycles"`
}

type CharacterReadinessReviewAudit struct {
	Input     CharacterReadinessReviewInput     `json:"input"`
	Receipt   CharacterChapterReadiness         `json:"receipt"`
	ModelView *CharacterReadinessModelBindingV1 `json:"model_view,omitempty"`
}

func NewCharacterReadinessReviewInput(context CharacterReadinessContext, session CharacterActivationSession, cycles []CharacterActivationCycle, reviewProtocol string) (CharacterReadinessReviewInput, error) {
	var input CharacterReadinessReviewInput
	checked, err := FinalizeCharacterReadinessContext(context)
	if err != nil {
		return input, err
	}
	if checked.Digest != context.Digest || context.Digest != session.ChapterContextDigest || context.GenerationID != session.GenerationID || context.Chapter != session.Chapter {
		return input, fmt.Errorf("readiness context is not the session's frozen chapter context")
	}
	if err := ValidateCharacterActivationSession(session); err != nil {
		return input, err
	}
	if session.Phase != "assessing" || len(cycles) != len(session.CycleDigests) {
		return input, fmt.Errorf("readiness must assess the exact pending cycle chain")
	}
	for i := range cycles {
		if cycles[i].Digest != session.CycleDigests[i] {
			return input, fmt.Errorf("readiness cycle differs from committed session")
		}
	}
	if err := ValidateCharacterReadinessContextPolicySources(context, cycles[0].Evidence.Stimulus.Sources); err != nil {
		return input, err
	}
	trace, err := BuildCharacterReadinessTrace(cycles)
	if err != nil {
		return input, err
	}
	if context.Version != CharacterReadinessReviewPolicyV2 {
		stripCharacterReadinessV2Trace(&trace)
	}
	if err := validatePlanningV2Digest("readiness protocol", reviewProtocol); err != nil {
		return input, err
	}
	policy := CharacterReadinessReviewPolicy
	if context.Version == CharacterReadinessReviewPolicyV2 {
		policy = CharacterReadinessReviewPolicyV2
	}
	input = CharacterReadinessReviewInput{Policy: policy, ReviewProtocol: reviewProtocol, SessionDigest: session.Digest, Context: context, Trace: trace, RemainingCycles: session.MaxCycles - len(cycles)}
	input.Requirements, err = characterReadinessRequirements(context)
	return input, err
}

func stripCharacterReadinessV2Trace(trace *CharacterReadinessTrace) {
	if trace == nil {
		return
	}
	for i := range trace.Cycles {
		trace.Cycles[i].BeforePhysicalRoot = ""
		trace.Cycles[i].AfterPhysicalRoot = ""
		for j := range trace.Cycles[i].Actions {
			trace.Cycles[i].Actions[j].DecisionReason = ""
		}
	}
}

func characterReadinessRequirements(context CharacterReadinessContext) ([]CharacterReadinessRequirement, error) {
	var requirements []CharacterReadinessRequirement
	for _, contract := range context.HardContracts {
		hash, err := characterAgentDigest(contract)
		if err != nil {
			return nil, err
		}
		requirements = append(requirements, CharacterReadinessRequirement{ID: "hard_" + strings.TrimPrefix(hash, "sha256:"), Contract: contract, DueNow: context.Chapter == context.BookLastChapter})
	}
	for _, obligation := range context.Obligations {
		if obligation.Hardness != ObligationHardnessV2("hard") {
			continue
		}
		requirements = append(requirements, CharacterReadinessRequirement{ID: obligation.ID, Contract: obligation.Contract, DueNow: obligation.DueNow})
	}
	return requirements, nil
}

func CharacterReadinessReviewInputDigest(input CharacterReadinessReviewInput) (string, error) {
	if (input.Policy != CharacterReadinessReviewPolicy && input.Policy != CharacterReadinessReviewPolicyV2) || len(input.Trace.Cycles) == 0 || (input.Policy == CharacterReadinessReviewPolicyV2) != (input.Context.Version == CharacterReadinessReviewPolicyV2) {
		return "", fmt.Errorf("invalid readiness review input")
	}
	context, err := FinalizeCharacterReadinessContext(input.Context)
	if err != nil {
		return "", err
	}
	if context.Digest != input.Context.Digest || input.RemainingCycles < 0 {
		return "", fmt.Errorf("readiness context/bounds mismatch")
	}
	for _, digest := range []string{input.ReviewProtocol, input.SessionDigest, input.Trace.FinalPhysicalRoot} {
		if err := validatePlanningV2Digest("readiness source binding", digest); err != nil {
			return "", err
		}
	}
	requirements, err := characterReadinessRequirements(context)
	if err != nil {
		return "", err
	}
	left, _ := json.Marshal(requirements)
	right, _ := json.Marshal(input.Requirements)
	if string(left) != string(right) {
		return "", fmt.Errorf("readiness input altered the frozen hard requirements")
	}
	for i, cycle := range input.Trace.Cycles {
		if cycle.Index != i+1 || cycle.StoryTime == nil || cycle.StoryTime.Chapter != input.Context.Chapter {
			return "", fmt.Errorf("readiness trace has invalid cycle identity/time")
		}
		digests := []string{cycle.CycleDigest, cycle.ArbitrationDigest}
		if input.Policy == CharacterReadinessReviewPolicyV2 {
			digests = append(digests, cycle.BeforePhysicalRoot, cycle.AfterPhysicalRoot)
		} else if cycle.BeforePhysicalRoot != "" || cycle.AfterPhysicalRoot != "" {
			return "", fmt.Errorf("legacy readiness trace cannot contain soft-event source roots")
		}
		for _, action := range cycle.Actions {
			if input.Policy == CharacterReadinessReviewPolicyV2 {
				if strings.TrimSpace(action.DecisionReason) == "" {
					return "", fmt.Errorf("soft-event readiness requires the character's exact decision reason")
				}
			} else if action.DecisionReason != "" {
				return "", fmt.Errorf("legacy readiness trace cannot contain soft-event decision reasons")
			}
		}
		for _, digest := range digests {
			if err := validatePlanningV2Digest("readiness cycle source", digest); err != nil {
				return "", err
			}
		}
	}
	return characterAgentDigest(input)
}

func FinalizeCharacterReadinessReview(input CharacterReadinessReviewInput, verdict CharacterReadinessVerdict) (CharacterChapterReadiness, error) {
	var result CharacterChapterReadiness
	inputDigest, err := CharacterReadinessReviewInputDigest(input)
	if err != nil {
		return result, err
	}
	allowed := map[string]bool{input.Trace.FinalPhysicalRoot: true}
	actual := map[string]bool{input.Trace.FinalPhysicalRoot: true}
	for _, cycle := range input.Trace.Cycles {
		allowed[cycle.CycleDigest], allowed[cycle.ArbitrationDigest] = true, true
		actual[cycle.CycleDigest], actual[cycle.ArbitrationDigest] = true, true
		if input.Policy == CharacterReadinessReviewPolicyV2 {
			allowed[cycle.BeforePhysicalRoot], allowed[cycle.AfterPhysicalRoot] = true, true
			actual[cycle.BeforePhysicalRoot], actual[cycle.AfterPhysicalRoot] = true, true
		}
		for _, action := range cycle.Actions {
			allowed[action.ProposalDigest] = true
		}
	}
	validateRefs := func(refs []string) error {
		if len(refs) == 0 || len(refs) > 12 {
			return fmt.Errorf("readiness must cite 1-12 actual evidence references")
		}
		actualCited := false
		for _, ref := range refs {
			if !allowed[ref] {
				return fmt.Errorf("readiness cites evidence outside its exact input")
			}
			actualCited = actualCited || actual[ref]
		}
		if !actualCited {
			return fmt.Errorf("a proposal alone cannot prove readiness; cite an actual cycle/arbitration/state")
		}
		return nil
	}
	if err := validateRefs(verdict.EvidenceRefs); err != nil {
		return result, err
	}
	checks := map[string]CharacterReadinessContractCheck{}
	for _, check := range verdict.ContractChecks {
		if _, duplicate := checks[check.ContractID]; duplicate {
			return result, fmt.Errorf("duplicate readiness contract check")
		}
		if err := validateRefs(check.EvidenceRefs); err != nil {
			return result, err
		}
		checks[check.ContractID] = check
	}
	if len(checks) != len(input.Requirements) {
		return result, fmt.Errorf("readiness must assess every hard requirement exactly once")
	}
	var impossible []string
	for _, requirement := range input.Requirements {
		check, ok := checks[requirement.ID]
		if !ok {
			return result, fmt.Errorf("readiness omitted a hard requirement")
		}
		switch check.Status {
		case "satisfied", "preserved", "pending", "impossible":
		default:
			return result, fmt.Errorf("invalid readiness contract status")
		}
		if check.Status == "impossible" {
			impossible = append(impossible, requirement.ID)
		}
		if verdict.Decision == "ready_for_plan" && ((requirement.DueNow && check.Status != "satisfied") || check.Status == "impossible") {
			return result, fmt.Errorf("unfulfilled due/impossible hard requirement cannot authorize chapter planning")
		}
	}
	last := input.Trace.Cycles[len(input.Trace.Cycles)-1]
	if last.HardContractStatus == "infeasible" {
		if verdict.Decision != "hard_conflict" {
			return result, fmt.Errorf("readiness cannot override an arbitrated hard conflict")
		}
		impossible = append(impossible, last.HardContractConflicts...)
	}
	if verdict.Decision == "hard_conflict" && len(impossible) == 0 {
		return result, fmt.Errorf("hard conflict requires an impossible hard contract, not a cycle/budget limit")
	}
	if verdict.Decision != "hard_conflict" && len(impossible) > 0 {
		return result, fmt.Errorf("impossible hard contracts require hard_conflict")
	}
	version := CharacterReadinessReviewedVersion
	if input.Policy == CharacterReadinessReviewPolicyV2 {
		if err := validateCharacterReadinessSoftEvent(input, verdict, impossible); err != nil {
			return result, err
		}
		version = CharacterReadinessReviewedVersionV3
	} else if verdict.SoftEvent != nil {
		return result, fmt.Errorf("legacy readiness verdict cannot classify a soft event")
	}
	result = CharacterChapterReadiness{Version: version, GenerationID: input.Context.GenerationID, Chapter: input.Context.Chapter, CycleDigest: last.CycleDigest, ReviewProtocol: input.ReviewProtocol, InputDigest: inputDigest, Decision: verdict.Decision, Reason: verdict.Reason, EvidenceRefs: verdict.EvidenceRefs, ContractChecks: verdict.ContractChecks, SoftEvent: verdict.SoftEvent, UnresolvedHardContracts: impossible}
	return FinalizeCharacterChapterReadiness(result)
}

func validateCharacterReadinessSoftEvent(input CharacterReadinessReviewInput, verdict CharacterReadinessVerdict, impossible []string) error {
	soft := verdict.SoftEvent
	if soft == nil {
		return fmt.Errorf("soft-event readiness verdict requires an explicit outcome")
	}
	validateSoftRefs := func(refs []string) error {
		if len(refs) == 0 || len(refs) > 12 {
			return fmt.Errorf("soft-event outcome must cite 1-12 exact evidence references")
		}
		allowed := map[string]bool{input.Trace.FinalPhysicalRoot: true}
		for _, cycle := range input.Trace.Cycles {
			allowed[cycle.CycleDigest], allowed[cycle.ArbitrationDigest] = true, true
			allowed[cycle.BeforePhysicalRoot], allowed[cycle.AfterPhysicalRoot] = true, true
			for _, action := range cycle.Actions {
				allowed[action.ProposalDigest] = true
			}
		}
		for _, ref := range refs {
			if !allowed[ref] {
				return fmt.Errorf("soft-event outcome cites evidence outside its exact input")
			}
		}
		return nil
	}
	if err := validateSoftRefs(soft.EvidenceRefs); err != nil {
		return err
	}
	closing := soft.Outcome == CharacterSoftEventOccurred || soft.Outcome == CharacterSoftEventRejected || soft.Outcome == CharacterSoftEventSuperseded
	if closing {
		if verdict.Decision != "ready_for_plan" || strings.TrimSpace(soft.ActorRef) == "" || strings.TrimSpace(soft.ProposalRef) == "" || strings.TrimSpace(soft.CharacterReason) == "" || strings.TrimSpace(soft.WorldConsequence) == "" {
			return fmt.Errorf("closed soft event requires ready_for_plan and exact actor/proposal/reason/consequence")
		}
		var cycle *CharacterReadinessCycleView
		for i := range input.Trace.Cycles {
			for _, ref := range soft.EvidenceRefs {
				if ref != input.Trace.Cycles[i].CycleDigest && ref != input.Trace.Cycles[i].ArbitrationDigest {
					continue
				}
				if cycle != nil && cycle != &input.Trace.Cycles[i] {
					return fmt.Errorf("soft-event outcome cites more than one actual cycle")
				}
				cycle = &input.Trace.Cycles[i]
			}
		}
		if cycle == nil {
			return fmt.Errorf("soft-event closure must cite its same-cycle actual result")
		}
		var action *CharacterReadinessAction
		for j := range cycle.Actions {
			candidate := &cycle.Actions[j]
			if candidate.ProposalDigest == soft.ProposalRef {
				if action != nil {
					return fmt.Errorf("soft-event proposal reference is ambiguous within its cited cycle")
				}
				action = candidate
			}
		}
		if action == nil || soft.ActorRef != action.AgentID || strings.TrimSpace(soft.CharacterReason) != strings.TrimSpace(action.DecisionReason) {
			return fmt.Errorf("soft-event actor/reason is not the exact selected character action")
		}
		consequence := strings.TrimSpace(soft.WorldConsequence)
		if consequence != strings.TrimSpace(action.ImmediateResult) && consequence != strings.TrimSpace(action.StateAfter) {
			return fmt.Errorf("soft-event consequence is not an exact arbitrated result")
		}
		proposalCited, actualCited := false, false
		for _, ref := range soft.EvidenceRefs {
			proposalCited = proposalCited || ref == soft.ProposalRef
			actualCited = actualCited || ref == cycle.CycleDigest || ref == cycle.ArbitrationDigest
		}
		if !proposalCited || !actualCited {
			return fmt.Errorf("soft-event closure must cite its proposal and same-cycle actual result")
		}
		if cycle.StoryTime == nil || (cycle.StoryTime.EndDay <= cycle.StoryTime.StartDay && cycle.BeforePhysicalRoot == cycle.AfterPhysicalRoot) {
			return fmt.Errorf("soft-event closure requires an observable world/character/relationship/knowledge consequence")
		}
		return nil
	}
	if soft.ActorRef != "" || soft.ProposalRef != "" || soft.CharacterReason != "" || soft.WorldConsequence != "" {
		return fmt.Errorf("non-closing soft-event outcome cannot invent an actor action or consequence")
	}
	switch soft.Outcome {
	case CharacterSoftEventPending:
		if verdict.Decision != "continue" || len(impossible) != 0 {
			return fmt.Errorf("DEFERRED requires continue without an impossible hard contract")
		}
	case CharacterSoftEventHardUnsatisfied:
		if verdict.Decision != "hard_conflict" || len(impossible) == 0 {
			return fmt.Errorf("HARD_CONTRACT_UNSATISFIED requires a proven impossible hard contract")
		}
	default:
		return fmt.Errorf("unsupported soft-event readiness outcome")
	}
	return nil
}

func ValidateCharacterReadinessReviewAudit(audit CharacterReadinessReviewAudit) error {
	if audit.ModelView != nil {
		codec, err := NewCharacterReadinessModelCodecV1(audit.Input)
		if err != nil {
			return err
		}
		if *audit.ModelView != codec.Binding() {
			return fmt.Errorf("readiness model view differs from its canonical input")
		}
	}
	want, err := FinalizeCharacterReadinessReview(audit.Input, CharacterReadinessVerdict{Decision: audit.Receipt.Decision, Reason: audit.Receipt.Reason, EvidenceRefs: audit.Receipt.EvidenceRefs, ContractChecks: audit.Receipt.ContractChecks, SoftEvent: audit.Receipt.SoftEvent})
	if err != nil {
		return err
	}
	left, _ := json.Marshal(want)
	right, _ := json.Marshal(audit.Receipt)
	if string(left) != string(right) {
		return fmt.Errorf("readiness review receipt/input mismatch")
	}
	return nil
}
