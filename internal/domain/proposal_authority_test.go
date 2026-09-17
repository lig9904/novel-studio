package domain

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestStoryProposalRegistryCannotSelfApprove(t *testing.T) {
	proposal := StoryProposalV1{
		ID: "PROPOSAL-P001", Statement: "凤凰具有火属性。", Authority: StoryProposalAuthority,
		Status: StoryProposalStatusPending, SourceRefs: []string{"controlled-test"},
	}
	registry, err := FinalizeStoryProposalRegistryV1(StoryProposalRegistryV1{Proposals: []StoryProposalV1{proposal}})
	if err != nil {
		t.Fatal(err)
	}
	if registry.Policy != StoryProposalPolicyV1 || registry.Proposals[0].Digest == "" || registry.Digest == "" {
		t.Fatalf("proposal registry was not content-addressed: %+v", registry)
	}
	bad := proposal
	bad.HumanApproved, bad.AcceptedCanon = true, true
	if _, err := FinalizeStoryProposalRegistryV1(StoryProposalRegistryV1{Proposals: []StoryProposalV1{bad}}); err == nil {
		t.Fatal("proposal registry minted approval/canon authority")
	}
}

func TestStoryProposalClaimIdentityBlocksParaphraseAndStructuredForm(t *testing.T) {
	subjects, objects, err := DeriveStoryProposalClaimTermsV1("凤凰具有火属性。")
	if err != nil {
		t.Fatal(err)
	}
	proposal := StoryProposalV1{Statement: "凤凰具有火属性。", SubjectTerms: subjects, ObjectTerms: objects}
	for _, value := range []string{"凤凰具备火属性", "凤凰的属性是火", "凤凰具有火属性，但没有任何限制。", "凤凰具有火属性，治愈能力尚未批准。", "凤凰具有火属性且不能治愈。"} {
		if !StoryProposalMatchesTextV1(proposal, value) {
			t.Fatalf("proposal paraphrase bypassed identity terms: %s", value)
		}
	}
	if !StoryProposalMatchesValueV1(proposal, map[string]any{"character": "凤凰", "attribute": "火"}) {
		t.Fatal("structured proposal claim bypassed identity terms")
	}
	if !StoryProposalMatchesValueV1(proposal, map[string]any{"character": "凤凰", "attribute": "火", "healing_ability": "TBD"}) {
		t.Fatal("unrelated TBD attribute masked positive structured proposal claim")
	}
	for _, value := range []any{
		"凤凰仍在云端。岸边有人点火。",
		"凤凰不具有火属性。",
		"凤凰的火属性仍为 TBD。",
		"凤凰看着村民生火。",
		"凤凰拥有一支火把。",
		"凤凰看见村民具有火属性。",
		[]any{map[string]any{"character": "凤凰", "action": "观察"}, map[string]any{"character": "村民", "action": "生火"}},
		map[string]any{"character": "凤凰", "attribute": "火属性 TBD"},
		map[string]any{"character": "凤凰", "attribute": "无火属性"},
	} {
		if StoryProposalMatchesValueV1(proposal, value) {
			t.Fatalf("legal negative/unrelated value matched proposal: %#v", value)
		}
	}
	bad := StoryProposalV1{
		ID: "PROPOSAL-P002", Statement: "凤凰具有火属性。", Authority: StoryProposalAuthority, Status: StoryProposalStatusPending,
		SourceRefs: []string{"controlled-test"}, SubjectTerms: []string{"无关主语"}, ObjectTerms: []string{"无关对象"},
	}
	if _, err := FinalizeStoryProposalRegistryV1(StoryProposalRegistryV1{Proposals: []StoryProposalV1{bad}}); err == nil {
		t.Fatal("model-supplied unrelated claim identity was accepted")
	}
}

func TestProjectedBundleRejectsProposalParaphraseAndStructuredClaim(t *testing.T) {
	_, _, bundles := planningV2TestChain(t, 1)
	proposalBinding := func(t *testing.T) SourceBindingV2 {
		return SourceBindingV2{
			Kind: "proposal", Authority: SourceAuthorityProposal, SourceID: "PROPOSAL-P001",
			SourceDigest: planningV2TestDigest(t, "proposal-p001"), ExactReferences: []string{"proposal:PROPOSAL-P001"},
			UsableFacts: []string{"凤凰具有火属性。"}, Transformation: "只允许隔离讨论", DoNotUse: []string{"不得写成事实"},
		}
	}
	t.Run("paraphrase", func(t *testing.T) {
		bundle := bundles[0]
		bundle.ChapterPlan.Goal = "凤凰具备火属性"
		bundle.SourceBindings = append(bundle.SourceBindings, proposalBinding(t))
		bundle.RenderContext, _ = BindProjectedRenderContextV2(bundle.RenderContext, bundle)
		bundle.RenderContextSHA256, _ = ComputePlanningV2JSONDigest(bundle.RenderContext)
		bundle.BundleDigest = planningV2MustBundleDigest(t, bundle)
		if err := ValidateProjectedChapterBundle(bundle); err == nil || !strings.Contains(err.Error(), "pending proposal claim") {
			t.Fatalf("proposal paraphrase entered bundle: %v", err)
		}
	})
	t.Run("structured", func(t *testing.T) {
		bundle := bundles[0]
		bundle.SourceBindings = append(bundle.SourceBindings, proposalBinding(t))
		var context map[string]any
		if err := json.Unmarshal(bundle.RenderContext, &context); err != nil {
			t.Fatal(err)
		}
		context["proposal_probe"] = map[string]any{"character": "凤凰", "attribute": "火"}
		bundle.RenderContext, _ = json.Marshal(context)
		bundle.RenderContext, _ = BindProjectedRenderContextV2(bundle.RenderContext, bundle)
		bundle.RenderContextSHA256, _ = ComputePlanningV2JSONDigest(bundle.RenderContext)
		bundle.BundleDigest = planningV2MustBundleDigest(t, bundle)
		if err := ValidateProjectedChapterBundle(bundle); err == nil || !strings.Contains(err.Error(), "pending proposal claim") {
			t.Fatalf("structured proposal claim entered bundle: %v", err)
		}
	})
}

