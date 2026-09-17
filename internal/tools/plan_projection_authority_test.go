package tools

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPlanToolsExposeProjectionAuthorityAtFactBearingFields(t *testing.T) {
	structure, err := json.Marshal(NewPlanStructureTool(nil).Schema())
	if err != nil {
		t.Fatal(err)
	}
	details, err := json.Marshal(NewPlanDetailsTool(nil).Schema())
	if err != nil {
		t.Fatal(err)
	}
	combined := string(structure) + string(details)
	for _, required := range []string{
		"不能恢复未发生的 Soft Outline 事件",
		"不得新增事件、实体、知识或未来义务",
		"不得借候选锚点创建新资源、设备或证据",
		"事实性物件/设备/读数/状态必须来自最终simulation evidence",
		"角色已实际作出的选择；不得由Planner重选",
		"不能由对白蓝图创造",
		"不得创建新的未来Story Obligation",
		"非因果表现",
	} {
		if !strings.Contains(combined, required) {
			t.Fatalf("plan tool schema missing projection authority rule %q", required)
		}
	}
}
