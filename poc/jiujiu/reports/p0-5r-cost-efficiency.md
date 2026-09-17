# P0-5R Cost and Efficiency

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

| Metric | Historical DeepSeek baseline | Sol Medium Run 1 |
|---|---:|---:|
| Top-level runs | 2 | 1 |
| Planner calls | 14 | 7 completed |
| Grounding calls | 9 | 6 started / 5 completed |
| Input tokens | 1,030,223 | 484,522 completed |
| Output tokens | 98,537 | 19,047 completed |
| Cache read | 657,280 | 20,480 completed |
| Total tokens | 1,128,760 | 503,569 completed |
| Elapsed | 887.773 s | 601.633 s |
| Legal plans | 0 | 0 |
| Accepted bundles | 0 | 0 |

```text
ACTUAL COST: UNAVAILABLE
TOKENS PER ACCEPTED BUNDLE: N/A — NO ACCEPTED BUNDLE
TIME PER ACCEPTED BUNDLE: N/A — NO ACCEPTED BUNDLE
MODEL CALLS PER ACCEPTED BUNDLE: N/A — NO ACCEPTED BUNDLE
```

The WAL contains configured accounting values for completed Sol calls, but those are not treated as a provider invoice or actual billed cost. The interrupted final Grounding call has no persisted token receipt.
