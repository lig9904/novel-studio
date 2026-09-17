# P0-4 Planner Grounding Root Cause and Fix

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Status: **IMPLEMENTED AND VERIFIED IN ISOLATED P0-4 WORKTREE**

## Root cause

Two independent defects occurred at different layers:

1. **DeepSeek Flash exhausted the 6,144-token output allowance in reasoning and never emitted the required tool call.** Host fail-closed behavior was correct.
2. **The model-facing pointer contract was underspecified.** The prompt listed causal fields by name while the JSON serialized them under `plan.causal_simulation`. Sol reliably found the story contradictions and used the tool, but omitted that namespace in causal-field pointers. Strict Host evidence validation correctly rejected both verdicts.

No defect was found in Story Simulation, Character Decision, World Arbitration, Readiness, receipt persistence or the Host's exact quote/pointer validator.

## Code change

- `internal/agents/plan_grounding.go`
  - states exact `/plan/contract/...` and `/plan/causal_simulation/...` namespaces;
  - includes concrete path examples;
  - requires pre-submission pointer existence and literal quote checks;
  - keeps the existing strict Host validator unchanged;
  - resolves an independent `plan_grounding` role.
- `internal/bootstrap/config.go` and `models.go`
  - register `plan_grounding` as a project model role;
  - preserve backward compatibility by inheriting `world_arbiter` when the role is absent.
- `internal/host/usage.go` and `cmd/novel-studio/pipeline_project_all_usage.go`
  - retain `plan_grounding` as its own durable/replayed accounting role instead of folding it into `world_arbiter`.
- `cmd/novel-studio/check_cmd.go`
  - includes the independently routed Grounding model in `--check` connectivity validation.
- `poc/jiujiu/project/config.json`
  - routes only `plan_grounding` to `codex/gpt-5.6-sol` with `medium` reasoning;
  - leaves Character and World Arbiter on their existing DeepSeek routes.
- protocol digest snapshot tests were updated because a deliberate model-facing protocol change must invalidate old Grounding cache identity.

## Validation

```text
go test ./...       PASS
go vet ./...        PASS
go build ./...      PASS
```

Additional checks:

- routing inheritance and explicit separation tests: PASS;
- durable Host/project-all accounting separation and CLI `--check` role coverage: PASS;
- exact prompt/tool namespace regression: PASS;
- project config load and runtime selection probe: `plan_grounding=codex/gpt-5.6-sol/medium`, `world_arbiter=deepseek/deepseek-flash/high`: PASS;
- `git diff --check`: PASS;
- config example copies match: PASS;
- fixed live Sol Medium verdict validation: PASS twice.

Independent GPT-6 Astra / High review initially found three P2 integration gaps in usage attribution, CLI model checking and report scope. All three were fixed and the follow-up review returned **PASS with no remaining P0/P1/P2**. Runtime reviewer metadata verified task `/root/reviewer`, model `gpt-6-astra`, effort `high`; the local Codex thread identifier is intentionally omitted from the public archive.

## State boundary

The controlled runs invoked only Planner Grounding Reviewer. They did not rerun Story Simulation, alter Character decisions, change World Arbitration or Readiness, write a formal plan, publish a Chapter 1 bundle, seal/promote/render, or modify global Codex configuration.

The valid Sol verdict is `pass=false`, so Chapter 1 remains blocked by real Planner content conflicts. Those findings are now actionable for a later Planner repair using the retained partial; this task did not mutate the frozen plan to force publication.
