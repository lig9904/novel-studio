# P0-5R Repair Entry Audit

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

## A — Resume the historical failed candidate

The repository has a formal sealed-convergence replan path, but it requires an active sealed-v2 frozen Plan, a valid promotion, an uncommitted chapter and an exhausted render ledger at its failure limit. P0-5R has no accepted Plan or Bundle and does not meet these preconditions.

```text
FORMAL ENTRY EXISTS: YES, for sealed convergence only
APPLICABLE TO P0-5R HISTORICAL PARTIAL: NO
CURRENT APPLICABLE FORMAL ENTRY: NONE
```

The generic `plan_details` tool can resume a partial inside an already running planning session, but no public CLI was found that safely selects this unsealed failed candidate, rebinds its current context/receipt set and resumes only Planner repair.

## B — Build a new isolated Planner candidate from the same frozen world result

The historical P0-5/P0-5R invocation copied a workspace whose current simulation checkpoint was already ready, executed Planner → Grounding and verified frozen bytes. The project-specific harness itself does not preflight that checkpoint as a hard invariant and does not disable the Project-All Character/World fallback. If the copied checkpoint is absent or not ready, `RunProjectedChapterPlanning` may run Story State again. It also resumes the copied partial rather than automatically creating a fresh candidate identity.

Generic `--pipeline --stages project-all` is not a safe B entry: if its copied checkpoint is missing or not ready, it may execute Character/World simulation.

```text
TEST HARNESS EXISTS: YES — P0-5 frozen harness
FORMAL CLI ENTRY EXISTS: NO
HARD-GUARANTEED B ENTRY: NO
HISTORICAL CONTROLLED INVOCATION: NO STORY-STATE RERUN OBSERVED
```

No new CLI or recovery authority was added in this task. A future positive Runtime must receive separate authorization. Reuse of the harness would first require an explicit ready/current checkpoint preflight and a fail-closed prohibition on Character/World fallback; otherwise a generic isolated-candidate lifecycle must be implemented and reviewed.
