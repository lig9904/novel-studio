# Phase 0 Resume — Fresh Corrected Input Attempt

Status: **STOPPED BEFORE REHEARSAL GATE**

This is a continuation of Phase 0. It does not authorize Phase 1, core-source changes, Promote, Render, or Accepted Canon writes.

## Why this attempt was run

The prior host-only `outline-repair-file` operation 0 succeeded but was necessarily followed by `plan_structure` and `expand_arc`; the latter regenerated the repaired chapters. The fixed upstream baseline has no post-expansion repair boundary.

This attempt therefore tested the remaining legal input-level alternative: start a new isolated TEST_SCAFFOLD project with the corrected Gate rules frozen in the original author source before Architect and outline-all run.

## Input and model

- Model: `deepseek/deepseek-flash` for all roles.
- Corrected source: `poc/jiujiu/project/test-scaffold-resume-prompt.md`.
- Official `poc/jiujiu/input/jiujiu-poc-input.md`: unchanged, SHA-256 `7ef95a96f17ac3290a0bff2a3d00afd0097bc919cc7797165dda428a8d28715a`.
- Core source: unchanged from baseline `5e4d6912bae5f20a6224fa84ecd0e34b7838118e`.

## Architect result

The first fresh attempt selected `architect_short` and produced only a flat outline; unmodified `outline-all` correctly rejected the missing `layered_outline`.

The second fresh attempt explicitly selected the Architect long/layered path. Architect readiness initially found a TEST_SCAFFOLD identity collision: 九尾狐 and 凤凰 shared the placeholder alias `（正式名 TBD·TEST_SCAFFOLD 占位）`, so the registry resolved both as one AgentID and physical-state validation reported `duplicate actor`. Removing the shared placeholder alias through the supported `--refresh-architect --architect-target characters` path resolved the issue. Architect readiness then passed.

## outline-all result

Operations 1 and 2 completed and have durable receipts:

- operation 1 `plan_structure` after digest: `sha256:0d5ddae309c7f0f75bf15ecde0c4704c437376a5810db0973d811be57dc765e0`;
- operation 2 `map_contracts` after digest: `sha256:67f2656a255d6558883f56252214b4b9255777f98dc844ce04f48a12bc6c41bd`.

Operation 3 `expand_arc` was attempted twice with DeepSeek Flash. Both attempts exhausted the direct Architect four-turn structured-output limit without producing `0003.receipt.json`.

The execution remains:

- `status=building`;
- `completed_action_count=2`;
- pending operation 3 `expand_arc` for V1A1, span 3;
- no published outline-all receipt;
- no final outline-all digest.

## Gate consequence

Because outline-all never published, zero-init, preplan, rehearse-arc, `ready_for_detail`, project-all, and Character Agent were not run. The pre-outline-all chapter draft is not a published full-book contract and cannot be used as Gate evidence.

## Classification

- Corrected TEST_SCAFFOLD author input: **ACCEPTED BY ARCHITECT READINESS**.
- Shared placeholder alias collision: **TEST_SCAFFOLD input defect, corrected**.
- Original operation-0 overwrite defect: **still present in the product design; not used in this attempt**.
- Initial blocker: **DeepSeek Flash structured-output convergence at outline-all operation 3**.
- Later cross-model result: DeepSeek V4 Pro, GLM-5.3, and GPT-5.6 Sol did not remove the boundary. GPT-5.6 Sol reached operation 3 repeatedly and proved the final rejection is a positive-payoff validator conflict with negative/host-only contracts, not output length. See `phase0-resume-model-comparison-blocker.md`.
- Core change required for a general solution: **YES**, or an equivalent first-class contract classification that keeps invariants, prohibitions, unresolved conditions, and host-only boundaries out of positive chapter payoff validation.
- World Arbiter too strict: **not applicable**, because rehearsal did not run.

## Stop condition

STOP. Repeating the same four-turn call is not a bounded test strategy. No core source was modified, no official Canon was added, and no Promote/Render/Phase 1 action occurred.
