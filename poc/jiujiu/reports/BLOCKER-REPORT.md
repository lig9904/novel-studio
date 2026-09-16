# BLOCKER REPORT

## Blocker

`BLOCKER-P0-001`: Novel Studio has no authorized Phase 0 entrypoint for initial character state plus isolated proposals/tests without first running `outline-all`.

## Reproduction

1. `./scripts/run-local.sh --pipeline --help` says `--init-only` completes full-book navigation and zero-init.
2. `cmd/novel-studio/pipeline_cmd.go` expands `--init-only` to `architect,outline-all,zero-init`.
3. The same help text defines `outline-all` as all volume/arc/chapter contracts.
4. `cmd/novel-studio/zero_init_cmd.go` requires a published `outline-all` before explicit `--zero-init` may run.
5. No CLI command exists for proposal-only/character-state-only/Hard Canon/Knowledge Boundary/Character Logic/proposal-isolation evaluation.

## Conflict with authorized Phase 0

The input task explicitly prohibits Season Planning and prohibits core-source modifications. The only native initialization route performs full-book planning; the alternative is to add a new core evaluation route or manually fabricate authoritative state. Neither is permitted.

## Preserved state

- Story-generation calls: 0.
- Connectivity calls: 1 (PASS).
- `data/` story run files: 0.
- Accepted canon changes: none.
- Candidate/pending state: none.
- Core source changes: none.
- Failure snapshot saved before any reset/rollback.

## Required human decision

Choose a later authorized path after review:

1. permit a non-production, non-canon Phase 0 evaluation harness under `poc/jiujiu/` that uses the existing Codex transport and Store contracts but is not an existing Novel Studio product entrypoint; or
2. authorize a scoped core change adding a first-class `phase0-eval` entrypoint; or
3. relax the Season Planning prohibition and allow the native `architect,outline-all,zero-init` path.

No option was assumed in this run.

## Continuation after human authorization

The user subsequently authorized `outline-all` strictly as TEST_SCAFFOLD. `BLOCKER-P0-001` was therefore closed for this run without modifying core source.

The original pipeline successfully completed Architect, outline-all, zero-init, physical resource repair, initial world tick, preplan, multiple rebase/recovery cycles, and several Architect/World Arbiter rehearsals.

## Final blocker

`BLOCKER-P0-002_REHEARSAL_NOT_READY`: the final durable World Arbiter report has digest `sha256:b700daa72c2d853ec872ec21fcc04f186d817906341e326d359581883e38611b` and `ready_for_detail=false`.

The report is valid and speculative. Its own summary says the selected conditional path has no hard-contract reachability failure, but the stored result retains 14 unresolved items and one missing material check. The supported `project-all` command refuses before any Character Agent model call when the current rehearsal is not ready.

No further scaffold expansion was attempted. Phase 0 stops with NO-GO, no accepted chapter, no proposal promotion, no render, and no core-source change.
