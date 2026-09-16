# Phase 0 project configuration

- The configuration is intentionally project-scoped and contains no API key.
- All first-round roles use the Codex CLI subscription provider and `gpt-5.6-sol`; no fallback model is configured.
- Novel Studio implements Planner as the `writer` role, so the task's Planner/Medium requirement maps to `roles.writer.reasoning_effort = medium`.
- Architect, World Arbiter, Reviewer, and Coordinator use High; Character uses Medium.
- Architect `max_turns` is set to 20 after the first unmodified outline-all attempt reproducibly exhausted the default four turns without publishing a candidate. The failed attempt and interruption remain in the audit logs.
- Drafter and Editor are pinned to the same model even though Season Planning and episode rendering are outside Phase 0.
- Embedding/Qdrant and notifications are disabled for this bounded local PoC.
- A connectivity check or story generation consumes Codex subscription usage and is not performed by merely loading this file.
