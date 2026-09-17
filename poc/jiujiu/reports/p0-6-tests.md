# P0-6 Tests

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Date: 2026-09-17

Status: **PASS**

Verified Core SHA: `c7ef060180d3c86c7e7f89a84a892238dc080a57`

## P001 acceptance matrix

| Case | Expected | Result |
|---|---|---|
| Persist P001 as pending proposal | PASS | PASS |
| Model submits approval/status/match identity | REJECT | PASS |
| Dedicated proposal list returns P001 with labels | PASS | PASS |
| Shared `novel_context` exposes P001 | REJECT | PASS |
| Proposal RAG enters fact recall/receipt | REJECT | PASS |
| Relabel P001 as chapter-summary fact | REJECT | PASS |
| Pending queue or legacy index reintroduces P001 | REJECT | PASS |
| Foundation/summary/world/Character/Plan/commit asserts P001 | REJECT | PASS |
| Bundle uses exact, supported paraphrase or structured P001 | REJECT | PASS |
| Proposal source downgrades Authority | REJECT | PASS |
| Negative/TBD statement mentions P001 | PASS | PASS |
| Phoenix and fire belong to different sentences/actors | PASS | PASS |
| Phoenix holds a torch / sees another fire character | PASS | PASS |
| Mixed structure has `attribute=fire` and unrelated `healing=TBD` | REJECT | PASS |
| Shadow registry missing/replaced | REJECT | PASS |

## Deterministic verification

```text
go test ./... -count=1
PASS

go vet ./...
PASS

go build ./...
PASS

go test -race ./internal/store ./internal/tools -run 'Proposal|P06' -count=1
PASS

git diff --check
PASS
```

The final full suite ran after all Reviewer fixes. No model or Story Simulation was executed.
