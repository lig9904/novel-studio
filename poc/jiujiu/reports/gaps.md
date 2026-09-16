# Gaps

## GAP-001 — P0 — No Phase 0 evaluation entrypoint

Input: approved canon plus a requirement to test initial character state, isolated proposals, Hard Canon, Knowledge Boundary, Character Logic, Proposal Isolation, and Recovery without Season Planning.

Observed interface: `--init-only` always runs `outline-all`; direct zero-init requires published outline-all; no proposal/test-only CLI exists.

Impact: the requested Phase 0 cannot use the actual Novel Studio Story Engine without violating scope or changing core code.

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
