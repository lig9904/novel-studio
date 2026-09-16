# Phase 0 Resume — Model Comparison Blocker

Status: **STOPPED BEFORE OUTLINE-ALL PUBLICATION**

This comparison used the same corrected-input TEST_SCAFFOLD foundation and did not change core source, official input, Canon, Promote, Render, or Phase 1 state.

## Models compared

- `deepseek-flash`: completed operations 1 and 2; operation 3 exhausted the four-turn limit in two bounded runs.
- `deepseek-v4-pro`: completed operations 1 and 2; operation 3 exhausted the four-turn limit after long remote responses.
- `glm-5.3`: live Coding Plan catalog and connectivity check passed; operation 1 completed, but operation 2 exhausted its four turns while correcting contract mapping.
- `codex/gpt-5.6-sol`: completed operations 1 and 2; operation 3 produced eight complete structured candidates across two bounded runs, but every candidate was rejected.

## GPT-5.6 Sol evidence

The GPT-5.6 Sol attempt retained durable receipts for:

- operation 1 `plan_structure`;
- operation 2 `map_contracts`.

Operation 3 remained pending and no `0003.receipt.json` was created. The model outputs were complete (roughly 8–10K runes per turn); no `finish_reason=length`, output truncation, context-length error, or token-limit error occurred.

The final rejections were `planned_resolution_evidence_missing`. The validator required chapter 3 `core_event/scenes` to positively realize every planned resolution as an actor + action + terminal state and explicitly rejected negation, quoted-only evidence, or a future/host plan.

## Structural incompatibility

Several bound contracts are not story events and cannot be positively enacted without changing their meaning:

- do not create a sender, messenger, credential, document, or historical evidence;
- keep the opportunity `UNVERIFIED`, sender unavailable, and credential absent;
- keep unavailable materials only as non-blocking unresolved conditions;
- stop at one volume / one arc / three chapters and do not render prose;
- keep TEST_SCAFFOLD outside official Canon and Human Approval;
- keep `fac_test_scaffold_boundary` as a host audit boundary, never an in-world actor.

These are negative invariants, unresolved-state constraints, or host execution boundaries. The current outline-all contract layer maps them to chapter payoff contracts, while the payoff validator refuses negated/non-event evidence. A model cannot simultaneously preserve the author boundary and satisfy the positive-event validator for every such item.

## Classification

- Output length insufficient: **NO**.
- Context window insufficient: **NO evidence**.
- Single-model weakness: **NO; reproduced across four model configurations**.
- TEST_SCAFFOLD wording issue: **partly; the author-contract selection bound host/negative rules as story payoff contracts**.
- Product contract design issue: **YES; contract kinds do not distinguish positive story payoff from invariant, prohibition, unresolved condition, or host-only boundary**.
- Core change required for a general solution: **YES**, or an equivalent first-class contract classification that keeps non-story constraints out of positive chapter payoff validation.

## Stop condition

STOP. The new structural blocker is confirmed. Repeating model calls cannot make a prohibited/non-event condition into a positive story event without violating the author contract. Rehearsal, `ready_for_detail`, project-all, and Character Agent were not reached.

