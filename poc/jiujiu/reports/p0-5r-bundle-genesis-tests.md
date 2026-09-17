# P0-5R Bundle Genesis Deterministic Tests

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

| Case | Expected | Result |
|---|---|---|
| First bundle, nil prior list | derived generation genesis + base state | PASS |
| First bundle, empty prior list | identical to nil prior list | PASS |
| Second bundle | Chapter 1 digest + projected post-state | PASS |
| Empty predecessor digest | builder rejects invalid digest | PASS |
| Formatted fake predecessor | chain rejects non-genesis tail | PASS |
| Genesis from wrong generation | chain rejects wrong genesis | PASS |

The passing first/second cases use real generic artifacts, `buildPipelineProjectedChapterBundle`, bundle validation and chain validation. The fake/wrong cases deliberately reach the chain authority check rather than relying only on digest syntax.

```text
go test ./cmd/novel-studio -run TestPipelineProjectAllTailAuthorityMatrix -count=1
PASS
```

These are synthetic deterministic bundles. They are not a Chapter 1 Model Runtime Bundle and do not alter the historical Runtime result.
