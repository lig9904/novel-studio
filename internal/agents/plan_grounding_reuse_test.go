package agents

import (
	"strings"
	"testing"

	"github.com/chenhongyang/novel-studio/internal/bootstrap"
	"github.com/chenhongyang/novel-studio/internal/domain"
)

func TestProjectedPlanReuseRequiresCurrentGroundingProtocol(t *testing.T) {
	cfg := bootstrap.Config{Provider: "local", ModelName: "grounding-model", Providers: map[string]bootstrap.ProviderConfig{"local": {Type: "openai"}}}
	models, err := bootstrap.NewModelSet(cfg)
	if err != nil {
		t.Fatal(err)
	}
	simulation := domain.ChapterWorldSimulation{Sources: []string{domain.PlanGroundingPolicyV1}}
	reviewer := NewPlanGroundingReviewer(cfg, models, nil)
	resolved, err := reviewer.ResolveForSimulation(simulation)
	if err != nil {
		t.Fatal(err)
	}
	plan := &domain.ChapterPlan{GroundingReview: &domain.PlanGroundingReceipt{
		ReviewProtocol: resolved.Protocol,
		Verdict:        domain.PlanGroundingVerdict{Pass: true},
	}}
	if err := validateProjectedPlanGroundingProtocol(cfg, models, simulation, plan); err != nil {
		t.Fatal(err)
	}
	plan.GroundingReview.ReviewProtocol = "sha256:" + strings.Repeat("0", 64)
	if err := validateProjectedPlanGroundingProtocol(cfg, models, simulation, plan); err == nil || !strings.Contains(err.Error(), "protocol is stale") {
		t.Fatalf("stale grounded plan was reusable: %v", err)
	}
}
