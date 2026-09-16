# Phase 0 Resume

Status: **IN PROGRESS**

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
