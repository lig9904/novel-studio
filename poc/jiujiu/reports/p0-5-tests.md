# P0-5 Tests

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Status: **CODE / STATIC / REGRESSION PASS; RUNTIME ACCEPTANCE BLOCKED**

## Required semantic cases

| Case | Plan behavior | Expected | Coverage and result |
|---|---|---|---|
| A | Actual buoy approach plus removable non-causal water-ripple rendering | PASS | Contract/schema tests preserve removable sensory presentation |
| B | Unsourced buoy rope | REJECT | Contract tests reject new resource; current Grounding replay rejects rope in positive RAG fields |
| C | Numeric water level without a gauge | REJECT | Contract tests reject unsourced device/measurement; current Grounding replay rejects gauge detail |
| D | Unsourced on-site visitor | REJECT | Contract tests reject new actor/meeting/dialogue |
| E | Phoenix gains UNKNOWN 九九 identity | REJECT | Contract tests require knowledge/communication authority |
| F | Two real Story Facts organized into one scene | PASS | Contract/schema tests preserve grouping without invented causality |

A–F are contract-level coverage. B/C additionally have current-protocol negative runtime evidence. Because no legal final plan was produced within the run cap, A/F do not have a positive bundle-level runtime acceptance result.

## Automated coverage

- `internal/agents/plan_projection_authority_test.go`
- `internal/tools/plan_projection_authority_test.go`
- `assets/load_test.go`
- `internal/agents/plan_grounding_reuse_test.go`
- `internal/agents/plan_grounding_test.go`
- `internal/agents/plan_grounding_snapshot_test.go`
- `cmd/novel-studio/pipeline_p0_5_live_test.go`

## Verification executed

```text
go test ./...
PASS

go test ./cmd/novel-studio -run 'TestLiveP05FrozenPlannerProjection|TestLiveP05GroundingRejectsRAGAuthorityBypass' -count=1
PASS (final scaffold projection compiled; live tests remain opt-in)

go vet ./...
PASS

go build ./...
PASS

git diff --check
PASS

p0-5 fixture: shasum -a 256 -c SHA256SUMS
PASS, 17/17 files

p0-5 runtime evidence: shasum -a 256 -c SHA256SUMS
PASS, 6/6 files
```

The full repository suite passed after the authority/Grounding/stale-receipt changes. The final subsequent edit only narrowed the opt-in fixture scanner to positive fields; its package test and formatting check passed.

## Frozen contracts

- P0-2 frozen simulation and activation-evidence file hashes remained unchanged before and after runtime.
- P0-3 Canon/coherence code was not changed.
- P0-4 routing remains Sol/Medium; strictness was strengthened. The current protocol returned one structured verdict with five exact findings for the unsafe saved plan.
- Knowledge Boundary and Character Autonomy regression suites passed as part of `go test ./...`.

## Independent review

GPT-6 Astra / High: **PASS; no remaining P0/P1/P2.** The review covered the Planner authority changes, positive Grounding channels, stale protocol gate and the fixture scanner's exclusion of source/negative metadata.
