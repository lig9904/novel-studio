# P0-5R Historical Partial / RAG Repair Summary

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

```text
HISTORICAL CAUSAL ATTRIBUTION: INCONCLUSIVE
REPOSITORY BEHAVIOR REPRODUCTION: REPRODUCED
GENERIC CONTRACT DEFECT: CONFIRMED
FIX: VERIFIED

A ENTRY: FORMAL SEALED REPLAN EXISTS, NOT APPLICABLE; NO FORMAL ENTRY FOR THIS UNSEALED PARTIAL
B ENTRY: HISTORICAL TEST INVOCATION ONLY; NO HARD-GUARANTEED ENTRY OR FORMAL CLI

PAID MODEL RUNTIME IN THIS DIAGNOSTIC TASK = NO
STORY SIMULATION RERUN IN THIS DIAGNOSTIC TASK = NO
CHAPTER 1 POSITIVE RUNTIME IN THIS DIAGNOSTIC TASK = NOT RUN
CORE SHA: 99eec8bf53f098e0d9e61f492a84d320aae33977
```

The confirmed defect was the absence of an explicit removal operation for eligible current-receipt RAG fact rows: even `external_reference_plan: []` restored them after merge. The focused model-facing schema also omitted this field and its patch semantics.

The generic fix makes an empty array a deliberate clear operation, preserves omission/resume behavior, keeps valid non-empty cross-batch RAG progress, publishes the complete row schema and documents clear-then-replace repair. Grounding, simulation authority, Proposal isolation and Canon authority were not weakened.

Historical P0-5R causation remains inconclusive. Four retained Grounding snapshots show unsupported fields surviving across reviews, but the raw Planner patch and merge/source-anchor intermediates were not retained. The synthetic reproduction proves current code behavior and the fix; it cannot prove what the historical Planner submitted.

P0-5R remains `BLOCKED`, model capability remains `INCONCLUSIVE`, and Chapter 1 Bundle remains uncreated.
