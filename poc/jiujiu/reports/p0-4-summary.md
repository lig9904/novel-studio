# P0-4 Planner Grounding Summary

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Status: **DIAGNOSIS PASS / CORE FIX PASS / ROUTING PASS / CHAPTER 1 STILL UNPUBLISHED**

Branch: `codex/p0-4-planner-grounding`

Base core commit: `1fb49d33d58a2137679fbf4c8f9e280c2bfa532a`

P0-2 evidence commit: `5850b04023812a200b0c4c0bbdee4962f14ec82d`

## MODEL ROUTING

```text
Task Level: LEVEL 2 core protocol diagnosis and repair
Primary: /root — GPT-5.6 Sol / High
Delegated: /root/explorer — GPT-5.6 Terra / Medium
Delegated: /root/worker — GPT-5.6 Luna / Medium
Reviewer: /root/reviewer — GPT-6 Astra / High
Escalations: LEVEL 4 independent review after core protocol/routing change
Reason: preserve root ownership while separating call-path exploration, evidence indexing and independent risk review.
```

## Answer to the primary question

Planner Grounding was blocked by **two separate failure layers**:

1. **Model capability/tool compliance:** DeepSeek Flash used the complete 6,144-token output allowance as reasoning and returned zero tool calls on the frozen 90,311-byte activation packet.
2. **Prompt/schema protocol usability:** Sol Medium returned exactly one tool call twice under the original protocol, but reliably omitted the serialized `/plan/causal_simulation/...` namespace for causal fields. The strict Host correctly rejected both verdicts.

Therefore this was neither a pure routing-only Case A nor a Host-validator defect. The allowed Hypothesis B phase produced a narrow prompt/schema correction; after that correction Sol Medium produced Host-valid structured verdicts twice.

## Business inputs and outputs

Input:

- the retained P0-2 Planner partial;
- the same Chapter 1 simulation;
- the same two-cycle Character Activation evidence;
- the same Character decisions, World Arbitration, Readiness, Canon and Soft Event facts.

Output:

- read-only path report: `poc/jiujiu/reports/p0-4-grounding-current-path.md`;
- immutable comparison fixture: `poc/jiujiu/snapshots/p0-4-grounding-fixture/`;
- model comparison: `poc/jiujiu/reports/p0-4-grounding-model-comparison.md`;
- root cause and fix report: `poc/jiujiu/reports/p0-4-grounding-root-cause-and-fix.md`;
- raw per-run structured evidence: `poc/jiujiu/reports/p0-4-grounding-comparison/*.json`;
- independently routed `plan_grounding` project role using `codex/gpt-5.6-sol` at `medium` reasoning.

## Interface and collaboration behavior

- The Reviewer still receives one exact activation packet and one tool.
- The Host still requires exactly one `submit_plan_grounding_verdict` call and performs strict JSON, path, quote and evidence binding checks.
- Missing `plan_grounding` configuration inherits `world_arbiter` for backward compatibility.
- Explicit `plan_grounding` configuration does not move Character or World Arbiter models.
- Durable usage, project-all accounting, replay and CLI `--check` retain the independent Grounding role.

## Acceptance evidence

| Check | Result |
|---|---|
| DeepSeek original baseline reproduced | FAIL as expected: 0 tool calls |
| Sol Medium original protocol run 1 | FAIL: causal path namespace omitted |
| Sol Medium original protocol run 2 | FAIL: same defect class |
| Sol Medium fixed protocol run 1 | PASS: one Host-valid verdict |
| Sol Medium fixed protocol run 2 | PASS: one Host-valid verdict |
| Story Facts equal across fix validation | PASS; identical excluding model-bound `review_protocol` |
| Fixture checksum verification | PASS |
| `go test ./...` | PASS |
| `go vet ./...` | PASS |
| `go build ./...` | PASS |
| Independent Astra/High review | PASS; no remaining P0/P1/P2 |
| Global Codex configuration modified | NO |
| Story Simulation rerun | NO |

## Current downstream state

Both fixed Sol verdicts are structurally valid `pass=false` classifications. They identify real Planner additions not supported by frozen Story Facts. No Chapter 1 plan was rewritten or finalized in this task, and no bundle was published. Character Decision, World Arbitration, Readiness and P0-2/P0-3 evidence remain frozen.
