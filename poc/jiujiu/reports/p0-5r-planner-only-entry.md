# P0-5R Planner-only Controlled Recovery Entry

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Baseline: `37dd8ac59ac3edb1d264ce65388aed9cb58c5225`

## Entry

`RunPlannerOnlyProjectedChapterPlanning` is an explicit program entry over the existing Project-All runner. Default `RunProjectedChapterPlanning` keeps its existing Character/World fallback behavior. The controlled entry accepts a content-addressed `PlannerOnlyRecoveryContract` prepared from the isolated workspace and never impersonates sealed-convergence replan.

Execution order:

```text
short preparation lock
  → capture generation/chapter/context/protocol/simulation/checkpoint/evidence/partial digests
project-all execution lock bound to execution_id + unique attempt nonce
  → strict preflight before ModelSet/provider work
  → no embedding initialization
  → reuse exact ready simulation and validated activation/readiness evidence
  → ensure existing local craft receipt relation
  → revalidate the full recovery contract immediately before provider work
  → Host novel_context(planning)
  → production Planner loop with novel_context + plan_structure + plan_details only
  → current Grounding inside plan_details
  → existing formal-plan/checkpoint/receipt/render-context gates
```

The strict preflight requires:

- exact generation, chapter and projected planning-context digest;
- exact current `chapter_world_simulation` checkpoint bytes, sequence and digest;
- `ChapterWorldSimulationStatus(...).ready=true`;
- exact simulation ID and deterministic simulation digest;
- exact project-all context source token in simulation sources;
- complete Character Activation/Readiness evidence, or complete legacy Character evidence;
- an unfinalized Planner partial still bound to the current simulation/rewrite sources;
- exact partial digest and current Planner protocol digest;
- no existing formal Plan.

Missing or stale inputs return a precondition error before any model call. The Planner-only branch receives the validated simulation directly, so neither Character Agent fallback nor the generic World Simulator loop can become reachable. The mode also omits the RAG embedder and `craft_recall` model tool; it retains Host context, current local receipts, accounting, PlanStructure, PlanDetails and Grounding.

Each call uses a unique lock owner even when two callers present the same recovery contract. Concurrent attempts therefore cannot treat the shared execution ID as lock reentry; only one reaches the Planner.

## Applicable state

```text
APPLICABLE:
  isolated copy + ready current simulation/checkpoint + valid evidence
  + exact context/generation/protocol binding + source-bound historical partial

NOT APPLICABLE:
  missing/not-ready/stale simulation; wrong generation/chapter/context;
  evidence mismatch; stale partial; formal Plan already exists;
  fresh-candidate creation; sealed/render recovery; production publication
```

The P0-5 opt-in harness now prepares the contract and calls this entry. It writes the contract and bounded trace only when separately authorized to run. No CLI or GUI was added.
