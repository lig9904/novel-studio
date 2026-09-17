# P0-5R Partial and RAG Patch Semantics

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

## Confirmed pre-fix behavior

- Re-submitting `plan_structure` retained the staged `causal_simulation` only while the existing structure remained bound to the current simulation/rewrite sources. The retained value remained a partial candidate; it did not become a formal Plan or Story Fact.
- Omitted `causal_simulation` keys retained their prior values.
- Most submitted keys used replacement semantics. `context_sources`, `review_refinement`, and character-keyed arrays had explicit merge exceptions.
- Merge was followed by source-anchor processing, which normalized current receipts, rebound simulation-derived fields, removed stale receipt tokens and could rehome or preserve eligible current-receipt fact rows.
- A complete current-receipt RAG fact row was identified by `query_or_need + ordered source_refs + usable_details + transformation_rule + do_not_use`. Receipt metadata was not part of that authored identity.
- A schema-valid `external_reference_plan: []` was replaced with an empty array during merge, then the old eligible fact rows were restored by source-anchor processing. There was therefore no public explicit-clear operation.
- The focused `plan_details` schema did not publish `external_reference_plan`, although runtime guidance and Host processing used it.

The fixed baseline plus the new tests reproduced three failures: missing model-facing schema, empty-array restoration, and stale/corrected coexistence after an ineffective clear.

## Generic contract defect

```text
GENERIC CONTRACT DEFECT: CONFIRMED
```

The public tool said same-name fields were overwritten, while an eligible RAG fact row survived even an explicit empty-array patch. This made omission and deliberate removal indistinguishable and prevented a bounded repair from clearing rejected authored content without manipulating files outside the tool contract.

This is a generic partial-update defect. It is independent of 九九 terms, model brand and any specific unsupported fact.

## Fix

- Omission continues to retain staged progress.
- A non-empty `external_reference_plan` batch continues to preserve eligible current-receipt fact rows, so later craft-only batches do not lose valid RAG progress.
- An explicit empty array now clears the field and skips staged fact-row restoration for that patch.
- A corrected row can then be submitted in the next batch and rebound to the current receipt.
- `null` remains outside the public schema and is not treated as an explicit removal request.
- The focused schema now exposes the full external-reference row and documents omission, non-empty preservation, full identity, empty-array clearing and clear-then-replace repair.
- Same refs with different legitimate `usable_details`, `transformation_rule`, query or `do_not_use` remain distinct. The fix does not collapse rows by refs alone.
- An identical complete identity is deduplicated.

## Deterministic matrix

| Case | Result |
|---|---|
| Source-bound `plan_structure` resubmission | Partial retained; no formal Plan promotion |
| Omitted `external_reference_plan` | Retained |
| Explicit empty array | Cleared |
| Direct raw `null` | Not explicit removal; outside schema |
| Same refs, changed `usable_details` | Distinct identities coexist |
| Same refs, changed `transformation_rule` | Distinct identities coexist |
| Same refs, multiple legitimate purposes | Preserved as distinct full identities |
| Identical complete identity | Deduplicated |
| Stale/mixed/incomplete row | Not preserved |
| Alias before normalization | Fails closed unless unambiguous current alias normalization applies |
| Craft-only replacement | Current eligible fact progress retained; craft replacement remains replacement |
| Empty clear followed by corrected row | Exactly one corrected row remains |
| Stale Grounding receipt | Cannot pass a changed/current input digest |
| Failed finalize/resume | Partial survives and can receive later batches |

These are synthetic repository tests, not reconstructed historical P0-5R patches.
