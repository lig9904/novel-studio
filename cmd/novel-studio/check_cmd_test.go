package main

import "testing"

func TestCheckRolesIncludeDedicatedPlanGroundingModel(t *testing.T) {
	for _, role := range checkRoles {
		if role == "plan_grounding" {
			return
		}
	}
	t.Fatal("--check omits the independently routed plan_grounding model")
}
