# Phase 0 project configuration

- The configuration is intentionally project-scoped and contains no API key.
- After the bounded DeepSeek Flash outline-all attempt failed twice at operation 3's four-turn ceiling, the user explicitly authorized `deepseek-v4-pro` for the Architect role. Other roles remain on `deepseek-flash`.
- The Coding endpoint's live `/models` catalog exposed `glm-5.3`; the user then selected it for the Architect rerun. `glm-5.3-flash` and `glm-5.2` remain available in the same provider inventory. The Coding Plan key is referenced only through `GLM_CODING_API_KEY` and remains in macOS Keychain.
- After the GLM-5.3 comparison was stopped during operation 2, the user selected `codex/gpt-5.6-sol` for the next corrected-input Architect rerun. GLM and DeepSeek providers remain configured and their credentials remain outside the repository.
- For the separately authorized P0-2 rerun, Architect was switched back to the retained `deepseek-flash` route so every model role used DeepSeek; the Keychain credential remained external to the repository.
- Scaffold construction used the Codex CLI subscription provider and `gpt-5.6-sol`. After the user supplied and authorized an official DeepSeek API key, the first resumed Character/World rehearsal used `deepseek-v4-pro`; after repeated slow schema-repair turns, the user explicitly selected the current `deepseek-flash` model for the next task. All subsequent roles use `deepseek-flash`; no fallback is configured and the credential is referenced only through `DEEPSEEK_API_KEY`.
- Novel Studio implements Planner as the `writer` role, so the task's Planner/Medium requirement maps to `roles.writer.reasoning_effort = medium`.
- Architect, World Arbiter, Reviewer, and Coordinator use High; Character uses Medium.
- Architect `max_turns` is set to 20 after the first unmodified outline-all attempt reproducibly exhausted the default four turns without publishing a candidate. The failed attempt and interruption remain in the audit logs.
- Character Agent uses the baseline's current physical-state data protocol `v2`, execution policy `v3`, and a bounded maximum of four activation cycles. The earlier explicit data protocol `v1` failed rehearsal capability validation before a model call; execution policy `v1` then produced an Arbiter report proving that the TEST_SCAFFOLD World State probe lacked artifact write/read capability. Both failures are retained.
- Drafter and Editor are pinned to the same model even though Season Planning and episode rendering are outside Phase 0.
- Embedding/Qdrant and notifications are disabled for this bounded local PoC.
- A connectivity check or story generation consumes Codex subscription usage and is not performed by merely loading this file.
