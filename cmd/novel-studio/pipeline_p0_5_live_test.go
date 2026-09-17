package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/chenhongyang/novel-studio/assets"
	"github.com/chenhongyang/novel-studio/internal/agents"
	"github.com/chenhongyang/novel-studio/internal/bootstrap"
	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/store"
	"github.com/chenhongyang/novel-studio/internal/testutil"
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
		filepath.Join(isolated, "meta", "project_all_state.json"),
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
	plannerProtocol := agents.ProjectAllPlanningProtocolWithActivationProducer(
		assets.Load(cfg.Style).Prompts.Planner,
		generation.CharacterAgentProtocol,
		generation.MaxCharacterActivationCycles,
		generation.CharacterActivationPolicy,
		cfg.CharacterAgents.FrozenActivationProducer,
	)
	recoveryContract, err := agents.PreparePlannerOnlyRecoveryContract(
		st,
		"p0-5-controlled-recovery-"+time.Now().UTC().Format("20060102T150405.000000000"),
		generationID,
		chapter,
		projected.ContextDigest,
		plannerProtocol,
	)
	if err != nil {
		t.Fatal(err)
	}
	p05WriteJSON(t, filepath.Join(resultDir, "planner-only-recovery-contract.json"), recoveryContract)
	tracePath := filepath.Join(resultDir, "planner-only-trace.jsonl")
	if err := os.WriteFile(tracePath, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	recoveryRequest := agents.PlannerOnlyRecoveryRequest{
		Contract: recoveryContract,
		Trace: func(event agents.PlannerOnlyTraceEvent) error {
			return p05AppendJSONL(tracePath, event)
		},
	}

	usagePath := filepath.Join(liveDir, store.UsageAuditPath)
	usageBefore, _ := os.ReadFile(usagePath)
	accounting, err := newPipelineProjectAllAccounting(context.Background(), cfg, store.NewStore(liveDir), st, generationID)
	if err != nil {
		t.Fatal(err)
	}
	started := time.Now()
	artifacts, runErr := agents.RunPlannerOnlyProjectedChapterPlanning(
		accounting.ctx,
		cfg,
		assets.Load(cfg.Style),
		isolated,
		chapter,
		projected.ContextDigest,
		generation.CharacterAgentProtocol,
		boundary,
		recoveryRequest,
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
	var activationEvidence domain.CharacterActivationChapterEvidence
	p05ReadJSON(t, filepath.Join(isolated, "meta", "character_agents", "activation_sessions", generationID, "000001", "chapter_evidence.json"), &activationEvidence)
	auditPath := filepath.Join(isolated, "meta", "planning", "grounding", strings.TrimPrefix(plan.GroundingReview.InputDigest, "sha256:")+".json")
	var groundingAudit domain.PlanGroundingAudit
	p05ReadJSON(t, auditPath, &groundingAudit)
	p05WriteJSON(t, filepath.Join(resultDir, "chapter-001.plan.json"), plan)
	p05WriteJSON(t, filepath.Join(resultDir, "plan-grounding-audit.json"), groundingAudit)
	postCheck, err := p05AuthorityPostCheck(*plan, *artifacts.WorldSimulation, activationEvidence, groundingAudit)
	if err != nil {
		t.Fatal(err)
	}
	p05WriteJSON(t, filepath.Join(resultDir, "authority-post-check.json"), postCheck)
	if postCheck.Status != "PASS" {
		t.Fatalf("authority post-check rejected final plan: %+v", postCheck.Findings)
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

	previousBundleDigest, preStateRoot, err := pipelineProjectAllTail(generation, nil)
	if err != nil {
		t.Fatal(err)
	}
	bundle, nextRegistry, err := buildPipelineProjectedChapterBundle(
		generation,
		*outline,
		previousBundleDigest,
		preStateRoot,
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

type p05AuthorityPostCheckFinding struct {
	Path   string `json:"path"`
	Kind   string `json:"kind"`
	Detail string `json:"detail"`
}

type p05AuthorityPostCheckResult struct {
	Schema                 string                         `json:"schema"`
	Classification         string                         `json:"classification"`
	Status                 string                         `json:"status"`
	GroundingInputDigest   string                         `json:"grounding_input_digest"`
	GroundingReceiptDigest string                         `json:"grounding_receipt_digest"`
	ReceiptBinding         string                         `json:"receipt_binding"`
	SourceBinding          string                         `json:"source_binding"`
	ProposalRegistryGate   string                         `json:"proposal_registry_gate"`
	Findings               []p05AuthorityPostCheckFinding `json:"findings"`
}

// p05AuthorityPostCheck is a fixture-scoped fail-closed gate between current
// Grounding and Bundle construction. It validates the exact receipt/input
// binding and blocks the previously observed unsupported fox tide-mark risk.
// It is deliberately not a general natural-language authority engine.
func p05AuthorityPostCheck(
	plan domain.ChapterPlan,
	simulation domain.ChapterWorldSimulation,
	evidence domain.CharacterActivationChapterEvidence,
	audit domain.PlanGroundingAudit,
) (p05AuthorityPostCheckResult, error) {
	result := p05AuthorityPostCheckResult{
		Schema: "p0-5r-authority-post-check.v1", Classification: "TEST_SCAFFOLD",
		Status: "FAIL", ReceiptBinding: "FAIL",
		SourceBinding: "NOT_AVAILABLE_PRE_BUNDLE", ProposalRegistryGate: "NOT_EXECUTED",
		Findings: []p05AuthorityPostCheckFinding{},
	}
	if plan.GroundingReview == nil || !plan.GroundingReview.Verdict.Pass || len(plan.GroundingReview.Verdict.Findings) != 0 {
		return result, nil
	}
	if err := domain.ValidatePlanGroundingAudit(audit); err != nil {
		return result, err
	}
	planReceipt, err := json.Marshal(plan.GroundingReview)
	if err != nil {
		return result, err
	}
	auditReceipt, err := json.Marshal(audit.Receipt)
	if err != nil {
		return result, err
	}
	if !bytes.Equal(planReceipt, auditReceipt) {
		result.Findings = append(result.Findings, p05AuthorityPostCheckFinding{
			Path: "/grounding_review", Kind: "STALE_OR_MISMATCHED_RECEIPT",
			Detail: "Plan receipt object differs from the Host-validated current audit receipt",
		})
		return result, nil
	}
	input, err := domain.NewActivationPlanGroundingInput(plan, simulation, evidence, plan.GroundingReview.ReviewProtocol)
	if err != nil {
		return result, err
	}
	inputDigest, err := domain.PlanGroundingInputDigest(input)
	if err != nil {
		return result, err
	}
	result.GroundingInputDigest = inputDigest
	result.GroundingReceiptDigest = plan.GroundingReview.Digest
	if inputDigest != plan.GroundingReview.InputDigest ||
		audit.Receipt.InputDigest != inputDigest ||
		audit.Receipt.Digest != plan.GroundingReview.Digest ||
		audit.Receipt.ReviewProtocol != plan.GroundingReview.ReviewProtocol {
		result.Findings = append(result.Findings, p05AuthorityPostCheckFinding{
			Path: "/grounding_review", Kind: "STALE_OR_MISMATCHED_RECEIPT",
			Detail: "final Plan, current Grounding input and receipt are not exactly bound",
		})
		return result, nil
	}
	result.ReceiptBinding = "PASS"

	claims, err := p05PositiveAuthorityClaims(plan)
	if err != nil {
		return result, err
	}
	for _, claim := range claims {
		if p05UnsupportedFoxErasure(claim.Text) {
			result.Findings = append(result.Findings, p05AuthorityPostCheckFinding{
				Path: claim.Path,
				Kind: "UNSUPPORTED_CAUSAL_PRESSURE", Detail: "fox tide-mark erasure risk is not established by frozen authority",
			})
		}
	}
	for _, forbidden := range []string{
		"凤凰具有火属性", "凤凰火属性", "凤凰火焰", "凤凰治愈", "凤凰涅槃", "凤凰重生", "凤凰不死",
	} {
		for _, claim := range claims {
			if p05AffirmativePhrase(claim.Text, forbidden) {
				result.Findings = append(result.Findings, p05AuthorityPostCheckFinding{
					Path: claim.Path, Kind: "PROPOSAL_OR_UNESTABLISHED_ABILITY", Detail: forbidden,
				})
			}
		}
	}
	if len(result.Findings) == 0 {
		result.Status = "PASS"
	}
	return result, nil
}

type p05AuthorityClaim struct {
	Path string
	Text string
}

func p05PositiveAuthorityClaims(plan domain.ChapterPlan) ([]p05AuthorityClaim, error) {
	plan.GroundingReview = nil
	raw, err := json.Marshal(plan)
	if err != nil {
		return nil, err
	}
	var root any
	if err := json.Unmarshal(raw, &root); err != nil {
		return nil, err
	}
	excluded := map[string]bool{
		"grounding_review": true, "forbidden_moves": true, "forbidden_usage": true,
		"forbidden_comedy": true, "forbidden_reward_patterns": true, "forbidden_endings": true,
		"forbidden_direct_use": true, "do_not_use": true, "source_refs": true,
		"query_or_need": true, "reveal_budget": true, "knowledge_boundary": true,
		"context_sources": true, "information_gaps": true, "review_refinement": true,
		"scene_constraints": true,
	}
	claims := []p05AuthorityClaim{}
	var walk func(string, any)
	walk = func(path string, value any) {
		switch node := value.(type) {
		case map[string]any:
			keys := make([]string, 0, len(node))
			for key := range node {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			for _, key := range keys {
				if excluded[key] {
					continue
				}
				walk(path+"/"+strings.ReplaceAll(strings.ReplaceAll(key, "~", "~0"), "/", "~1"), node[key])
			}
		case []any:
			for index, item := range node {
				walk(path+"/"+fmt.Sprint(index), item)
			}
		case string:
			if strings.TrimSpace(node) != "" {
				claims = append(claims, p05AuthorityClaim{Path: path, Text: node})
			}
		}
	}
	walk("/plan", root)
	return claims, nil
}

func p05UnsupportedFoxErasure(value string) bool {
	value = strings.ReplaceAll(strings.TrimSpace(value), " ", "")
	if value == "" || !(strings.Contains(value, "潮痕") || strings.Contains(value, "标记")) {
		return false
	}
	water := []string{"浪花", "海浪", "潮水", "水流", "浪头"}
	erasure := []string{"抹掉", "冲掉", "冲散", "冲刷", "洗掉", "消失", "毁掉", "灭失", "褪去"}
	valueRunes := []rune(value)
	for _, effect := range erasure {
		effectRunes := []rune(effect)
		for i := 0; i+len(effectRunes) <= len(valueRunes); i++ {
			if string(valueRunes[i:i+len(effectRunes)]) != effect || !p05AffirmativeAt(valueRunes, effectRunes, i) {
				continue
			}
			start, end := i-24, i+len(effectRunes)+12
			if start < 0 {
				start = 0
			}
			if end > len(valueRunes) {
				end = len(valueRunes)
			}
			context := string(valueRunes[start:end])
			for _, cause := range water {
				if strings.Contains(context, cause) {
					return true
				}
			}
		}
	}
	return false
}

func p05AffirmativePhrase(value, phrase string) bool {
	valueRunes, phraseRunes := []rune(value), []rune(phrase)
	if len(phraseRunes) == 0 || len(valueRunes) < len(phraseRunes) {
		return false
	}
	for i := 0; i+len(phraseRunes) <= len(valueRunes); i++ {
		if string(valueRunes[i:i+len(phraseRunes)]) == phrase && p05AffirmativeAt(valueRunes, phraseRunes, i) {
			return true
		}
	}
	return false
}

func p05AffirmativeAt(valueRunes, phraseRunes []rune, index int) bool {
	if index < 0 || index+len(phraseRunes) > len(valueRunes) || string(valueRunes[index:index+len(phraseRunes)]) != string(phraseRunes) {
		return false
	}
	start := index - 24
	if start < 0 {
		start = 0
	}
	prefix := string(valueRunes[start:index])
	if boundary := strings.LastIndexAny(prefix, "。；;！!?，,"); boundary >= 0 {
		prefix = prefix[boundary+1:]
	}
	for _, marker := range []string{
		"不得", "禁止", "不可", "不能", "不会", "不应", "不允许", "未调用", "未发生", "没有", "无权", "避免",
		"删除", "删去", "移除", "去掉", "剔除", "不得保留", "不再保留",
	} {
		if strings.Contains(prefix, marker) {
			return false
		}
	}
	return true
}

func TestP05AuthorityPostCheckRejectsKnownOffscreenFact(t *testing.T) {
	if !p05UnsupportedFoxErasure("浪花随时会抹掉潮痕，含义与来源都未确认") ||
		!p05UnsupportedFoxErasure("潮水可能把私密标记冲刷消失") {
		t.Fatal("known unsupported fox erasure semantics escaped the post-check")
	}
	if p05UnsupportedFoxErasure("潮痕的含义、来源与行动价值均未确认") {
		t.Fatal("supported uncertainty boundary was misclassified as erasure")
	}
	if p05UnsupportedFoxErasure("浪花不会抹掉潮痕；删除‘浪花会抹掉潮痕’这一未裁决描述") {
		t.Fatal("negated fact and deletion note were misclassified as affirmative erasure")
	}
	if !p05UnsupportedFoxErasure("浪花不会抹掉潮痕；另一时段浪花会抹掉潮痕") {
		t.Fatal("a later affirmative occurrence borrowed the earlier negation")
	}
	plan := domain.ChapterPlan{CausalSimulation: domain.ChapterCausalSimulation{OffscreenStage: []domain.CharacterStageRecord{{
		Character: "九尾狐", Pressure: "浪花随时会抹掉潮痕，含义与来源都未确认",
	}}}}
	claims, err := p05PositiveAuthorityClaims(plan)
	if err != nil {
		t.Fatal(err)
	}
	matched := false
	for _, claim := range claims {
		if claim.Path == "/plan/causal_simulation/offscreen_character_stage/0/pressure" && p05UnsupportedFoxErasure(claim.Text) {
			matched = true
		}
	}
	if !matched {
		t.Fatal("known runtime field path escaped the positive scalar projection")
	}
}

func TestP05AuthorityPostCheckBindsReceiptAndExcludesNegativeMetadata(t *testing.T) {
	evidence := testutil.CharacterActivationChapter(t)
	simulation, err := domain.BuildCharacterActivationSimulation(evidence, "tick_fixture", nil)
	if err != nil {
		t.Fatal(err)
	}
	base := domain.ChapterPlan{Chapter: simulation.Chapter, Notes: "删除‘浪花会抹掉潮痕’这一未裁决描述", CausalSimulation: domain.ChapterCausalSimulation{
		WorldSimulationID:   simulation.SimulationID,
		ProtagonistDecision: simulation.ProtagonistProjection.ChosenDecision,
		EndingContract:      domain.EndingConsequenceContract{ForbiddenEndings: []string{"不得让凤凰治愈"}},
		DialogueBlueprints:  []domain.DialogueSceneBlueprint{{DoNotUse: []string{"不得让凤凰治愈"}}},
	}}
	makeAudit := func(plan domain.ChapterPlan, verdict domain.PlanGroundingVerdict) (domain.ChapterPlan, domain.PlanGroundingAudit) {
		t.Helper()
		input, err := domain.NewActivationPlanGroundingInput(plan, simulation, evidence, evidence.ProtocolDigest)
		if err != nil {
			t.Fatal(err)
		}
		receipt, err := domain.FinalizePlanGroundingReceipt(input, verdict)
		if err != nil {
			t.Fatal(err)
		}
		plan.GroundingReview = &receipt
		return plan, domain.PlanGroundingAudit{Input: input, Receipt: receipt}
	}

	passingPlan, passingAudit := makeAudit(base, domain.PlanGroundingVerdict{Pass: true, Findings: []domain.PlanGroundingFinding{}})
	result, err := p05AuthorityPostCheck(passingPlan, simulation, evidence, passingAudit)
	if err != nil || result.Status != "PASS" {
		t.Fatalf("valid current receipt or negative metadata was rejected: result=%+v err=%v", result, err)
	}
	stale := passingPlan
	stale.Goal = "changed after receipt"
	result, err = p05AuthorityPostCheck(stale, simulation, evidence, passingAudit)
	if err != nil || result.Status != "FAIL" || len(result.Findings) == 0 || result.Findings[0].Kind != "STALE_OR_MISMATCHED_RECEIPT" {
		t.Fatalf("post-receipt Plan mutation escaped: result=%+v err=%v", result, err)
	}

	unsafe := base
	unsafe.CausalSimulation.ExternalRefs = []domain.ExternalReferencePlan{{UsableDetails: []string{"浪花会把私密潮痕冲掉"}}}
	unsafePlan, unsafeAudit := makeAudit(unsafe, domain.PlanGroundingVerdict{Pass: true, Findings: []domain.PlanGroundingFinding{}})
	result, err = p05AuthorityPostCheck(unsafePlan, simulation, evidence, unsafeAudit)
	if err != nil || result.Status != "FAIL" || len(result.Findings) == 0 {
		t.Fatalf("unsupported causal-beat erasure escaped: result=%+v err=%v", result, err)
	}

	rejecting := base
	rejecting.CausalSimulation.OffscreenStage = []domain.CharacterStageRecord{{
		Character: "九尾狐", Location: "观察点", Pressure: "浪花会抹掉潮痕",
	}}
	input, err := domain.NewActivationPlanGroundingInput(rejecting, simulation, evidence, evidence.ProtocolDigest)
	if err != nil {
		t.Fatal(err)
	}
	sourceQuote := input.Activation.Cycles[0].Decisions[0].Decision.ImmediateResult
	rejectVerdict := domain.PlanGroundingVerdict{Pass: false, Findings: []domain.PlanGroundingFinding{{
		Kind: "outcome", PlanPath: "/plan/causal_simulation/offscreen_character_stage/0/pressure",
		PlanQuote: "浪花会抹掉潮痕", SourcePath: "/activation/cycles/0/decisions/0/decision/immediate_result",
		SourceQuote: sourceQuote, Explanation: "未经权威支持的正向因果压力",
	}}}
	rejectReceipt, err := domain.FinalizePlanGroundingReceipt(input, rejectVerdict)
	if err != nil {
		t.Fatal(err)
	}
	tampered := rejecting
	tamperedReceipt := rejectReceipt
	tamperedReceipt.Verdict = domain.PlanGroundingVerdict{Pass: true, Findings: []domain.PlanGroundingFinding{}}
	tampered.GroundingReview = &tamperedReceipt
	result, err = p05AuthorityPostCheck(tampered, simulation, evidence, domain.PlanGroundingAudit{Input: input, Receipt: rejectReceipt})
	if err != nil || result.Status != "FAIL" || len(result.Findings) == 0 || result.Findings[0].Kind != "STALE_OR_MISMATCHED_RECEIPT" {
		t.Fatalf("tampered verdict with copied digests escaped: result=%+v err=%v", result, err)
	}
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

func p05AppendJSONL(path string, value any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	if _, err := file.Write(append(raw, '\n')); err != nil {
		_ = file.Close()
		return err
	}
	if err := file.Sync(); err != nil {
		_ = file.Close()
		return err
	}
	return file.Close()
}

func TestP05PlannerOnlyTraceAppendsDurably(t *testing.T) {
	path := filepath.Join(t.TempDir(), "planner-only-trace.jsonl")
	for sequence := 1; sequence <= 2; sequence++ {
		if err := p05AppendJSONL(path, agents.PlannerOnlyTraceEvent{
			Sequence: sequence, ExecutionID: "synthetic", Stage: "preflight", Result: "PASS",
		}); err != nil {
			t.Fatal(err)
		}
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	lines := bytes.Split(bytes.TrimSpace(raw), []byte{'\n'})
	if len(lines) != 2 {
		t.Fatalf("durable trace lines=%d want=2", len(lines))
	}
	for index, line := range lines {
		var event agents.PlannerOnlyTraceEvent
		if err := json.Unmarshal(line, &event); err != nil || event.Sequence != index+1 {
			t.Fatalf("trace line %d invalid: event=%+v err=%v", index, event, err)
		}
	}
}
