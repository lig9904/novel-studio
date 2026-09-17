# P0-5R Planner-only Deterministic Tests

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

All model activity used in-memory fake/spy models. No provider or Novel Studio business model was called.

| Case | Planner calls | Grounding calls | Character/World/Readiness calls | Result |
|---|---:|---:|---:|---|
| Valid frozen sources and bound partial | 4 fake calls stopped by existing Planner stop guard | 0 | 0 | Entered the production Planner loop only |
| Wrong generation | 0 | 0 | 0 | FAIL CLOSED |
| Wrong chapter | 0 | 0 | 0 | FAIL CLOSED |
| Wrong planning context | 0 | 0 | 0 | FAIL CLOSED |
| Wrong checkpoint digest | 0 | 0 | 0 | FAIL CLOSED |
| Wrong partial digest | 0 | 0 | 0 | FAIL CLOSED |
| Wrong Planner protocol | 0 | 0 | 0 | FAIL CLOSED |
| Missing simulation artifact | 0 | 0 | 0 | FAIL CLOSED |
| Checkpoint-current but semantically not-ready simulation | 0 | 0 | 0 | FAIL CLOSED |
| Missing exact projected-context source | 0 | 0 | 0 | FAIL CLOSED |
| Tampered activation/readiness evidence | 0 | 0 | 0 | FAIL CLOSED |
| Parent cancellation | 0 | 0 | 0 | FAIL CLOSED |
| Trace event limit exceeded | 0 | 0 | 0 | FAIL CLOSED |
| Same contract used concurrently | second attempt 0 | 0 | 0 | Unique attempt owner rejects reentry |
| Trace sink fails after first tool | exactly 1 fake call | 0 | 0 | Sticky fatal cancels run; no subsequent provider call |
| Oversized/error-bearing trace event | 0 | 0 | 0 | Keys/text/event bytes bounded; error reduced to classification |

The valid synthetic attempt revalidated the original contract after the fake Planner stopped, proving that simulation, checkpoint, evidence and historical partial stayed unchanged. Invalid attempts also verified that the partial digest did not change.

The opt-in harness writes its recovery contract before execution and appends each trace event to a synced JSONL file. A deterministic file test confirms earlier events remain readable without waiting for successful completion.

Existing deterministic suites additionally cover:

- exact Grounding input/protocol binding and stale-receipt rejection;
- explicit RAG row clear, clear-then-replace and no premature auto-finalize;
- normal Project-All Character/World recovery behavior;
- accounting start/record/flush and canceled-call behavior;
- current Plan checkpoint and simulation causality gates.

Verification:

```text
go test ./... -count=1: PASS
go vet ./...: PASS
go build ./...: PASS
targeted planner-only/tools race: PASS
git diff --check: PASS
```

The full deterministic checks were repeated in a detached worktree at exact Core `70bc232a4bc7aac1d4746fab6cddc7f21ec56ba1`. The Evidence-only harness update was verified separately in the final combined tree.

The live P0-5 positive test was not executed.
