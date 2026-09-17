# P0-5R Planner-only Trace Contract

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Tracing is opt-in and supplied by the controlled caller. Default Project-All does not create this evidence. The trace is bounded to 128 events by default and 1024 maximum; exceeding the bound sets a sticky fatal error, cancels the run and blocks every later Planner provider call.

Each event contains only:

```text
sequence
execution_id
stage / tool
sanitized argument keys
argument digest
before partial digest
after merge/source-anchor state digest
persisted partial digest
validation result / bounded error
```

Raw prompts, credentials, authentication headers, private reasoning and Codex session data are not recorded. Tool arguments are represented by at most 32 top-level keys, each at most 64 runes, plus a deterministic digest. Stage/tool labels are bounded to 64 runes, errors are reduced to fixed classifications, and each serialized event must be at most 4096 bytes.

The P0-5 harness writes the recovery contract before execution. Its callback appends and `fsync`s each trace event as JSONL, so a Host error or process-level timeout retains all events completed before interruption. A trace sink failure becomes `TRACE_SINK_FAILED` and terminates the run instead of degrading into an ordinary recoverable tool error.

`plan_details` emits these internal phases:

```text
before_merge
after_merge
after_source_anchor
persisted
validation
```

Generic tool wrappers also record `tool_start` and `tool_end`. Host operations record preflight, craft receipt, pre-Planner revalidation and `novel_context(planning)`. Grounding retains its existing structured audit keyed by protocol/input digest; invalid results remain Host errors and do not become passing receipts.

The machine-readable file `p0-5r-planner-only-runtime-evidence/synthetic-trace-example.json` is a format example only. It is not a reconstructed historical trace and marks nonexecuted stages explicitly.

## Future real-run budgets

| Control | Existing behavior |
|---|---|
| Planner turns | `min(configured writer max_turns, 36)` |
| Grounding response | one model response, exactly one verdict tool call, max output 6144 tokens |
| Context/input budget | current Planner/Grounding validators remain active |
| Cancellation | caller context cancels before or during Planner/Grounding; incomplete work does not become success |
| Usage | existing started/completed WAL and accounting hooks |
| Trace | 128 default / 1024 hard maximum |
| Overall Planner-only deadline | GAP — supplied by caller context; this task does not add a new budget system |
| Per-call hard timeout | GAP for broad Planner/Grounding calls; existing typed hard timeout is limited to frozen-render roles |

No future timeout or token cost may be inferred from the historical ten-minute harness timeout.
