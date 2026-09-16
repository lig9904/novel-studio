package domain_test

import (
	"strings"
	"testing"

	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/testutil"
)

func TestReadinessChecksDueContractsWithoutForcingFutureEnding(t *testing.T) {
	_, _, _, input := testutil.CharacterReadiness(t, false)
	verdict := testutil.ReadyVerdict(input)
	r, err := domain.FinalizeCharacterReadinessReview(input, verdict)
	if err != nil {
		t.Fatal(err)
	}
	if r.Version != domain.CharacterReadinessReviewedVersion || r.InputDigest == "" || r.Decision != "ready_for_plan" {
		t.Fatal("missing reviewed readiness binding")
	}
	if err := domain.ValidateCharacterReadinessReviewAudit(domain.CharacterReadinessReviewAudit{Input: input, Receipt: r}); err != nil {
		t.Fatal(err)
	}
	_, _, _, finalInput := testutil.CharacterReadiness(t, true)
	if _, err := domain.FinalizeCharacterReadinessReview(finalInput, verdict); err == nil {
		t.Fatal("final chapter accepted an ending deferred to the future")
	}
}

func TestReadinessRejectsMissingInventedOrOnlyIntendedEvidence(t *testing.T) {
	for _, mode := range []string{"missing-check", "pending-due", "unknown-ref", "proposal-only", "limit-is-conflict", "altered-requirements", "altered-context"} {
		t.Run(mode, func(t *testing.T) {
			_, _, _, input := testutil.CharacterReadiness(t, false)
			verdict := testutil.ReadyVerdict(input)
			switch mode {
			case "missing-check":
				verdict.ContractChecks = verdict.ContractChecks[:1]
			case "pending-due":
				verdict.ContractChecks[1].Status = "pending"
			case "unknown-ref":
				verdict.EvidenceRefs = []string{"invented-result"}
			case "proposal-only":
				verdict.EvidenceRefs = []string{input.Trace.Cycles[0].Actions[0].ProposalDigest}
			case "limit-is-conflict":
				input.RemainingCycles = 0
				verdict.Decision = "hard_conflict"
			case "altered-requirements":
				input.Requirements[1].DueNow = false
			case "altered-context":
				input.Context.HardContracts = nil
			}
			if _, err := domain.FinalizeCharacterReadinessReview(input, verdict); err == nil {
				t.Fatal("invalid readiness verdict accepted")
			}
		})
	}
}

func TestReadinessReceiptTamperingFailsVerification(t *testing.T) {
	_, _, _, input := testutil.CharacterReadiness(t, false)
	r, err := domain.FinalizeCharacterReadinessReview(input, testutil.ReadyVerdict(input))
	if err != nil {
		t.Fatal(err)
	}
	r.Reason = "重新解释但保留旧摘要"
	if err := domain.ValidateCharacterReadinessReviewAudit(domain.CharacterReadinessReviewAudit{Input: input, Receipt: r}); err == nil {
		t.Fatal("readiness receipt tampering passed")
	}
}

func softEventReadinessFixture(t *testing.T) (domain.CharacterReadinessReviewInput, domain.CharacterReadinessVerdict) {
	t.Helper()
	_, _, _, input := testutil.CharacterReadiness(t, false)
	input.Policy = domain.CharacterReadinessReviewPolicyV2
	input.Context.Version = domain.CharacterReadinessReviewPolicyV2
	input.Context.Digest = ""
	var err error
	input.Context, err = domain.FinalizeCharacterReadinessContext(input.Context)
	if err != nil {
		t.Fatal(err)
	}
	cycle := &input.Trace.Cycles[0]
	cycle.BeforePhysicalRoot = "sha256:" + strings.Repeat("1", 64)
	cycle.AfterPhysicalRoot = "sha256:" + strings.Repeat("2", 64)
	cycle.Actions[0].DecisionReason = "当前认知与爱面子使我拒绝照旧行动"
	action := cycle.Actions[0]
	verdict := testutil.ReadyVerdict(input)
	verdict.SoftEvent = &domain.CharacterReadinessSoftEvent{Outcome: domain.CharacterSoftEventRejected, ActorRef: action.AgentID, ProposalRef: action.ProposalDigest,
		CharacterReason: action.DecisionReason, WorldConsequence: action.ImmediateResult, EvidenceRefs: []string{action.ProposalDigest, cycle.ArbitrationDigest}}
	return input, verdict
}

