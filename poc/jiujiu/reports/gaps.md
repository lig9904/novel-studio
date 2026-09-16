# Gaps

## GAP-001 — CLOSED BY AUTHORIZATION — No Phase 0 evaluation entrypoint

Input: approved canon plus a requirement to test initial character state, isolated proposals, Hard Canon, Knowledge Boundary, Character Logic, Proposal Isolation, and Recovery without Season Planning.

Observed interface: `--init-only` always runs `outline-all`; direct zero-init requires published outline-all; no proposal/test-only CLI exists.

Resolution in this run: the user explicitly authorized `outline-all` as TEST_SCAFFOLD. The original blocker remains historical evidence but no longer blocks initialization.

Acceptance for a future fix: a first-class, non-production evaluation route must load explicit canon, create candidate-only branches, execute the real Character Agent and World Arbiter, persist usage/checkpoints, forbid Accepted Canon writes, and produce replayable evidence without outline/season/episode planning.

## GAP-002 — P2 — Full-package test timeout is host-load-sensitive

`go test -count=1 ./...` failed twice because two fake-Codex MCP inventory subprocesses exceeded a fixed 5-second deadline under package parallelism. Both isolated tests and the complete `-p 1` suite passed unchanged.

Acceptance: the exact required command passes reliably on this Mac or the timeout/process design is made scheduling-tolerant without weakening fail-closed isolation.

## GAP-003 — P2 — Container build provenance is missing

The container built and passed health, but `novel-studio --version` inside the image reports `dev`, `commit: unknown`, and `built: unknown`.

Acceptance: Docker build injects and verifies the source commit and build timestamp.

## GAP-004 — P3 — Race helper scripts require newer GNU Bash

CI helper scripts use `mapfile`, unavailable in macOS Bash 3.2. They passed unchanged under Homebrew Bash 5.3.20.

Acceptance: document GNU Bash as a local prerequisite or make scripts portable.

## GAP-005 — P0 — Rehearsal cannot become ready, so project-all never starts

The final speculative World Arbiter report says the selected conditional path has no hard-contract reachability failure, yet it retains 14 unresolved items, one missing material check, and `ready_for_detail=false`. The supported `project-all` command then refuses before any Character Agent model call.

Impact: Character Autonomy, actual Knowledge Boundary behavior, Proposal Isolation, and the detailed World State simulation cannot be validated. This blocks Phase 0 use of the story engine.

Acceptance: the same TEST_SCAFFOLD input must produce a `ready_for_detail=true` report, or validation must clearly distinguish non-blocking unresolved context from an actual reachability blocker without weakening fail-closed behavior.

## GAP-006 — P1 — Structured rehearsal convergence is bounded below observed repair need

Architect and World Arbiter rehearsal loops allow four turns. GPT-5.6 Sol, DeepSeek V4 Pro, and DeepSeek Flash repeatedly needed more schema/capability corrections, especially for qualitative observation, secret mechanisms, and future artifact dependencies.

Acceptance: representative complex rehearsals submit valid tool arguments within the bounded turn budget, or the product supplies deterministic schema guidance/repair that does not create an unbounded retry loop.

## GAP-007 — P1 — Current DeepSeek Flash alias has no monetary attribution

The system recorded 1,148,649 input and 319,706 output tokens for `deepseek-flash`, but `meta/usage.json` reports zero cost because the local registry does not price the current alias. Actual provider billing is therefore not represented in the project ledger.

Acceptance: current official model aliases resolve to current price metadata or are explicitly marked `monetary_cost = unavailable` rather than zero-cost.
