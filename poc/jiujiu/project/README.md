# Phase 0 project configuration

- The configuration is intentionally project-scoped and contains no API key.
- Scaffold construction used the Codex CLI subscription provider and `gpt-5.6-sol`. After the user supplied and authorized an official DeepSeek API key, the first resumed Character/World rehearsal used `deepseek-v4-pro`; after repeated slow schema-repair turns, the user explicitly selected the current `deepseek-flash` model for the next task. All subsequent roles use `deepseek-flash`; no fallback is configured and the credential is referenced only through `DEEPSEEK_API_KEY`.
- Novel Studio implements Planner as the `writer` role, so the task's Planner/Medium requirement maps to `roles.writer.reasoning_effort = medium`.
- Architect, World Arbiter, Reviewer, and Coordinator use High; Character uses Medium.
- Architect `max_turns` is set to 20 after the first unmodified outline-all attempt reproducibly exhausted the default four turns without publishing a candidate. The failed attempt and interruption remain in the audit logs.
- Character Agent uses the baseline's current physical-state data protocol `v2`, execution policy `v3`, and a bounded maximum of four activation cycles. The earlier explicit data protocol `v1` failed rehearsal capability validation before a model call; execution policy `v1` then produced an Arbiter report proving that the TEST_SCAFFOLD World State probe lacked artifact write/read capability. Both failures are retained.
- Drafter and Editor are pinned to the same model even though Season Planning and episode rendering are outside Phase 0.
- Embedding/Qdrant and notifications are disabled for this bounded local PoC.
- A connectivity check or story generation consumes Codex subscription usage and is not performed by merely loading this file.
