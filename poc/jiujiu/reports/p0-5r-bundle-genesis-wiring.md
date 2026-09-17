# P0-5R Bundle Genesis Wiring

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

## Production authority

The production Project-All path loads current bundles, calls `pipelineProjectAllTail`, and passes its two outputs to `buildPipelineProjectedChapterBundle`:

```text
no previous bundle
  → DeriveProjectedChainGenesisV2(generation)
  → generation.BaseStateRoot

existing bundles
  → latest bundle.BundleDigest
  → latest bundle.ProjectedPostStateRoot
```

The bundle validator requires a syntactically valid predecessor digest. The chain validator then proves that the first predecessor equals the exact generation genesis, or that a later predecessor equals the previous bundle digest.

## Controlled harness defect and fix

The P0-5 controlled harness passed `""` directly. It now calls the same `pipelineProjectAllTail(generation, nil)` helper used by production and passes both returned authority values to the bundle builder.

It does not hard-code a hash, skip validation, weaken Chapter 1 rules or create a P0-specific genesis algorithm. The change is confined to the controlled P0 harness; production code already had the correct behavior.

```text
BUNDLE GENESIS HARNESS WIRING: PASS
PRODUCTION GENERIC FIX: NO FIX REQUIRED
MODEL-RUNTIME BUNDLE CREATED IN THIS TASK: NO
```
