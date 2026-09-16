# Model usage

Status at blocker stop: **connectivity check completed; no story-generation call started**.

| Role | Model | Effort | Calls | Input tokens | Cached tokens | Output tokens | Duration | Estimated cost | Failure | Retry |
|---|---|---:|---:|---:|---:|---:|---:|---|---|---:|
| Connectivity check (shared target) | gpt-5.6-sol | xhigh* | 1 | 10,333 | 0 | 5 | 9.763s | monetary_cost = unavailable | none | 0 |
| Coordinator | gpt-5.6-sol | high | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | not started | 0 |
| Architect | gpt-5.6-sol | high | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | not started | 0 |
| Character Agent | gpt-5.6-sol | medium | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | not started | 0 |
| World Arbiter | gpt-5.6-sol | high | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | not started | 0 |
| Planner (`writer`) | gpt-5.6-sol | medium | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | not started | 0 |
| Reviewer | gpt-5.6-sol | high | 0 | 0 | 0 | 0 | 0s | monetary_cost = unavailable | not started | 0 |

The repository's unit tests use fake Codex executables and synthetic usage fixtures; those are not counted as story-model calls.

\* The repository's `--check` implementation probes the unique provider/model target once with its own check effort. This does not change the configured production-role efforts.

The run uses a logged-in Codex Pro subscription rather than an API key. Accurate monetary attribution for that subscription call is unavailable, so no API list-price estimate is presented as actual cost.
