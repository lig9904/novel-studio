# P0-6 Code Change

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Date: 2026-09-17

Status: **CORE IMPLEMENTED AND VERIFIED**

Core commit: `c7ef060180d3c86c7e7f89a84a892238dc080a57`

## Domain and persistence

- Added typed pending proposal registry and deterministic digests.
- Added Host-derived claim identity and fixed, non-model-writable authority/status/approval.
- Added explicit Source Binding Authority with legacy-compatible omission.
- Bound the registry into Project-All foundation snapshots and live/shadow workspace manifests.

## Context and retrieval

- Added dedicated submit/list proposal capabilities.
- Kept proposal bodies out of shared `novel_context`.
- Classified proposal RAG as design-only.
- Revalidated incoming chunks, pending queues, existing indexes and fact receipt replay.

## Fact-bearing gates

The shared proposal guard is applied at foundation, summary, world tick, Character observation/decision, World arbitration, direct simulation, staged/one-shot planning, chapter commit/rewrite recovery and Project-All bundle/recovery boundaries.

## Preserved protocols

No Character Decision, World Arbitration, Readiness, Canon, Story Simulation or Planner Grounding result was changed. No frozen system prompt was edited. No P001/九九 production branch exists; P001 appears only in tests and PoC evidence.
