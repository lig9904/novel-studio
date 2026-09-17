# P0-2 Test Results

Date: 2026-09-17

Result: **PASS**

## Five-state semantic cases

| Required semantic | Repository outcome | Result |
|---|---|---|
| Accepted | `OCCURRED` | PASS |
| Plain rejection without consequence | `DEFERRED` + `continue` | PASS |
| Rejected with actual consequence | `REJECTED_WITH_CONSEQUENCE` | PASS |
| Actual choice supersedes direction | `SUPERSEDED_BY_ACTUAL_CHOICE` | PASS |
| Unresolved evidence | `DEFERRED` + `continue` | PASS |
| Hard contract impossible | `HARD_CONTRACT_UNSATISFIED` + `hard_conflict` | PASS |

Negative tests reject paraphrased character reasons, invented consequences, proposal-only proof, wrong/multiple-cycle proof, no actual progress, unknown evidence aliases and impossible-hard-contract bypass.

## Freeze-order cases

- current V3 producer freezes a v2 readiness context before first dispatch: PASS;
- binding contains exact producer, activation policy and Soft Event policy inventory: PASS;
- activation session binds the final context digest: PASS;
- new stimulus with legacy frozen context: FAIL CLOSED;
- v2 context with the Soft Event marker removed: FAIL CLOSED;
- producer mismatch or downgraded policy inventory: FAIL CLOSED.

## Commands

- `go test -p 1 -count=1 ./...`: PASS.
- targeted `go test -race` for Character Readiness, Character Activation, Soft Event, World Arbitration, Project-all and P0-3 boundaries: PASS.
- `go vet ./...`: PASS.
- `go build ./...`: PASS.
- full `internal/domain`, `internal/agents`, and `internal/store` packages: PASS.

P0-3 protected canon and derived coherence regressions remained PASS.
