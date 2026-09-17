# P0-5R Sol Medium Run 1

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Result: **FAIL — no valid Plan before harness timeout; terminal Planner capability was not observed**

Primary classification: **E — harness global timeout; Planner max-turn was not observed**

Contributing classifications: **D — legitimate Grounding rejection**, **A — one invalid structured Grounding verdict rejected by Host**, **H — recovery from historical partial**.

| Metric | Result |
|---|---:|
| Planner model | GPT-5.6 Sol / Medium |
| Grounding model | GPT-5.6 Sol / Medium |
| Elapsed | 601.633 s |
| Planner calls | 7 started / 7 completed |
| Grounding calls | 6 started / 5 completed |
| Completed input tokens | 484,522 |
| Completed output tokens | 19,047 |
| Completed cache read | 20,480 |
| Completed total tokens | 503,569 |
| Current-protocol Host-valid verdicts | 4 |
| Grounding pass=true | 0 |
| Formal plan | NO |
| Bundle | NO |
| Actual cost | UNAVAILABLE |

All four persisted current-protocol verdicts were structured and `pass=false`, with finding counts 4, 8, 3 and 4. One Grounding response cited an unverifiable excerpt and was correctly rejected rather than persisted. The final Grounding attempt was interrupted by Go's 10-minute test deadline.

The final partial had repaired its structure/required beats but still contained 27 keyword matches across 19 distinct positive-field paths spanning ending, entertainment, reward, longform and external-reference fields. No stale passing receipt was reused.
