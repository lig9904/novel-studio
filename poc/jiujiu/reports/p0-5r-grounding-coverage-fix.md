# P0-5R Grounding Coverage Fix Decision

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

```text
GROUNDING COVERAGE CODE CHANGE: NO CODE CHANGE REQUIRED
```

The full `ChapterPlan` is already included in `PlanGroundingInput.Plan`; activation mode adds the verified per-cycle decision trace and POV before/after/new-knowledge state. It does not expose every actor's original observation or create per-leaf source bindings. The fox claim and the retained activation evidence needed to identify its unsupported fragment were model-visible. Host receipt validation accepts exact `/plan/causal_simulation/offscreen_character_stage/...` findings.

`SourceBindingV2` is built after Grounding and is correctly retained as a bundle-time postcondition. Its current per-entry transformation bindings cover external references, while RAG/craft are receipt-level; it does not independently bind every `grounding_details` or `reality_support_plan` leaf. Feeding it backward as a new Grounding authority would create another protocol layer and would not address the observed probabilistic false negative.

Because the issue is `REVIEWER JUDGMENT FAILURE`, the task's authorization condition for a coverage repair—`FACT_BEARING + UNSUPPORTED + NO AUTHORITY COVERAGE`—is not met. No prompt, schema, source-binding or Grounding implementation was changed.

This decision does not declare the unsupported claim acceptable. It keeps the deterministic coverage question separate from future model-quality/stability evaluation.
