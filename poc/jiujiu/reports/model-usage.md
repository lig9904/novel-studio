# Model usage

Status at Phase 0 stop: **scaffold/rehearsal calls completed; Character Agent calls never started**.

| Role | Model | Effort | Calls | Input tokens | Cached tokens | Output tokens | Duration | Estimated cost | Failure | Retry |
|---|---|---:|---:|---:|---:|---:|---:|---|---|---:|
| Scaffold aggregate | gpt-5.6-sol | high/medium | 64 recorded | 2,363,252 | 356,480 | 170,846 | individual calls 40–195s; aggregate unavailable | ledger estimate USD 11.3652346; subscription charge unavailable | MCP timeout, capacity, schema rejection | multiple |
| Rehearsal aggregate | deepseek-v4-pro | high | 3 recorded | 130,913 | 95,104 | 53,590 | connectivity 7.019s; long rehearsal >15m before interruption | ledger USD 0.062544967 | malformed/tool-contract retries | multiple |
| Scaffold/rehearsal aggregate | deepseek-flash | high/medium | 29 recorded | 1,148,649 | 853,888 | 319,706 | connectivity 1.237s; long rehearsal about 3–8m | monetary_cost = unavailable; ledger incorrectly records zero | schema rejection; final report not ready | multiple |
| Unknown/synthetic accounting rows | unavailable | unavailable | 4 recorded | included in overall | included | included | unavailable | monetary_cost = unavailable | missing model attribution | n/a |
| Character Agent | not dispatched | medium | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | blocked before project-all | 0 |
| Planner (`writer`) | not dispatched | medium | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | blocked before project-all | 0 |
| Reviewer | not dispatched | high | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | outside reached stages | 0 |

Final durable overall ledger: input 3,642,814; output 544,142; cache read 1,305,472; recorded cost USD 11.427779567. This total is incomplete because the current `deepseek-flash` alias was unpriced and 7 assistant usages were missing.

The repository's unit tests use fake Codex executables and synthetic usage fixtures; those are not counted above.

\* The repository's `--check` implementation probes the unique provider/model target once with its own check effort. This does not change the configured production-role efforts.

GPT calls used the logged-in Codex Pro subscription; DeepSeek calls used the explicitly authorized official API key through an environment variable. No credential was committed or printed to a report. Aggregate Phase 0 wall time was approximately 2 hours 5 minutes including retries, inspections, rebuilds, and human-directed model switches.
