# P0-5R Repair Lineage

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Baseline: `9b2e92cecc86c1cb81e09fa13785fb73f05b3932`

## Evidence availability

| Node | Status |
|---|---|
| Historical partial before Sol run | RETAINED in the isolated copied workspace |
| Raw Planner `plan_structure` / `plan_details` arguments | NOT RETAINED |
| Per-call Planner patch payload | NOT RETAINED |
| Per-call merge intermediate | NOT RETAINED |
| Per-call source-anchor intermediate | NOT RETAINED |
| Host-valid Grounding input + receipt snapshots | RETAINED, 4 |
| Invalid Grounding raw response | NOT RETAINED |
| Final partial | RETAINED |
| Interrupted final Grounding call | STARTED WAL ONLY; response not retained |

The four retained Grounding files contain `{input, receipt}`. They are not raw Planner transcripts. Their hashes and receipt findings can prove the Plan state reviewed at those boundaries, but cannot prove which fields the Planner omitted, replaced or attempted to delete.

## Partial chronological order

The preserved original workspace gives this filesystem-mtime order:

| Order | Input digest suffix | Local mtime (+08:00) | Verdict |
|---:|---|---|---|
| 1 | `ecdb60a5...` | 2026-09-17 15:35:24 | `pass=false`, 4 findings |
| 2 | `9b2e8c00...` | 2026-09-17 15:37:12 | `pass=false`, 3 findings |
| 3 | `6f4a4db2...` | 2026-09-17 15:39:59 | `pass=false`, 4 findings |
| 4 | `9314d546...` | 2026-09-17 15:43:04 | `pass=false`, 8 findings |

The final partial and usage files were updated at approximately 15:43:42. Filesystem mtime is local retained metadata, not embedded receipt authority and not Git-portable. It proves a useful partial order only; the Grounding filename order and hash order are not chronology.

## Representative fields

| Field | Retained evidence | Attribution |
|---|---|---|
| `goal` | All four Grounding snapshots match each other; final partial differs | A later state change is visible, but raw patch/merge/source-anchor nodes are unavailable |
| `hook` | First three snapshots match; fourth contains the supported two-observation hook; final matches the fourth | Final state adopted the corrected hook; exact patch is unavailable |
| `required_beats` | First three retain buoy-rope/gauge/visitor claims; fourth is corrected; final matches the fourth | Final state adopted corrected beats; exact patch is unavailable |
| `ending_consequence_contract` | All four snapshots and final are identical | No evidence that the Planner submitted a correction |
| `emotional_logic` | All four snapshots and final are identical | No evidence that the Planner submitted a correction |
| `reader_entertainment_plan` | All snapshots match; final representation differs | A state/representation change is visible; its author and stage are not retained |
| `reader_reward_plan` | All four snapshots and final are identical | No evidence that the Planner submitted a correction |
| `longform_opening` | All four snapshots and final are identical | No evidence that the Planner submitted a correction |
| `external_reference_plan` | All four snapshots and final are identical | No evidence of an actual historical delete/replace patch; Host restoration is not proven |

The machine-readable hashes and partial order are in `p0-5r-repair-runtime-evidence/historical-lineage.json`.

## Historical conclusion

```text
HISTORICAL CAUSAL ATTRIBUTION: INCONCLUSIVE
```

Current code behavior can be reproduced synthetically, but the missing historical Planner patches prevent binding that behavior to the P0-5R timeout. No model response was reconstructed or invented.