func TestProjectedBundleRejectsExplicitProposalAsStoryFact(t *testing.T) {
	_, _, bundles := planningV2TestChain(t, 1)
	bundle := bundles[0]
	bundle.ChapterPlan.Goal = "凤凰具有火属性。"
	bundle.SourceBindings = append(bundle.SourceBindings, SourceBindingV2{
		Kind: "proposal", Authority: SourceAuthorityProposal, SourceID: "PROPOSAL-P001",
		SourceDigest: planningV2TestDigest(t, "proposal-p001"), ExactReferences: []string{"proposal:PROPOSAL-P001"},
		UsableFacts: []string{"凤凰具有火属性。"}, Transformation: "只允许隔离讨论", DoNotUse: []string{"不得写成事实"},
	})
	bundle.RenderContext, _ = BindProjectedRenderContextV2(bundle.RenderContext, bundle)
	bundle.RenderContextSHA256, _ = ComputePlanningV2JSONDigest(bundle.RenderContext)
	bundle.BundleDigest = planningV2MustBundleDigest(t, bundle)
	if err := ValidateProjectedChapterBundle(bundle); err == nil || !strings.Contains(err.Error(), "pending proposal claim") {
		t.Fatalf("explicit proposal became bundle fact: %v", err)
	}

	bundle = bundles[0]
	bundle.SourceBindings = append(bundle.SourceBindings, SourceBindingV2{
		Kind: "proposal", Authority: SourceAuthorityProposal, SourceID: "PROPOSAL-P001",
		SourceDigest: planningV2TestDigest(t, "proposal-p001"), ExactReferences: []string{"proposal:PROPOSAL-P001"},
		UsableFacts: []string{"凤凰具有火属性。"}, Transformation: "只允许隔离讨论", DoNotUse: []string{"不得写成事实"},
	})
	bundle.ChapterPlan.Contract.ForbiddenMoves = append(bundle.ChapterPlan.Contract.ForbiddenMoves, "不得断言凤凰具有火属性。")
	if len(bundle.POVPlan.Scenes) == 0 {
		t.Fatal("fixture lacks POV scene")
	}
	bundle.POVPlan.Scenes[0].POVDoesNotKnow = append(bundle.POVPlan.Scenes[0].POVDoesNotKnow, "凤凰具有火属性。")
	bundle.RenderContext, _ = BindProjectedRenderContextV2(bundle.RenderContext, bundle)
	bundle.RenderContextSHA256, _ = ComputePlanningV2JSONDigest(bundle.RenderContext)
	bundle.BundleDigest = planningV2MustBundleDigest(t, bundle)
	if err := ValidateProjectedChapterBundle(bundle); err != nil {
		t.Fatalf("negative proposal boundary was mistaken for a fact: %v", err)
	}
}

func TestProjectedBundleProposalBindingCannotDowngradeAuthority(t *testing.T) {
	_, _, bundles := planningV2TestChain(t, 1)
	for _, authority := range []string{"", SourceAuthorityCanon} {
		bundle := bundles[0]
		bundle.SourceBindings = append(bundle.SourceBindings, SourceBindingV2{
			Kind: "proposal", Authority: authority, SourceID: "PROPOSAL-P001",
			SourceDigest: planningV2TestDigest(t, "proposal-p001"), ExactReferences: []string{"proposal:PROPOSAL-P001"},
			UsableFacts: []string{"凤凰具有火属性。"}, Transformation: "只允许隔离讨论", DoNotUse: []string{"不得写成事实"},
		})
		bundle.RenderContext, _ = BindProjectedRenderContextV2(bundle.RenderContext, bundle)
		bundle.RenderContextSHA256, _ = ComputePlanningV2JSONDigest(bundle.RenderContext)
		bundle.BundleDigest = planningV2MustBundleDigest(t, bundle)
		if err := ValidateProjectedChapterBundle(bundle); err == nil || !strings.Contains(err.Error(), "lacks PROPOSAL authority") {
			t.Fatalf("proposal authority downgrade %q was accepted: %v", authority, err)
		}
	}
}
