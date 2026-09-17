# P0-5R Final Runtime Host Defect

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Status: **CONFIRMED CONTROLLED-HARNESS WIRING BLOCKER — NOT FIXED IN THIS TASK**

The controlled attempt created a formal grounded Plan, then failed in Chapter 1 bundle construction. The P0-5 harness called `buildPipelineProjectedChapterBundle` with an empty `previousBundleDigest`. The bundle constructor copied that empty value into `ProjectedChapterBundle.PreviousBundleDigest`, and `ValidateProjectedChapterBundle` rejected it because digest fields must use `sha256:<64 lowercase hex>` form.

The chain validator separately derives the generation genesis and requires the first bundle's `previous_bundle_digest` to equal it. The production Project-All caller already calls `pipelineProjectAllTail`, which returns the derived genesis when no earlier bundle exists. The newly exposed defect is therefore in the P0-5 controlled Runtime harness wiring, not in the production tail calculation or validator.

No repair was made. The task contract prohibits:

```text
run → fail → edit code → rerun
```

The defect requires a separate deterministic harness-wiring task and a regression that exercises first-bundle genesis binding before any future paid Runtime.
