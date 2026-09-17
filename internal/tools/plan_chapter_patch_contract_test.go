package tools

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/store"
)

func currentRAGExternalRow(
	receipt domain.RAGFactReceipt,
	usable string,
	transformation string,
	doNotUse string,
) map[string]any {
	return map[string]any{
		"query_or_need":         receipt.Query,
		"source_type":           "RAG",
		"source_refs":           []any{receipt.Hits[0].Ref},
		"retrieved_at":          receipt.CreatedAt,
		"freshness_requirement": "当前项目事实 receipt",
		"usable_details":        []any{usable},
		"transformation_rule":   transformation,
		"do_not_use":            []any{doNotUse},
	}
}

func executePlanDetailsPatch(t *testing.T, st *store.Store, patch map[string]any) {
	t.Helper()
	raw, err := json.Marshal(map[string]any{
		"chapter":           1,
		"causal_simulation": patch,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := NewPlanDetailsTool(st).Execute(context.Background(), raw); err != nil {
		t.Fatalf("plan_details patch: %v", err)
	}
}

func loadPartialExternalRows(t *testing.T, st *store.Store) []any {
	t.Helper()
	partial, err := st.Drafts.LoadChapterPlanPartial(1)
	if err != nil || partial == nil {
		t.Fatalf("load partial: partial=%#v err=%v", partial, err)
	}
	merged, _ := partial["causal_simulation"].(map[string]any)
	rows, _ := merged["external_reference_plan"].([]any)
	return rows
}

func newPlanDetailsPatchContractFixture(t *testing.T) (*store.Store, domain.RAGFactReceipt, map[string]any) {
	t.Helper()
	st, receipt := newPlanDetailsRAGFactReceiptFixture(t, 1)
	if _, err := NewPlanStructureTool(st).Execute(context.Background(), planStructureArgs(1)); err != nil {
		t.Fatalf("plan_structure: %v", err)
	}
	original := currentRAGExternalRow(
		receipt,
		"旧的派生用途",
		"把当前项目事实转成旧的表现方式",
		"不得复制来源摘要",
	)
	executePlanDetailsPatch(t, st, map[string]any{"external_reference_plan": []any{original}})
	return st, receipt, original
}

func TestPlanDetailsSchemaPublishesExternalReferencePatchContract(t *testing.T) {
	t.Parallel()
	s := NewPlanDetailsTool(nil).Schema()
	properties, _ := s["properties"].(map[string]any)
	causal, _ := properties["causal_simulation"].(map[string]any)
	causalProperties, _ := causal["properties"].(map[string]any)
	external, _ := causalProperties["external_reference_plan"].(map[string]any)
	if external == nil || external["type"] != "array" {
		t.Fatalf("plan_details schema does not publish external_reference_plan: %#v", causalProperties)
	}
	description, _ := external["description"].(string)
	if !strings.Contains(description, "空数组") || !strings.Contains(description, "清除") {
		t.Fatalf("external_reference_plan schema does not disclose explicit-clear semantics: %q", description)
	}
}

func TestPlanStructureResubmissionPreservesBoundCandidateWithoutPromotion(t *testing.T) {
	st := newPhaseTestStore(t)
	if _, err := NewPlanStructureTool(st).Execute(context.Background(), planStructureArgs(1)); err != nil {
		t.Fatal(err)
	}
	executePlanDetailsPatch(t, st, map[string]any{"chapter_function": "仍待Grounding审查的候选用途"})
	if _, err := NewPlanStructureTool(st).Execute(context.Background(), planStructureArgs(1)); err != nil {
		t.Fatal(err)
	}
	partial, err := st.Drafts.LoadChapterPlanPartial(1)
	if err != nil || partial == nil {
		t.Fatalf("load resubmitted partial: partial=%#v err=%v", partial, err)
	}
	merged, _ := partial["causal_simulation"].(map[string]any)
	if got := merged["chapter_function"]; got != "仍待Grounding审查的候选用途" {
		t.Fatalf("source-bound plan_structure resubmission lost staged candidate: %#v", got)
	}
	if plan, err := st.Drafts.LoadChapterPlan(1); err != nil || plan != nil {
		t.Fatalf("retained partial was promoted to formal plan: plan=%#v err=%v", plan, err)
	}
}

func TestPlanDetailsOmissionRetainsCurrentRAGFactRow(t *testing.T) {
	st, _, _ := newPlanDetailsPatchContractFixture(t)
	executePlanDetailsPatch(t, st, map[string]any{"chapter_function": "只修改另一个字段"})
	if rows := loadPartialExternalRows(t, st); len(rows) != 1 {
		t.Fatalf("omission did not retain the staged current RAG row: %#v", rows)
	}
}

func TestPlanDetailsExplicitEmptyExternalReferencePlanClearsCurrentRows(t *testing.T) {
	st, _, _ := newPlanDetailsPatchContractFixture(t)
	executePlanDetailsPatch(t, st, map[string]any{"external_reference_plan": []any{}})
	if rows := loadPartialExternalRows(t, st); len(rows) != 0 {
		t.Fatalf("explicit empty external_reference_plan restored rows instead of clearing them: %#v", rows)
	}
}

func TestPlanDetailsNullExternalReferencePlanIsNotExplicitRemoval(t *testing.T) {
	st, _, _ := newPlanDetailsPatchContractFixture(t)
	executePlanDetailsPatch(t, st, map[string]any{"external_reference_plan": nil})
	if rows := loadPartialExternalRows(t, st); len(rows) != 1 {
		t.Fatalf("schema-invalid null unexpectedly acted as explicit removal: %#v", rows)
	}
}

func TestPlanDetailsCurrentRAGRowUsesCompleteIdentity(t *testing.T) {
	t.Run("identical row deduplicates", func(t *testing.T) {
		st, _, original := newPlanDetailsPatchContractFixture(t)
		executePlanDetailsPatch(t, st, map[string]any{"external_reference_plan": []any{original}})
		if rows := loadPartialExternalRows(t, st); len(rows) != 1 {
			t.Fatalf("identical current RAG row was duplicated: %#v", rows)
		}
	})

	t.Run("revised usable details remain a distinct purpose", func(t *testing.T) {
		st, receipt, _ := newPlanDetailsPatchContractFixture(t)
		revised := currentRAGExternalRow(receipt, "修订后的派生用途", "把当前项目事实转成旧的表现方式", "不得复制来源摘要")
		executePlanDetailsPatch(t, st, map[string]any{"external_reference_plan": []any{revised}})
		if rows := loadPartialExternalRows(t, st); len(rows) != 2 {
			t.Fatalf("different usable_details did not retain distinct full identities: %#v", rows)
		}
	})

	t.Run("revised transformation remains a distinct purpose", func(t *testing.T) {
		st, receipt, _ := newPlanDetailsPatchContractFixture(t)
		revised := currentRAGExternalRow(receipt, "旧的派生用途", "把当前项目事实转成修订后的表现方式", "不得复制来源摘要")
		executePlanDetailsPatch(t, st, map[string]any{"external_reference_plan": []any{revised}})
		if rows := loadPartialExternalRows(t, st); len(rows) != 2 {
			t.Fatalf("different transformation_rule did not retain distinct full identities: %#v", rows)
		}
	})
}

func TestCurrentRAGFactExternalRowKeyUsesCompleteAuthoredIdentity(t *testing.T) {
	_, receipt := newPlanDetailsRAGFactReceiptFixture(t, 2)
	exactOrder := []string{receipt.Hits[0].Ref, receipt.Hits[1].Ref}
	exactRefs := map[string]struct{}{receipt.Hits[0].Ref: {}, receipt.Hits[1].Ref: {}}
	base := currentRAGExternalRow(receipt, "用途一", "转换一", "禁用一")
	base["source_refs"] = []any{receipt.Hits[1].Ref, receipt.Hits[0].Ref}
	baseKey, ok := currentRAGFactExternalRowKey(base, exactRefs, exactOrder)
	if !ok {
		t.Fatalf("complete current RAG row did not produce an identity: %#v", base)
	}

	for name, mutate := range map[string]func(map[string]any){
		"query":          func(row map[string]any) { row["query_or_need"] = "用途二" },
		"usable":         func(row map[string]any) { row["usable_details"] = []any{"用途二"} },
		"transformation": func(row map[string]any) { row["transformation_rule"] = "转换二" },
		"do_not_use":     func(row map[string]any) { row["do_not_use"] = []any{"禁用二"} },
	} {
		t.Run(name, func(t *testing.T) {
			row := map[string]any{}
			for key, value := range base {
				row[key] = value
			}
			mutate(row)
			key, ok := currentRAGFactExternalRowKey(row, exactRefs, exactOrder)
			if !ok || key == baseKey {
				t.Fatalf("%s was not part of the complete authored identity: base=%q changed=%q", name, baseKey, key)
			}
		})
	}

	metadataOnly := map[string]any{}
	for key, value := range base {
		metadataOnly[key] = value
	}
	metadataOnly["retrieved_at"] = "2099-01-01T00:00:00Z"
	metadataOnly["freshness_requirement"] = "服务端会重新物化"
	metadataOnly["source_type"] = "rag_fact"
	metadataKey, ok := currentRAGFactExternalRowKey(metadataOnly, exactRefs, exactOrder)
	if !ok || metadataKey != baseKey {
		t.Fatalf("server metadata changed authored identity: base=%q changed=%q", baseKey, metadataKey)
	}
}

func TestPlanDetailsClearThenReplaceCurrentRAGRow(t *testing.T) {
	st, receipt, _ := newPlanDetailsPatchContractFixture(t)
	executePlanDetailsPatch(t, st, map[string]any{"external_reference_plan": []any{}})
	revised := currentRAGExternalRow(receipt, "修订后的唯一用途", "只形成合法表现细节", "不得新增故事事实")
	executePlanDetailsPatch(t, st, map[string]any{"external_reference_plan": []any{revised}})
	rows := loadPartialExternalRows(t, st)
	if len(rows) != 1 {
		t.Fatalf("clear-then-replace did not leave exactly one current row: %#v", rows)
	}
	row, _ := rows[0].(map[string]any)
	usable := stringSliceFromAny(row["usable_details"])
	if len(usable) != 1 || usable[0] != "修订后的唯一用途" {
		t.Fatalf("clear-then-replace retained stale authored content: %#v", rows)
	}
}

func TestPlanDetailsTraceRecordsBoundedPatchStages(t *testing.T) {
	st := newPhaseTestStore(t)
	if _, err := NewPlanStructureTool(st).Execute(context.Background(), planStructureArgs(1)); err != nil {
		t.Fatal(err)
	}
	events := []PlanDetailsTraceEvent{}
	tool := NewPlanDetailsTool(st).WithTraceRecorder(func(event PlanDetailsTraceEvent) error {
		events = append(events, event)
		return nil
	})
	args, _ := json.Marshal(map[string]any{
		"chapter": 1,
		"causal_simulation": map[string]any{
			"chapter_function": "合成trace验证",
		},
	})
	if _, err := tool.Execute(context.Background(), args); err != nil {
		t.Fatal(err)
	}
	want := []string{"before_merge", "after_merge", "after_source_anchor", "persisted", "validation"}
	if len(events) != len(want) {
		t.Fatalf("trace event count=%d want=%d: %+v", len(events), len(want), events)
	}
	for index, phase := range want {
		if events[index].Sequence != index+1 || events[index].Phase != phase {
			t.Fatalf("trace[%d]=%+v want phase=%s", index, events[index], phase)
		}
	}
	if len(events[0].PatchKeys) != 1 || events[0].PatchKeys[0] != "chapter_function" ||
		events[0].PatchDigest == "" || events[0].PartialDigest == "" ||
		events[1].StateDigest == "" || events[2].StateDigest == "" ||
		events[3].PartialDigest == "" || events[4].Result != "PASS" {
		t.Fatalf("trace omitted sanitized patch or state evidence: %+v", events)
	}
}
