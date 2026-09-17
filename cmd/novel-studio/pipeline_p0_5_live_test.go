package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/chenhongyang/novel-studio/assets"
	"github.com/chenhongyang/novel-studio/internal/agents"
	"github.com/chenhongyang/novel-studio/internal/bootstrap"
	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/store"
)

// Explicit paid P0-5 evaluation. It reuses a copied frozen Chapter 1
// simulation/evidence workspace and runs only Planner -> Plan Grounding. The
// resulting bundle is written as isolated TEST_SCAFFOLD evidence; it is never
// published, sealed, promoted, rendered or copied into accepted canon.
func TestLiveP05FrozenPlannerProjection(t *testing.T) {
	evalRoot := os.Getenv("NOVEL_STUDIO_P0_5_EVAL_ROOT")
	configPath := os.Getenv("NOVEL_STUDIO_P0_5_CONFIG")
	resultDir := os.Getenv("NOVEL_STUDIO_P0_5_RESULT_DIR")
	if evalRoot == "" || configPath == "" || resultDir == "" {
		t.Skip("requires explicit P0-5 copied eval root, project config and result directory")
	}
	const generationID = "pg2_b2354a98e8714bba2baca89a"
	const chapter = 1
	isolated := filepath.Join(evalRoot, ".project-all", generationID, "output", "novel")
	liveDir := filepath.Join(evalRoot, "output", "novel")
	if err := os.MkdirAll(resultDir, 0o755); err != nil {
		t.Fatal(err)
	}

	frozenPaths := []string{
		filepath.Join(isolated, "meta", "chapter_simulations", "001.json"),
		filepath.Join(isolated, "meta", "character_agents", "activation_sessions", generationID, "000001", "chapter_evidence.json"),
		filepath.Join(isolated, "outline.json"),
		filepath.Join(isolated, "layered_outline.json"),
	}
	before := make(map[string][]byte, len(frozenPaths))
	for _, path := range frozenPaths {
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		before[path] = raw
	}

	cfg, err := bootstrap.LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}
	cfg.FillDefaults()
	var projected domain.ProjectedPlanningContextV2
	p05ReadJSON(t, filepath.Join(isolated, "meta", "project_all_state.json"), &projected)
	var generation domain.PlanningGenerationV2
	p05ReadJSON(t, filepath.Join(liveDir, "meta", "planning", "v2", ".building", generationID, "generation.json"), &generation)
	var registry domain.ObligationRegistryV2
	p05ReadJSON(t, filepath.Join(liveDir, "meta", "planning", "v2", ".building", generationID, "obligation_registry.json"), &registry)

	st := store.NewStore(isolated)
	if err := st.Init(); err != nil {
		t.Fatal(err)
	}
	outline, err := st.Outline.GetChapterOutline(chapter)
	if err != nil || outline == nil {
		t.Fatalf("load chapter outline: %v", err)
	}
	boundary := agents.ProjectedArcBoundary{
		Volume:                       1,
		Arc:                          1,
		Title:                        "TEST_SCAFFOLD：潮边三项探针",
		Goal:                         "TEST_SCAFFOLD：在三章有界测试内验证角色自治、知识边界与有限海岸事件",
		FirstChapter:                 1,
		LastChapter:                  3,
		BookLastChapter:              3,
		CharacterProtocolPinned:      true,
		CharacterActivationPolicy:    generation.CharacterActivationPolicy,
		MaxCharacterActivationCycles: generation.MaxCharacterActivationCycles,
	}

	usagePath := filepath.Join(liveDir, store.UsageAuditPath)
	usageBefore, _ := os.ReadFile(usagePath)
	accounting, err := newPipelineProjectAllAccounting(context.Background(), cfg, store.NewStore(liveDir), st, generationID)
	if err != nil {
		t.Fatal(err)
	}
	started := time.Now()
	artifacts, runErr := agents.RunProjectedChapterPlanning(
		accounting.ctx,
		cfg,
		assets.Load(cfg.Style),
		isolated,
		chapter,
		projected.ContextDigest,
		generation.CharacterAgentProtocol,
		boundary,
		accounting.hooks(),
	)
	closeErr := accounting.close()
	if runErr != nil {
		t.Fatal(runErr)
	}
	if closeErr != nil {
		t.Fatal(closeErr)
	}
	if artifacts == nil || artifacts.Plan == nil || artifacts.WorldSimulation == nil {
		t.Fatal("Planner returned no formal artifacts")
	}
	plan := artifacts.Plan
	if plan.GroundingReview == nil || !plan.GroundingReview.Verdict.Pass || len(plan.GroundingReview.Verdict.Findings) != 0 {
		t.Fatalf("plan lacks a passing exact grounding receipt: %+v", plan.GroundingReview)
	}
	if plan.CausalSimulation.ProtagonistDecision != artifacts.WorldSimulation.ProtagonistProjection.ChosenDecision {
		t.Fatal("plan changed the frozen protagonist decision")
	}
	type externalProjection struct {
		UsableDetails      []string `json:"usable_details"`
		TransformationRule string   `json:"transformation_rule"`
	}
	type realityProjection struct {
		UsableDetail  string `json:"usable_detail"`
		TransformedAs string `json:"transformed_as"`
		ChapterUse    string `json:"chapter_use"`
	}
	type lensProjection struct {
		Target string `json:"target"`
		Move   string `json:"move"`
		Why    string `json:"why"`
	}
	type endingProjection struct {
		ConcreteAnchor  string `json:"concrete_anchor"`
		Consequence     string `json:"consequence"`
		NextChapterPull string `json:"next_chapter_pull"`
	}
	type groundingProjection struct {
		Detail        string `json:"detail"`
		TransformedAs string `json:"transformed_as"`
		SceneAnchor   string `json:"scene_anchor"`
	}
	type rewardProjection struct {
		FirstChapterSmallWin string                    `json:"first_chapter_small_win"`
		NewDebtOrCost        string                    `json:"new_debt_or_cost"`
		PayoffVisibility     string                    `json:"payoff_visibility"`
		RewardLadder         []domain.ReaderRewardStep `json:"reward_ladder"`
	}
	external := make([]externalProjection, 0, len(plan.CausalSimulation.ExternalRefs))
	for _, item := range plan.CausalSimulation.ExternalRefs {
		external = append(external, externalProjection{UsableDetails: item.UsableDetails, TransformationRule: item.TransformationRule})
	}
	entertainment := plan.CausalSimulation.EntertainmentPlan
	entertainment.ForbiddenComedy = nil
	reward := rewardProjection{
		FirstChapterSmallWin: plan.CausalSimulation.ReaderRewardPlan.FirstChapterSmallWin,
		NewDebtOrCost:        plan.CausalSimulation.ReaderRewardPlan.NewDebtOrCost,
		PayoffVisibility:     plan.CausalSimulation.ReaderRewardPlan.PayoffVisibility,
		RewardLadder:         plan.CausalSimulation.ReaderRewardPlan.RewardLadder,
	}
	longform := plan.CausalSimulation.LongformOpening
	longform.RevealBudget, longform.RetentionRisks = nil, nil
	reality := make([]realityProjection, 0, len(plan.CausalSimulation.RealitySupport))
	for _, item := range plan.CausalSimulation.RealitySupport {
		reality = append(reality, realityProjection{UsableDetail: item.UsableDetail, TransformedAs: item.TransformedAs, ChapterUse: item.ChapterUse})
	}
	literary := map[string]any{}
	if plan.CausalSimulation.LiteraryRendering != nil {
		lenses := make([]lensProjection, 0, len(plan.CausalSimulation.LiteraryRendering.ActiveLenses))
		for _, item := range plan.CausalSimulation.LiteraryRendering.ActiveLenses {
			lenses = append(lenses, lensProjection{Target: item.Target, Move: item.Move, Why: item.Why})
		}
		literary = map[string]any{"scene_modes": plan.CausalSimulation.LiteraryRendering.SceneModes, "active_lenses": lenses, "afterimage": plan.CausalSimulation.LiteraryRendering.Afterimage}
	}
	ending := endingProjection{
		ConcreteAnchor:  plan.CausalSimulation.EndingContract.ConcreteAnchor,
		Consequence:     plan.CausalSimulation.EndingContract.Consequence,
		NextChapterPull: plan.CausalSimulation.EndingContract.NextChapterPull,
	}
	grounding := make([]groundingProjection, 0, len(plan.CausalSimulation.GroundingDetails))
	for _, item := range plan.CausalSimulation.GroundingDetails {
		grounding = append(grounding, groundingProjection{Detail: item.Detail, TransformedAs: item.TransformedAs, SceneAnchor: item.SceneAnchor})
	}

	// This is a fixture-specific conservative check over positive, Drafter-consumed
	// fact surfaces. Source references and forbidden/negative metadata are excluded.
	factSurface, err := json.Marshal(map[string]any{
		"contract":          map[string]any{"required_beats": plan.Contract.RequiredBeats, "payoff_points": plan.Contract.PayoffPoints, "scene_anchors": plan.Contract.SceneAnchors},
		"environment_state": plan.CausalSimulation.EnvironmentState,
		"causal_beats":      plan.CausalSimulation.CausalBeats,
		"render_capacity":   plan.CausalSimulation.RenderCapacity.SceneUnits, "ending": ending,
		"outcome_shift": plan.CausalSimulation.OutcomeShift, "decision_points": plan.CausalSimulation.DecisionPoints,
		"entertainment": entertainment, "reward": reward, "longform": longform,
		"external_projection": external, "grounding_details": grounding,
		"reality_support": reality, "literary_rendering": literary,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"浮标绳", "水尺", "现场来客", "转述者", "接受声望机会", "凤凰火焰", "涅槃", "重生", "治愈", "不死"} {
		if bytes.Contains(factSurface, []byte(forbidden)) {
			t.Fatalf("fact-bearing plan surface retained unsupported content %q", forbidden)
		}
	}

	bundle, nextRegistry, err := buildPipelineProjectedChapterBundle(
		generation,
		*outline,
		"",
		generation.BaseStateRoot,
		artifacts,
		registry,
	)
	if err != nil {
		t.Fatal(err)
	}
	if err := domain.ValidateProjectedChapterBundle(bundle); err != nil {
		t.Fatal(err)
	}
	if bundle.Chapter != 1 || bundle.GenerationID != generationID || bundle.ChapterPlan.GroundingReview == nil || !bundle.ChapterPlan.GroundingReview.Verdict.Pass {
		t.Fatal("isolated bundle lost its exact grounded Chapter 1 identity")
	}

	p05WriteJSON(t, filepath.Join(resultDir, "chapter-001.plan.json"), plan)
	p05WriteJSON(t, filepath.Join(resultDir, "chapter-001.bundle.json"), bundle)
	p05WriteJSON(t, filepath.Join(resultDir, "obligation-registry-after-ch01.json"), nextRegistry)
	p05WriteJSON(t, filepath.Join(resultDir, "runtime-result.json"), map[string]any{
		"schema": "p0-5-plan-projection-runtime.v1", "status": "PASS",
		"generation_id": generationID, "chapter": chapter, "simulation_id": artifacts.WorldSimulation.SimulationID,
		"plan_grounding_pass": true, "grounding_receipt_digest": plan.GroundingReview.Digest,
		"plan_grounding_input_digest": plan.GroundingReview.InputDigest, "bundle_digest": bundle.BundleDigest,
		"protagonist_decision": plan.CausalSimulation.ProtagonistDecision,
		"elapsed_ms":           time.Since(started).Milliseconds(), "classification": "TEST_SCAFFOLD",
		"published": false, "sealed": false, "promoted": false, "rendered": false,
	})

	if auditPath := filepath.Join(isolated, "meta", "planning", "grounding", strings.TrimPrefix(plan.GroundingReview.InputDigest, "sha256:")+".json"); auditPath != "" {
		raw, err := os.ReadFile(auditPath)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(resultDir, "plan-grounding-audit.json"), raw, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	usageAfter, _ := os.ReadFile(usagePath)
	if len(usageAfter) >= len(usageBefore) {
		if err := os.WriteFile(filepath.Join(resultDir, "runtime-usage-new.jsonl"), usageAfter[len(usageBefore):], 0o644); err != nil {
			t.Fatal(err)
		}
	}
	for path, old := range before {
		now, err := os.ReadFile(path)
		if err != nil || !bytes.Equal(old, now) {
			t.Fatalf("frozen Story Simulation source changed: %s err=%v", path, err)
		}
	}
	t.Logf("P0-5 PASS plan=%s grounding=%s bundle=%s elapsed=%s", plan.Title, plan.GroundingReview.Digest, bundle.BundleDigest, time.Since(started))
}

// Explicit read-only semantic regression for the runtime-discovered RAG
// authority bypass. It does not run Planner or mutate the copied workspace.
func TestLiveP05GroundingRejectsRAGAuthorityBypass(t *testing.T) {
	evalRoot := os.Getenv("NOVEL_STUDIO_P0_5_EVAL_ROOT")
	configPath := os.Getenv("NOVEL_STUDIO_P0_5_CONFIG")
	resultDir := os.Getenv("NOVEL_STUDIO_P0_5_RESULT_DIR")
	if evalRoot == "" || configPath == "" || resultDir == "" {
		t.Skip("requires explicit P0-5 copied eval root, project config and result directory")
	}
	const generationID = "pg2_b2354a98e8714bba2baca89a"
	isolated := filepath.Join(evalRoot, ".project-all", generationID, "output", "novel")
	var plan domain.ChapterPlan
	var simulation domain.ChapterWorldSimulation
	var evidence domain.CharacterActivationChapterEvidence
	p05ReadJSON(t, filepath.Join(isolated, "drafts", "01.plan.json"), &plan)
	p05ReadJSON(t, filepath.Join(isolated, "meta", "chapter_simulations", "001.json"), &simulation)
	p05ReadJSON(t, filepath.Join(isolated, "meta", "character_agents", "activation_sessions", generationID, "000001", "chapter_evidence.json"), &evidence)
	cfg, err := bootstrap.LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}
	cfg.FillDefaults()
	models, err := bootstrap.NewModelSet(cfg)
	if err != nil {
		t.Fatal(err)
	}
	reviewer := agents.NewPlanGroundingReviewer(cfg, models, nil)
	reviewer, err = reviewer.ResolveForSimulation(simulation)
	if err != nil {
		t.Fatal(err)
	}
	input, err := domain.NewActivationPlanGroundingInput(plan, simulation, evidence, reviewer.Protocol)
	if err != nil {
		t.Fatal(err)
	}
	verdict, err := reviewer.Review(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	matched := false
	for _, finding := range verdict.Findings {
		if strings.HasPrefix(finding.PlanPath, "/plan/causal_simulation/external_reference_plan/") {
			matched = true
		}
	}
	if verdict.Pass || !matched {
		t.Fatalf("strengthened Grounding accepted RAG authority bypass: %+v", verdict)
	}
	p05WriteJSON(t, filepath.Join(resultDir, "grounding-rag-bypass-rejection.json"), map[string]any{
		"schema": "p0-5-grounding-rag-authority-regression.v1", "status": "PASS",
		"grounding_pass": verdict.Pass, "findings": verdict.Findings,
		"review_protocol": reviewer.Protocol, "classification": "TEST_SCAFFOLD",
		"planner_rerun": false, "story_simulation_rerun": false,
	})
}

func p05ReadJSON(t *testing.T, path string, out any) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, out); err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}
}

func p05WriteJSON(t *testing.T, path string, value any) {
	t.Helper()
	raw, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(raw, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
}
