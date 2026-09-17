# P0-5 Runtime

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Status: **BLOCKED**

Chapter 1 Bundle: **NOT CREATED**

## Frozen input

Runtime used the isolated copy at `poc/jiujiu/project/test-scaffold-p0-5-run` and the existing Chapter 1 simulation for generation `pg2_b2354a98e8714bba2baca89a`.

The following frozen files remained byte-identical:

| Evidence | SHA-256 |
|---|---|
| Chapter simulation | `5471807b513f333d9d3a2675e870f1e3d96498dcf63c140411eaa93b7d99f15e` |
| Character activation evidence | `fdbefdc0c001f18e0e15d448ab9fe1f5d8ac60487e82081b833dbc34d4fce4fa` |
| Soft Outline | `70b08086a81e1d2b04302a2332dc98c79f53aaec9dd7b4c81b5d2099e71f1849` |
| Layered Outline | `eeb73b94031e59ecac17d0dd70794cbd6aa1e3467f808f21672fb844fb733f55` |

Character Agent, World Arbiter and Readiness were not rerun. Planner remained on its existing DeepSeek Flash configuration. Planner Grounding used GPT-5.6 Sol / Medium.

## Bounded top-level runs

| Run | Scope | Result |
|---|---|---|
| Normal | Planner → Grounding repair loop, frozen simulation | Ended after about 555.953 s during the sixth Grounding step; no formal plan and no bundle |
| Confirmation | Same frozen inputs and configuration | Ended after about 331.820 s with a formal plan and historical old-protocol `pass=true`; a positive-field check found unsupported RAG-derived facts, so bundle construction stopped |

The allowed one normal run plus one confirmation were exhausted. No further Planner sampling occurred.

## Historical pass was invalidated

The saved plan's old receipt used:

```text
review_protocol = sha256:fce07afd0071cfde9afbf9467a7087018a4ca0a7f4aabf17766035d60cb5a9f9
input_digest    = sha256:6a822d41e309ab8cefa27e11e622d11084aa482c90b666f630bd4adb912ef5fb
pass            = true
findings        = 0
```

That plan still placed unsupported facts in positive `external_reference_plan` fields. The receipt is historical evidence only and cannot be reused by the current Host because the review protocol no longer matches.

## Current-protocol negative regression

A Grounding-only replay used the same saved plan and frozen evidence. It did not run Planner or any story-state stage.

```text
review_protocol = sha256:3f8e557cabb1ec2f832e03e3b2205dd71408e1e76ac893b86b849014976d7abb
structured verdicts = 1
pass = false
findings = 5
regression status = PASS
```

The findings identify:

1. an unsupported claim that two observation periods were continuous;
2. invented tide-return time pressure;
3. water gauge, retreat channel and tide window in positive usable details;
4. buoy rope and a channel-state claim in positive usable details;
5. rope-knot and cross-observation distance claims in a positive transformation rule.

Every finding contains an existing JSON pointer and exact quote. This proves the strengthened Grounding gate rejects the known positive-channel bypass.

## Usage evidence

Persisted WAL accounting for the two top-level runs records:

| Role | Calls | Input | Output | Cache read | Total tokens |
|---|---:|---:|---:|---:|---:|
| writer / Planner | 14 | 690,697 | 78,563 | 636,800 | 769,260 |
| plan_grounding | 9 | 339,526 | 19,974 | 20,480 | 359,500 |

One interrupted Grounding attempt has an explicit accounting gap with zero assigned tokens. The later direct Grounding-only regression is recorded separately and is not included in this WAL aggregate. Cost was not available.

## Runtime conclusion

The fix prevents known unsupported facts and stale receipt reuse, but the bounded Planner runs did not produce a plan that passes the current protocol. P0-5 therefore cannot claim runtime PASS.
