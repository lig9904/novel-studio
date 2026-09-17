package agents

import (
	"context"
	"testing"

	"github.com/chenhongyang/novel-studio/internal/bootstrap"
	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/voocel/agentcore"
)

func TestPlanGroundingReviewerResolvesBeforeCacheButDirectReviewKeepsSnapshot(t *testing.T) {
	cfg := bootstrap.Config{Provider: "local", ModelName: "judge-a", Providers: map[string]bootstrap.ProviderConfig{"local": {Type: "openai"}}}
	models, err := bootstrap.NewModelSet(cfg)
	if err != nil {
		t.Fatal(err)
	}
	var called []string
	models.SetAttemptDecorator(func(_ context.Context, _, _, name string, _ agentcore.ChatModel) agentcore.ChatModel {
		called = append(called, name)
		return &groundingProbeModel{args: `{"pass":true,"findings":[]}`}
	})
	initial := NewPlanGroundingReviewer(cfg, models, nil)
	if err := models.Swap("default", "local", "judge-b"); err != nil {
		t.Fatal(err)
	}
	input := domain.PlanGroundingInput{Policy: domain.PlanGroundingPolicyV1, ReviewProtocol: initial.Protocol}
	if _, err := initial.Review(context.Background(), input); err != nil {
		t.Fatal(err)
	}
	if len(called) != 1 || called[0] != "judge-a" {
		t.Fatalf("direct diagnostic protocol a ran hot-swapped model: %v", called)
	}
	resolved, err := initial.Resolve()
	if err != nil {
		t.Fatal(err)
	}
	if resolved.Protocol == initial.Protocol {
		t.Fatal("hot swap reused old cache protocol")
	}
	if _, err := resolved.Review(context.Background(), input); err == nil {
		t.Fatal("direct reviewer accepted another model's input protocol")
	}
	if len(called) != 1 {
		t.Fatal("mismatched review protocol reached provider")
	}
	if err := models.Swap("world_arbiter", "local", "judge-c"); err != nil {
		t.Fatal(err)
	}
	input.ReviewProtocol = resolved.Protocol
	if _, err := resolved.Review(context.Background(), input); err != nil {
		t.Fatal(err)
	}
	if len(called) != 2 || called[1] != "judge-b" {
		t.Fatalf("swap after cache-key resolution changed the paid target: %v", called)
	}
	latest, err := initial.Resolve()
	if err != nil {
		t.Fatal(err)
	}
	if latest.Protocol == resolved.Protocol {
		t.Fatal("new explicit override did not change cache identity")
	}
}

func TestPlanGroundingReviewerUsesDedicatedRoleWithoutMovingWorldArbiter(t *testing.T) {
	cfg := bootstrap.Config{
		Provider:  "local",
		ModelName: "default",
		Providers: map[string]bootstrap.ProviderConfig{"local": {Type: "openai"}},
		Roles: map[string]bootstrap.RoleConfig{
			"world_arbiter":  {Provider: "local", Model: "arbiter-model", ReasoningEffort: "high"},
			"plan_grounding": {Provider: "local", Model: "grounding-model", ReasoningEffort: "medium"},
		},
	}
	models, err := bootstrap.NewModelSet(cfg)
	if err != nil {
		t.Fatal(err)
	}
	var role, name string
	models.SetAttemptDecorator(func(_ context.Context, gotRole, _, gotName string, _ agentcore.ChatModel) agentcore.ChatModel {
		role, name = gotRole, gotName
		return &groundingProbeModel{args: `{"pass":true,"findings":[]}`}
	})
	reviewer := NewPlanGroundingReviewer(cfg, models, nil)
	input := domain.PlanGroundingInput{Policy: domain.PlanGroundingPolicyV1, ReviewProtocol: reviewer.Protocol}
	if _, err := reviewer.Review(context.Background(), input); err != nil {
		t.Fatal(err)
	}
	if role != "plan_grounding" || name != "grounding-model" {
		t.Fatalf("grounding reviewer routed to %s/%s, want plan_grounding/grounding-model", role, name)
	}
	if _, model, _ := models.CurrentSelection("world_arbiter"); model != "arbiter-model" {
		t.Fatalf("dedicated grounding route moved world_arbiter to %s", model)
	}
}
