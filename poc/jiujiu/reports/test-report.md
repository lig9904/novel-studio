# Test report

## Required baseline tests

| Check | Result | Evidence |
|---|---|---|
| `go test -count=1 ./...` attempt 1 | FAIL | Two tests exceeded the fixed 5s Codex MCP inventory timeout under full package parallelism. |
| `go test -count=1 ./...` attempt 2 | FAIL | Same two tests, same timeout signature. |
| Both failed tests in isolation | PASS | 0.50s and 0.75s without source changes. |
| Full suite with package serialization (`-p 1`) | PASS | All packages passed without source changes. |
| `go vet ./...` | PASS | No diagnostics. |

The two exact-command failures are not erased. The isolated tests and complete serialized run prove the code paths pass when the host is not running every package concurrently. Classification: **environmental scheduling sensitivity / fixed 5s subprocess timeout**, not a repaired or masked core failure.

## CI checks

| Check | Result |
|---|---|
| Vendored `third_party/litellm` Race suite | PASS |
| Agents Race sharding contract | PASS |
| Store Race sharding contract | PASS |
| `go test -race ... ./internal/tools ./services/dashboard` | PASS |
| Store Race shard 0/2 | PASS (164/327 top-level tests) |
| Store Race shard 1/2 | PASS (163/327 top-level tests) |
| Agents Race shard 0/4 | PASS (59/236 top-level tests) |
| Agents Race shard 1/4 | PASS (59/236 top-level tests) |
| Agents Race shard 2/4 | PASS (59/236 top-level tests) |
| Agents Race shard 3/4 | PASS (59/236 top-level tests) |

The contract scripts initially failed before testing because macOS Bash 3.2 lacks `mapfile`. They passed unchanged under Homebrew Bash 5.3.20, matching the GNU Bash behavior of the Ubuntu CI runner.

Raw evidence is under `poc/jiujiu/logs/`.

Current build/test gate: **PASS WITH RECORDED ENVIRONMENTAL CAVEAT**. No core source file was changed.

## TEST_SCAFFOLD Story Engine run

| Check | Result |
|---|---|
| Architect foundation | PASS after schema-correcting retries |
| outline-all | PASS in isolated TEST_SCAFFOLD, with recoverable operation receipts |
| zero-init and initial world tick | PASS |
| structured v2 physical resources | PASS |
| Phoenix minimum-knowledge input | PASS |
| fox-private resource omitted from 九九/Phoenix views | PASS |
| Arc rehearsal tool execution | PASS (draft and World Arbiter report stored) |
| Arc rehearsal readiness | FAIL (`ready_for_detail=false`) |
| direct project-all attempt | FAIL CLOSED before model call |
| Character Agent runtime | INCONCLUSIVE / not dispatched |
| seal | NOT RUN because project-all did not complete |
| promote/render | NOT RUN by authorization |
| rebase source/archive root equality | PASS |
| final pending usage calls | PASS (none) |

The final report digest is `sha256:b700daa72c2d853ec872ec21fcc04f186d817906341e326d359581883e38611b`. Its summary says the selected conditional path has no hard-contract reachability failure, but the durable gate remains not ready with 14 unresolved items and one missing material check. The system therefore cannot proceed to detailed Character/World simulation without a product fix or a different source contract.
# Runtime continuation update — 2026-09-16

The later authorized continuation supersedes the earlier “Character Agent not executed” status:

- Rehearsal Gate: PASS (`ready_for_detail=true`, report `41434f…3252`);
- Character Agent: EXECUTED for four candidate activation cycles;
- Character Autonomy: PASS;
- Knowledge Boundary: FAIL because global character-view rules injected forbidden identity facts into Phoenix's real observation packet;
- Proposal Isolation: INCONCLUSIVE for the exact PHOENIX-A/PHOENIX-B branch test;
- World State: PASS for candidate execution;
- Recovery: PASS;
- project-all chapter 1: incomplete at the four-cycle readiness ceiling;
- Promote/Render/Accepted Canon: not run / unchanged.