func TestSoftEventReadinessCaseAValidRejectionRequiresActualConsequence(t *testing.T) {
	input, verdict := softEventReadinessFixture(t)
	receipt, err := domain.FinalizeCharacterReadinessReview(input, verdict)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Version != domain.CharacterReadinessReviewedVersionV3 || receipt.Decision != "ready_for_plan" || receipt.SoftEvent == nil || receipt.SoftEvent.Outcome != domain.CharacterSoftEventRejected {
		t.Fatal("valid character rejection with actual consequence did not close the soft event")
	}
	if err := domain.ValidateCharacterReadinessReviewAudit(domain.CharacterReadinessReviewAudit{Input: input, Receipt: receipt}); err != nil {
		t.Fatal(err)
	}
	for _, mutate := range []func(*domain.CharacterReadinessReviewInput, *domain.CharacterReadinessVerdict){
		func(_ *domain.CharacterReadinessReviewInput, v *domain.CharacterReadinessVerdict) {
			v.SoftEvent.CharacterReason = "概括后的理由"
		},
		func(_ *domain.CharacterReadinessReviewInput, v *domain.CharacterReadinessVerdict) {
			v.SoftEvent.WorldConsequence = "并未实际发生后果"
		},
		func(_ *domain.CharacterReadinessReviewInput, v *domain.CharacterReadinessVerdict) {
			v.SoftEvent.EvidenceRefs = []string{v.SoftEvent.ProposalRef}
		},
		func(i *domain.CharacterReadinessReviewInput, _ *domain.CharacterReadinessVerdict) {
			i.Trace.Cycles[0].StoryTime.EndDay = i.Trace.Cycles[0].StoryTime.StartDay
			i.Trace.Cycles[0].AfterPhysicalRoot = i.Trace.Cycles[0].BeforePhysicalRoot
		},
	} {
		changedInput, changedVerdict := input, verdict
		soft := *verdict.SoftEvent
		soft.EvidenceRefs = append([]string(nil), verdict.SoftEvent.EvidenceRefs...)
		changedVerdict.SoftEvent = &soft
		mutate(&changedInput, &changedVerdict)
		if _, err := domain.FinalizeCharacterReadinessReview(changedInput, changedVerdict); err == nil {
			t.Fatal("rejection without exact reason and actual consequence was accepted")
		}
	}
}

func TestSoftEventReadinessCaseBActualChoiceSupersedesSoftOutline(t *testing.T) {
	input, verdict := softEventReadinessFixture(t)
	verdict.SoftEvent.Outcome = domain.CharacterSoftEventSuperseded
	receipt, err := domain.FinalizeCharacterReadinessReview(input, verdict)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Decision != "ready_for_plan" || receipt.SoftEvent == nil || receipt.SoftEvent.Outcome != domain.CharacterSoftEventSuperseded {
		t.Fatal("actual alternative choice did not supersede the soft outline")
	}
}

func TestSoftEventReadinessCaseCHardContractStillBlocksAutonomy(t *testing.T) {
	input, verdict := softEventReadinessFixture(t)
	verdict.Decision = "hard_conflict"
	for i := range verdict.ContractChecks {
		if verdict.ContractChecks[i].ContractID == "due_check" {
			verdict.ContractChecks[i].Status = "impossible"
		}
	}
	verdict.SoftEvent = &domain.CharacterReadinessSoftEvent{Outcome: domain.CharacterSoftEventHardUnsatisfied, EvidenceRefs: []string{input.Trace.Cycles[0].ArbitrationDigest}}
	receipt, err := domain.FinalizeCharacterReadinessReview(input, verdict)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Decision != "hard_conflict" || len(receipt.UnresolvedHardContracts) == 0 || receipt.SoftEvent.Outcome != domain.CharacterSoftEventHardUnsatisfied {
		t.Fatal("character autonomy incorrectly overrode an unsatisfied hard contract")
	}
	verdict.Decision = "ready_for_plan"
	if _, err := domain.FinalizeCharacterReadinessReview(input, verdict); err == nil {
		t.Fatal("hard-contract failure was incorrectly authorized for planning")
	}
}
