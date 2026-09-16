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
