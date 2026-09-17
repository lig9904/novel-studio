package agents

import (
	"strings"
	"testing"

	"github.com/chenhongyang/novel-studio/assets"
)

func TestPlannerProjectionAuthorityContractCoversRequiredSemanticCases(t *testing.T) {
	contract := assets.Load("").Prompts.Planner + projectAllPlannerBoundary + projectAllActivationPlannerBoundary
	for _, tc := range []struct {
		name  string
		parts []string
	}{
		{"A_non_causal_water_ripple_allowed", []string{"可添加删除后不会改变任何后续 Story Simulation 结果的感官/氛围表现", "不得持续为实体"}},
		{"B_unsourced_buoy_rope_rejected", []string{"新资源/道具/设备", "必须能回指"}},
		{"C_unsourced_numeric_gauge_rejected", []string{"测量", "需要 authority source"}},
		{"D_unsourced_visitor_rejected", []string{"新角色/现场人物", "新会面/通信/事件"}},
		{"E_unknown_identity_knowledge_rejected", []string{"新消息/线索/知识", "角色知识/状态"}},
		{"F_two_real_facts_may_share_scene", []string{"两个或多个真实 Story Facts 可以合并进同一 Scene", "不能新增它们之间未被裁决的因果关系"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, part := range tc.parts {
				if !strings.Contains(contract, part) {
					t.Fatalf("projection authority contract missing %q", part)
				}
			}
		})
	}
	for _, required := range []string{"Soft Outline", "如果删掉它，后续 Story Simulation 是否可能得到不同结果", "只修 Planner", "不得续推世界"} {
		if !strings.Contains(contract, required) {
			t.Fatalf("projection contract missing fail-closed rule %q", required)
		}
	}
	if !strings.Contains(contract, "RAG/craft/web 来源不能绕过最终 simulation") ||
		!strings.Contains(contract, "不能成为本章新事件、人物、资源、设备、通信、知识迁移或结果的替代 Authority") {
		t.Fatal("external sources can still bypass Story Simulation authority")
	}
	if !strings.Contains(contract, "黄金三章是质量目标，不是 Story Authority") ||
		!strings.Contains(contract, "不得为了章序模板补造能力使用") {
		t.Fatal("chapter-order quality template can still create Story Facts")
	}
}
