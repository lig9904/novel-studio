# P0-4 Planner Grounding Model Comparison

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Result: **CONTROLLED DIAGNOSIS COMPLETE — BOTH MODEL COMPLIANCE GAP AND PROTOCOL USABILITY DEFECT CONFIRMED**

`PASS` below means the Reviewer submitted one structured verdict that passed the complete Host schema and evidence-pointer validation. It does not mean the frozen Chapter Plan was grounded. Both repaired Sol runs correctly returned `pass=false` with valid findings.

## Frozen control

- Story input digest: `sha256:5a9a512d20a486d886c67e26e637b4ef56647b0cb5868b7df0da181f2be938c4`
- Story State excluding model-bound `review_protocol`: `sha256:dfec76f217ea6de6eddd87a5c7a251cb01fabad9c075a019103db57fa1f88fcb`
- P0-2 generation: `pg2_b2354a98e8714bba2baca89a`
- Chapter / simulation: `1` / `ch001-4db45dbca11c`
- Story Simulation was not rerun.

## Controlled results

| Run | Model / reasoning | Prompt/schema | Tool calls | Named verdicts | Host accepted | Duration | Usage in/out | Result |
|---|---|---|---:|---:|---|---:|---:|---|
| DeepSeek baseline | DeepSeek Flash / High | original | 0 | 0 | no | 33.314 s | 26,595 / 6,144 | **FAIL** |
| Sol run 1 | GPT-5.6 Sol / Medium | original | 1 | 1 | no; invalid plan pointer | 53.069 s | 40,429 / 1,890 | **FAIL** |
| Sol run 2 | GPT-5.6 Sol / Medium | original | 1 | 1 | no; invalid plan pointer | 35.347 s | 40,429 / 1,337 | **FAIL** |
| Sol fixed run 1 | GPT-5.6 Sol / Medium | explicit nested paths | 1 | 1 | yes | 71.672 s | 40,564 / 2,345 | **PASS** |
| Sol fixed run 2 | GPT-5.6 Sol / Medium | explicit nested paths | 1 | 1 | yes | 33.793 s | 40,564 / 1,115 | **PASS** |

No isolated Reviewer call retried internally; `retry_count=0` for every row.

## DeepSeek baseline diagnosis

The Host offered exactly one tool, but DeepSeek returned zero tool calls and no assistant text. All 6,144 output tokens were consumed by reasoning content (`thinking_runes=22,961`) before a tool submission. The Host therefore rejected the response with:

```text
plan grounding must return exactly one structured verdict
```

This confirms a model capacity/tool-compliance gap on this 90,311-byte exact activation packet. The core protocol should not be weakened to accept an absent verdict.

## Original Sol diagnosis

Sol Medium selected the tool correctly in both runs and produced syntactically valid arguments. Both responses failed evidence-pointer validation in the same way:

- correct contract path: `/plan/contract/required_beats/...`;
- incorrect causal paths: `/plan/render_capacity/...` and `/plan/ending_consequence_contract/...`;
- actual input paths: `/plan/causal_simulation/render_capacity/...` and `/plan/causal_simulation/ending_consequence_contract/...`.

The original prompt named causal fields but did not state their serialized `/plan/causal_simulation/...` namespace. The tool schema described `plan_path` only as a generic JSON pointer. Two independent Sol failures establish that the interface was not reliably usable even though the Host validator was correct.

## Fixed Sol validation

The repair made the real nested namespaces explicit in both the system prompt and `plan_path` schema description and required the model to verify that every pointer exists and every quote is a literal substring. The Host validator was not relaxed.

Two independent Sol Medium runs then submitted one verdict each; all paths and quotes validated. Both verdicts correctly rejected the frozen Planner plan because it added facts absent from the activation evidence, including the float-rope/water-gauge result and a visible visitor interaction.

## Hypothesis decision

```text
Hypothesis A — Model Capability: CONFIRMED AS ONE FAILURE LAYER
Hypothesis B — Protocol / Wiring Defect: CONFIRMED AS A SECOND FAILURE LAYER
```

Routing alone did not pass the original protocol, so the result is not Case A. A narrow prompt/schema repair was required. After that repair, Sol Medium passed twice, so the Planner Grounding role should be routed to Sol / Medium. DeepSeek Flash is proven to fail only under the original protocol/High/6,144-token conditions; it was not rerun after the repair, so no broader post-fix suitability claim is made.
