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

## Additional corrected-input continuation

A later isolated attempt froze the Gate correction in the original TEST_SCAFFOLD author source rather than applying it as operation 0. Architect readiness passed after removing one shared placeholder alias that had merged 九尾狐 and 凤凰 into the same AgentID.

The corrected author contracts then survived through outline-all operations 1 and 2. DeepSeek Flash failed to converge on operation 3 `expand_arc` twice within the direct Architect four-turn limit, so outline-all did not publish and the Rehearsal Gate was not reached. See `phase0-resume-fresh-input-blocker.md`. This does not change the Phase 0 decision.

Authorized DeepSeek V4 Pro, GLM-5.3, and GPT-5.6 Sol comparisons later reproduced the boundary. GPT-5.6 Sol completed operation 2 and submitted eight full operation-3 candidates; the validator rejected negative invariants and host-only rules because they were not positively enacted as chapter events. This confirms a contract-kind/payoff-validation design blocker rather than a context or output-length shortage. See `phase0-resume-model-comparison-blocker.md`.
