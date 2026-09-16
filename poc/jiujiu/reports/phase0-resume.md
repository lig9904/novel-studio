# Phase 0 Resume

Status: **COMPLETE / STOPPED AT OUTLINE REPAIR PUBLICATION**

This is a continuation of the original Phase 0, not Phase 1.

Baseline core commit: `5e4d6912bae5f20a6224fa84ecd0e34b7838118e`.

The branch HEAD contains prior `poc/jiujiu/**` evidence commits, but the core source tree has no differences from the fixed baseline. The official input copy still matches the source SHA-256 `7ef95a96f17ac3290a0bff2a3d00afd0097bc919cc7797165dda428a8d28715a`.

Confirmed prior Gate root cause:

- 14 `unresolved_items` are not readiness blockers.
- All 10 contract checks were `plausible` or `conditional`.
- The sole blocker was the selected-path material check “第一章：九九核查声望机会口信的来源与凭据” with `status=missing`.
- The resume prompt removes that operation from the selected path while preserving the unverified opportunity and all Character Agent choices.

Authorized sequence: rebase → outline-all → zero-init → preplan → rehearse-arc; only after a legal `ready_for_detail=true` may project-all continue.

Still prohibited: core-source changes, official Canon additions, fake sender/credential/evidence, digest edits, Gate bypass, direct Character Agent invocation, promote, render, and Phase 1 work.

## Execution result

1. rebase: PASS; old chapter-zero generation archived with accepted chapter still 0.
2. outline-all: PASS, but the freely generated outline still introduced a messenger-like contact and mandatory evidence-seeking behavior.
3. deterministic `outline-repair-file`: operation 0 PASS; all three chapters were replaced in the isolated candidate.
4. outline-all continuation: FAIL semantically; operation 3 overwrote the repaired chapter contracts before publication.
5. zero-init/preplan/rehearse-arc for the repaired outline: NOT RUN, because the final published outline no longer matched the authorized repair.
6. project-all and Character Agent: NOT RUN.

The supported pipeline offers no post-operation-3 repair stage. Continuing requires a core product fix or manual live-state mutation; both are outside Phase 0 authorization. Final decision remains NO-GO.
