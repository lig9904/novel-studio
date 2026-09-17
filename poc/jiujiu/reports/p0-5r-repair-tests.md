# P0-5R Repair Deterministic Verification

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Date: 2026-09-17

## Controlled pre-fix reproduction

The generic contract tests were applied to fixed baseline `9b2e92cecc86c1cb81e09fa13785fb73f05b3932` in a temporary detached worktree.

```text
RESULT: EXPECTED FAIL

TestPlanDetailsSchemaPublishesExternalReferencePatchContract
  focused schema omitted external_reference_plan

TestPlanDetailsExplicitEmptyExternalReferencePlanClearsCurrentRows
  explicit [] restored the old eligible current-receipt row

TestPlanDetailsClearThenReplaceCurrentRAGRow
  stale and corrected authored rows coexisted because clear did not take effect
```

## Final verification

```text
go test ./... -count=1
PASS

go vet ./...
PASS

go build ./...
PASS

go test -race ./internal/tools -run 'TestPlanDetailsSchemaPublishesExternalReferencePatchContract|TestPlanStructureResubmissionPreservesBoundCandidateWithoutPromotion|TestPlanDetailsOmissionRetainsCurrentRAGFactRow|TestPlanDetailsExplicitEmptyExternalReferencePlanClearsCurrentRows|TestPlanDetailsNullExternalReferencePlanIsNotExplicitRemoval|TestPlanDetailsCurrentRAGRowUsesCompleteIdentity|TestCurrentRAGFactExternalRowKeyUsesCompleteAuthoredIdentity|TestPlanDetailsClearThenReplaceCurrentRAGRow|TestPlanDetailsSequentialPatchesPreserveCurrentRAGFactRow|TestPlanDetailsDoesNotPreserveStaleOrIncompleteRAGFactRows|TestPlanDetailsRAGFactHitAliasesFailClosed|TestPlanDetailsFinalizeIncompleteListsMissing|TestPlanGrounding|TestPlanGroundingResolve' -count=1
PASS

git diff --check
PASS
```

The first full-suite run after the schema fix correctly detected three pinned Project-All planning protocol digests. The Planner tool contract is part of that digest, so their fixtures were advanced together. One-shot Character and activation actor/arbiter protocol digests remained unchanged. The final full suite passed.

After the final explicit-clear control-flow change, the full suite, vet, build, targeted race command and `git diff --check` were all rerun and passed. The results therefore cover Core commit `99eec8bf53f098e0d9e61f492a84d320aae33977` exactly.

No Novel Studio business model, Planner, Grounding Reviewer, Character Agent, World Arbiter or Story Simulation was executed.
