# P0-5 Simulation → Plan Projection Authority Summary

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

```text
P0-5: BLOCKED
CHAPTER 1 BUNDLE: NOT CREATED
WAITING FOR HUMAN REVIEW
```

## Final answers

1. **Planner Authority** — Planner may select, order, group, emphasize, express and organize authoritative Story Facts. It may add removable non-causal sensory/atmospheric rendering. It cannot decide what actually happened or create causal-persistent facts.
2. **Story Fact / Presentation Detail** — A detail is fact-bearing when it can persist, be held/used/measured/tracked, prove something, transfer knowledge, trigger action, change an outcome or create a future obligation. If deleting it can change later simulation, it requires authority. Removable wording, visualization, sensory texture and atmosphere are presentation.
3. **Creative capacity** — Preserved. Planner still controls scene grouping, pacing, emphasis, viewpoint and expression, and may combine multiple real facts in a scene without inventing causality.
4. **Unsupported facts** — `浮标绳`, `水尺` and `现场来客` originated in Soft Outline and were re-promoted after simulation left them unrealized. Planner prompt/schema authority was under-specified. The new Planner contract forbids this; current Grounding also rejects rope/gauge facts hidden in positive RAG fields.
5. **Grounding strictness** — Preserved and strengthened. It checks every positive Drafter-consumed channel, excludes negative/source metadata from occurrence claims, and stale passing receipts cannot be reused.
6. **P0-2/P0-3/P0-4 regressions** — None found in code/static regression or frozen digests. Full `go test ./...`, `go vet ./...`, `go build ./...` and independent review passed. P0-4's prior runtime evidence remains historical; the current protocol has new identity and direct negative evidence.
7. **Bounded runtime** — No. One normal Planner run and one confirmation were used. The confirmation produced an unsafe plan that only passed the old Grounding protocol; current Grounding returns `pass=false` with five exact findings.
8. **Chapter 1 Bundle** — Not created. No current-protocol `pass=true` receipt exists.
9. **New P0/P1** — No P0. Runtime found two P1-class protocol gaps: positive side-channel Grounding coverage and stale Grounding receipt reuse. Both are fixed, tested and Astra-reviewed. Runtime acceptance remains blocked by the run cap.

## Root Cause and fix

The primary Root Cause is a model-facing Planner authority contract gap, reinforced by fact-bearing tool descriptions. Existing Domain authority was sufficient, so no enum, migration or parallel evidence ledger was added. Runtime then exposed and closed the two Grounding/Host gaps above.

The implementation is generic, has no 九九 or model-specific production branch, does not modify Character Decision, World Arbitration, Readiness, Canon or Story Simulation, and does not weaken evidence requirements.

## Verification boundary

Code and protocol review: **PASS**.

Current-protocol rejection of the known unsafe plan: **PASS**.

Positive runtime acceptance and bundle creation: **BLOCKED**.

## MODEL ROUTING

```text
Task Level:
LEVEL 2 — Complex / Core Engineering

Primary:
/root — GPT-5.6 Sol / High (observed runtime)

Delegated:
/root/explorer — GPT-5.6 Terra / Medium (observed runtime)
/root/worker — GPT-5.6 Luna / Medium (observed runtime)

Escalations:
LEVEL 4 independent review because this changes fail-closed Authority and core planning protocol behavior.

Reviewer:
/root/reviewer — GPT-6 Astra / High (observed runtime), final review PASS, no remaining P0/P1/P2.

Reason:
Root owned Authority semantics and integration; Explorer traced the call path; Worker indexed frozen evidence; Astra independently searched for authority leaks, false passes and test blind spots.
```

## Git / publication state

Work is isolated on branch `codex/p0-5-plan-projection-authority` in its own worktree. Changes remain uncommitted because runtime acceptance did not pass. No global Codex configuration was changed and no PR was created.
