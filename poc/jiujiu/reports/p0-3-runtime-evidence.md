# P0-3 Same-PoC Runtime Evidence

Date: 2026-09-17

Status: **P0-3 PASS / P0-2 RUNTIME FAIL TO ACTIVATE / STOPPED**

Classification: **TEST_SCAFFOLD / NOT ACCEPTED CANON / NOT HUMAN APPROVAL**

## Native chain

- rebase: PASS;
- outline-all: PASS;
- zero-init: PASS;
- preplan: PASS;
- rehearse-arc: PASS (`ready_for_detail=true`);
- project-all Chapter 1: real Character Agent and World Arbiter executed;
- Chapter 2 model work: NOT RUN;
- Chapter 1 bundle: NOT CREATED;
- seal/promote/render: NOT RUN.

All model roles used `deepseek/deepseek-flash`. The retained credential stayed in macOS Keychain and was not committed.

## P0-2 activation finding

Generation `pg2_f2fc66b794ea7dd681ea0ec9` used character activation policy `chapter-activation-cycles.v3`. Its real first-cycle stimulus correctly contains:

- `character-soft-event-readiness:v1`;
- producer `character-agent-protocol:sha256:3a49b7bd92e39427a395234e6f11c7a98139b9622591c4b25a502d12f5ca596e`.

However, the readiness context was frozen before those producer policies were attached. The persisted context and audit therefore remained:

- context version: `chapter-readiness:arbitrated-events.v1`;
- input policy: `chapter-readiness:arbitrated-events.v1`;
- model view: `readiness-model-view.requirements-once-aliases.v1`;
- receipt version: `character-chapter-readiness.v2`;
- `soft_event`: `null`.

The real first cycle contains a lawful autonomous rejection by 九九 and an actual consequence, but the old readiness path returned `continue` and requested another cycle. Therefore this run does not prove either `REJECTED_WITH_CONSEQUENCE` or `SUPERSEDED_BY_ACTUAL_CHOICE`; the P0-2 five-state runtime contract was not activated.

This is a P0-2 wiring/freeze-order defect, not a DeepSeek classification result. The task explicitly prohibited modifying P0-2, so execution was stopped before completing cycle 2. No Chapter 1 bundle exists and no Chapter 2 model call occurred.

## Preserved boundaries

- Phoenix's real first-cycle observation contains none of 九九, 螭吻, 九尾狐, 泼水节, 独角龙 or `TEST_SCAFFOLD_SECRET_FOX_TIDE_MARK`.
- private fox observation remained private in the first arbitration.
- official input SHA-256 remained `7ef95a96f17ac3290a0bff2a3d00afd0097bc919cc7797165dda428a8d28715a`.
- live accepted/current chapter remained 0.
- the interrupted pipeline lease was released through a native no-key recovery call; no paid call or story evidence was created by that cleanup.
- P0-1, Proposal Isolation, Foreshadow/RISK-001 and Phase 1 were not modified or executed.

STOP. WAITING FOR HUMAN REVIEW.

Evidence: `poc/jiujiu/snapshots/phase0-p0-3-runtime/`.
