# Model usage

## Runtime generation

Generation `pg2_378bb2c691e5a07d2c2b8d02` used `deepseek-flash` for Character Agent, World Arbiter, and chapter readiness.

| Role | Ledger rows | Provider attempts | Input tokens | Cache-read tokens | Output tokens | Successful rows | Failed rows |
|---|---:|---:|---:|---:|---:|---:|---:|
| Character | 14 | 28 | 433,778 | 370,688 | 84,435 | 13 | 1 |
| World Arbiter | 6 | 16 | 486,550 | 407,808 | 91,358 | 5 | 1 |
| Chapter readiness | 4 | 4 | 36,929 | 4,608 | 33,129 | 4 | 0 |

Runtime subtotal: 24 ledger rows, 48 attempts, 957,257 input tokens, 783,104 cache-read tokens, and 208,922 output tokens.

`deepseek-flash` remained unpriced in the local registry, so `monetary_cost = unavailable`; zero must not be interpreted as free. Earlier scaffold/rehearsal Codex and DeepSeek usage remains recorded in the durable project ledgers and historical report versions.

Credentials were loaded from macOS Keychain into process environment only. No API key is stored in source, artifacts, or reports.
