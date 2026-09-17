# Codex Model Routing Policy

## Purpose

This repository uses project-scoped Codex instructions and custom agents to route work by risk and complexity. The policy does not change Story Engine behavior and does not require global Codex configuration.

The governing rule is: use the cheapest model that can reliably complete the task, while placing reliability ahead of cost.

## Routing Matrix

| Level | Work | Primary route | Expected model and effort |
| --- | --- | --- | --- |
| LEVEL 0 | Tests, builds, logs, evidence, inventories, diff summaries, fixed-format output, repetitive checks | `worker` | GPT-5.6 Luna / medium |
| LEVEL 1 | Exploration, call paths, symbols, single-file bugs, clear interfaces, ordinary test discovery and non-core analysis | `explorer` | GPT-5.6 Terra / medium |
| LEVEL 2 | Cross-module root cause, authority, state machines, recovery, concurrency and core Story Engine protocols | root | GPT-5.6 Sol / high |
| LEVEL 3 | Failed first repair, conflicting runtime and unit evidence, unresolved races or competing high-risk explanations | root escalation | GPT-5.6 Sol / xhigh |
| LEVEL 4 | Independent review of difficult or high-risk core changes | `reviewer` | GPT-6 Astra / high; xhigh only with evidence |

The root owns task classification, decomposition, core judgments, integration, the final implementation decision, and the final technical conclusion. Subagent results are inputs to that judgment rather than independent approval.

## Project Agents

### explorer

`explorer` is a read-heavy GPT-5.6 Terra/medium agent. It locates files and symbols, traces call paths, maps impact, examines schemas and APIs, identifies tests, and analyzes ordinary implementation or refactoring options.

It returns Canon Authority, Character Autonomy, Knowledge Boundary, World Arbiter, Recovery, Proposal Isolation, Human Approval, Evidence Authority, concurrency, race, and core pipeline-state decisions to the root.

### worker

`worker` is a GPT-5.6 Luna/medium mechanical agent. It runs specified existing checks, captures output and exit status, organizes evidence, inventories files, summarizes diffs, and produces fixed-format reports.

A failing check is a valid result. When the cause is unclear or a business-logic change would be required, it reports `RETURN TO ROOT` instead of changing a core protocol.

### reviewer

`reviewer` is a read-only GPT-6 Astra/high independent reviewer. It is reserved for core authority, state-machine, recovery, migration, concurrency, race-sensitive, fail-closed, and protocol changes. It reviews the task contract, diff, tests, runtime evidence, and known risks without assuming that the root conclusion is correct.

## Context Policy

- Worker: normally no inherited turns. Supply the exact directory, commands or files, and output format.
- Explorer: normally two to four recent turns; use no inherited turns when the task contract is self-contained.
- Reviewer: normally no inherited turns. Supply a neutral task contract, relevant diff, tests, runtime evidence, and known risks to reduce confirmation bias.

The current Codex host permits model and reasoning overrides when a subagent receives no history or a bounded number of recent turns. A full-history spawn inherits the parent model and reasoning effort.

The current host does not expose a custom-agent selector in `spawn_agent`. Matching `task_name` to a file under `.codex/agents/` is insufficient to apply the file's model and reasoning effort. The root must pass the routed model and effort explicitly and include the relevant role boundary in the child task. The TOML definitions remain usable by Codex clients that expose custom-agent selection.

## Concurrency

The expected capacity is one root and up to three concurrent subagents. Capacity is a limit, not a target.

Parallel reads are suitable. Before running tests concurrently, check for shared ports, databases, caches, and generated artifacts. Agents must not edit the same file concurrently. The reviewer normally runs after implementation and primary verification.

## Escalation

- Luna to Terra or root when failures are unclear, evidence conflicts, code interpretation becomes complex, or business logic must change.
- Terra to root Sol for cross-module work, authority, state machines, Canon, Knowledge, Recovery, races, or uncertain root cause.
- Sol/high to Sol/xhigh after a failed repair, runtime/test conflict, or unresolved high-risk protocol interpretation.
- Sol/xhigh to Astra only for a genuinely difficult independent review.

After a high-capability model resolves the risky decision, mechanical follow-up returns to Terra or Luna.

## Evidence and Reporting

Important task reports record:

```text
MODEL ROUTING

Task Level:
Primary:
Delegated:
Escalations:
Reviewer:
Reason:
```

Reports distinguish configured routing from observed runtime model, effort, status, and output. A saved TOML file alone is not runtime proof.

## Configuration Files

- `AGENTS.md`: routing, escalation, concurrency, and reporting policy.
- `.codex/agents/explorer.toml`: Terra/medium read-heavy explorer.
- `.codex/agents/worker.toml`: Luna/medium mechanical worker.
- `.codex/agents/reviewer.toml`: Astra/high independent reviewer.
