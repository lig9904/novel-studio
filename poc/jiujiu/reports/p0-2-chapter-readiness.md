# P0-2 Chapter Readiness

Date: 2026-09-17

Status: **CORE FIX PASS / SAME-POC FULL RUNTIME BLOCKED BEFORE ZERO-INIT**

Classification: **TEST_SCAFFOLD / NOT ACCEPTED CANON / NOT HUMAN APPROVAL**

## Implemented contract

The new readiness policy classifies the soft-outline result as exactly one of:

- `OCCURRED`
- `REJECTED_WITH_CONSEQUENCE`
- `SUPERSEDED_BY_ACTUAL_CHOICE`
- `DEFERRED`
- `HARD_CONTRACT_UNSATISFIED`

The Host, not the model, enforces the closure boundary. A closing result must bind the exact actor and proposal, copy the actual `decision_reason`, copy an actual `immediate_result` or `state_after`, cite the proposal and the same cycle's arbitration/cycle evidence, and prove real time or state progress. A bare refusal cannot close a chapter. `DEFERRED` remains `continue`; an impossible hard contract remains `hard_conflict`.

Historical readiness receipts and the previous six executable V3 producers retain their old policy, schema and digest semantics. The five-state contract is enabled only by a new frozen producer/source marker.

## Required cases

| Case | Expected | Result |
|---|---|---|
| A: lawful rejection with exact character reason and actual consequence | `REJECTED_WITH_CONSEQUENCE`, planning may proceed | PASS |
| A negative controls: paraphrased reason, invented consequence, proposal-only evidence, or no actual progress | reject | PASS |
| B: actual autonomous choice replaces the soft-outline action | `SUPERSEDED_BY_ACTUAL_CHOICE`, do not demand the old action again | PASS |
| C: hard contract is impossible | `HARD_CONTRACT_UNSATISFIED`, `hard_conflict` | PASS |

## Verification

- `go test -p 1 -count=1 ./...`: PASS.
- targeted `go test -race` for Readiness, producer and transaction paths: PASS.
- `go vet ./...`: PASS.
- `go build ./...`: PASS.
- user fork commit: `aaf50f558f47a00c6f1d2c2129a70d3bbcbcc589`.
- branch: `codex/physical-resource-semantics`.
- pushed to `origin`; no upstream PR opened.

## Same-PoC rerun

The explicit rebase completed and archived the prior exact project. DeepSeek Flash then completed all three outline-all model operations, but publication failed closed with `outline-all final candidate modified protected canon`.

The changed protected files were only the derived `meta/world_coherence_report.json/md`: the outline changed its source digest, while the current protected-canon root treats that derived report as immutable. Using the reviewed sparse outline repair reached the same failure. Restoring the exact archived, previously published outline exposed the opposite side of the same cycle: P0-1's legitimate `world_rules` visibility migration requires a refreshed coherence report, while refreshing that report invalidates the published outline receipt.

Therefore the required native chain did not reach `zero-init → preplan → rehearse-arc → project-all Chapter 1`. No claim is made that the old four-cycle loop has been eliminated in the real PoC runtime. Fixing the outline-all/protected-coherence contract would be a separate non-Readiness core change and was not made under the P0-2-only authorization.

Manual archive-copy recovery attempts were retained only as audit artifacts and are not counted as native pipeline evidence. The active project was returned to the product-generated chapter-zero rebase state.

## Safety boundary

- official input SHA-256: `7ef95a96f17ac3290a0bff2a3d00afd0097bc919cc7797165dda428a8d28715a` (unchanged);
- live current chapter: `0`;
- seal: NOT RUN;
- promote: NOT RUN;
- render: NOT RUN;
- Phase 1: NOT RUN;
- Proposal Isolation: NOT MODIFIED / NOT RUN;
- Foreshadow/RISK-001: NOT MODIFIED;
- DeepSeek configuration and Keychain credential: retained.

STOP for human review. The code-level P0-2 invariant is verified; the same-PoC runtime acceptance remains blocked by a separately scoped outline publication contract.
