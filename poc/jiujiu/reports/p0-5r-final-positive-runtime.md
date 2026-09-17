# P0-5R Final Positive Runtime Evidence

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

## Entry and deadline

```text
ENTRY: RunPlannerOnlyProjectedChapterPlanning
DEADLINE: 15 minutes
TOP-LEVEL ATTEMPTS: 1
RESULT: TIMEOUT
ELAPSED: 900.993 seconds
```

The fresh evaluation copy began with one unfinalized, source-bound partial and
no formal Plan. Planner-only preflight passed before the first provider call.
The recovery contract pinned generation, chapter, planning context, Planner
protocol, Simulation, Simulation checkpoint, activation evidence and initial
partial digest.

The durable trace retains 70 contiguous events through the final completed
Planner persist. It records ten Planner calls and eight Grounding calls. The
last Grounding call was in flight when `go test -timeout=15m` stopped the only
top-level attempt.

## Grounding results

```text
GROUNDING CALLS STARTED: 8
GROUNDING CALLS WITH TERMINAL USAGE: 7
HOST-VALID STRUCTURED RECEIPTS: 6
VALID REJECTION RECEIPTS: 6
PASSING RECEIPTS: 0
HOST-REJECTED UNVERIFIABLE RESPONSES: 1
TIMED-OUT GROUNDING CALLS: 1
```

The valid receipts rejected unsupported float-rope displacement, water-gauge
anomalies, live visitors, new communication, tide deadlines and unsupported
ending consequences. Host did not turn any rejection or invalid response into
a formal Plan.

## Frozen-state result

`frozen-comparison.json` proves every frozen item byte-identical. The only
changed item is `drafts/01.plan.partial.json`, which is explicitly mutable and
whose initial identity was bound by the recovery contract.

```text
FROZEN STATE UNCHANGED: YES
STORY STATE RERUN: NO
CHARACTER AGENT CALLS: 0
WORLD ARBITER CALLS: 0
READINESS CALLS: 0
STORY SIMULATION CALLS: 0
```
