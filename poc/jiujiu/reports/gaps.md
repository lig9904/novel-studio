# Gaps

## GAP-001 — CLOSED BY AUTHORIZATION — No Phase 0 evaluation entrypoint

Input: approved canon plus a requirement to test initial character state, isolated proposals, Hard Canon, Knowledge Boundary, Character Logic, Proposal Isolation, and Recovery without Season Planning.

Observed interface: `--init-only` always runs `outline-all`; direct zero-init requires published outline-all; no proposal/test-only CLI exists.

Resolution in this run: the user explicitly authorized `outline-all` as TEST_SCAFFOLD. The original blocker remains historical evidence but no longer blocks initialization.

Acceptance for a future fix: a first-class, non-production evaluation route must load explicit canon, create candidate-only branches, execute the real Character Agent and World Arbiter, persist usage/checkpoints, forbid Accepted Canon writes, and produce replayable evidence without outline/season/episode planning.

## GAP-002 — CLOSED — Full-package MCP inventory timeout was host-load-sensitive

The fixed five-second inventory timeout had only about one second of margin on this Mac. It now defaults to 15 seconds, supports `providers.*.mcp_inventory_timeout_sec` (1–120), retains per-call refresh/no cache, and preserves size/server limits. Full serial tests, race, vet, build, and real GPT-5.6 Sol outline-all calls passed.

Acceptance: the exact required command passes reliably on this Mac or the timeout/process design is made scheduling-tolerant without weakening fail-closed isolation.

## GAP-003 — P2 — Container build provenance is missing

The container built and passed health, but `novel-studio --version` inside the image reports `dev`, `commit: unknown`, and `built: unknown`.

Acceptance: Docker build injects and verifies the source commit and build timestamp.

## GAP-004 — P3 — Race helper scripts require newer GNU Bash

CI helper scripts use `mapfile`, unavailable in macOS Bash 3.2. They passed unchanged under Homebrew Bash 5.3.20.

Acceptance: document GNU Bash as a local prerequisite or make scripts portable.

## GAP-005 — CLOSED — Rehearsal readiness semantics

The new typed resource/observation/requiredness contracts produced report `41434f…3252` with `ready_for_detail=true`. Optional missing material no longer blocks, and project-all started real Character Agents.

Impact: Character Autonomy, actual Knowledge Boundary behavior, Proposal Isolation, and the detailed World State simulation cannot be validated. This blocks Phase 0 use of the story engine.

Acceptance: the same TEST_SCAFFOLD input must produce a `ready_for_detail=true` report, or validation must clearly distinguish non-blocking unresolved context from an actual reachability blocker without weakening fail-closed behavior.

## GAP-006 — PARTIALLY CLOSED / P1 — Structured convergence

outline-all now honors an explicit Architect setting up to a bounded 20 turns; rehearsal honors bounded settings up to 8; Character and World Arbiter defaults use their existing ceiling of 8. This allowed the real pipeline to progress. Structured outputs still consumed many retries, so deterministic repair guidance remains desirable.

Acceptance: representative complex rehearsals submit valid tool arguments within the bounded turn budget, or the product supplies deterministic schema guidance/repair that does not create an unbounded retry loop.

## GAP-007 — P1 — Current DeepSeek Flash alias has no monetary attribution

The system recorded 1,148,649 input and 319,706 output tokens for `deepseek-flash`, but `meta/usage.json` reports zero cost because the local registry does not price the current alias. Actual provider billing is therefore not represented in the project ledger.

Acceptance: current official model aliases resolve to current price metadata or are explicitly marked `monetary_cost = unavailable` rather than zero-cost.

## GAP-008 — P0 — Global character views can violate per-character knowledge boundaries

World rules have one global `character_view`. The TEST_SCAFFOLD published hard-canon visibility text containing 九九/九尾狐 identities and 泼水节 distinctions. The host copied that text into Phoenix's real observation packet even though Phoenix's minimum-knowledge contract forbade those facts.

Acceptance: character-facing rule projection must support actor scope or redact entity-specific facts against each actor's authorized knowledge. A globally safe process rule must not become a global identity dossier.

## GAP-009 — P0 — Chapter readiness can harden a soft outline after lawful refusal

九九 made a valid autonomous choice not to pursue an unsupported reputation opportunity. Four cycles produced actual observation/time/resource consequences, yet readiness continued because the soft-outline buoy/opportunity event had not become a sufficiently eventful narrative unit.

Acceptance: readiness must distinguish unmet hard contracts from soft-guidance non-occurrence. A lawful Character Agent refusal may alter or cancel a soft beat and still close the chapter when the resulting consequences form a complete bounded unit.

## GAP-010 — P1 — No first-class alternative proposal branch isolation harness

Runtime proposals are isolated by generation/cycle/round/actor, but the product has no supported Phase-0 harness to create PHOENIX-A with a branch-only secret and independently run PHOENIX-B against the same frozen baseline.

Acceptance: provide an isolated candidate-branch evaluation API with explicit parent root, branch-local facts, no Accepted Canon write, and a deterministic cross-branch leak assertion.
