# Novel Studio Codex Instructions

## Scope

- Apply these instructions to this repository.
- Do not treat Codex routing work as authorization to change Story Engine business code.
- Preserve existing worktree changes and untracked runtime evidence.
- Keep the root agent on the session-selected model. The expected project default is GPT-5.6 Sol with high reasoning; do not change global Codex configuration from this repository.

## Model Routing & Delegation Policy

Choose the cheapest model that can reliably complete the task. Reliability takes priority over cost. The root agent owns task classification, root-cause analysis, architecture, protocols, state machines, core Story Engine judgments, integration of subagent results, the final implementation decision, and the final technical conclusion.

Classify each material task before delegating:

### LEVEL 0 — Mechanical

Examples: run existing tests or builds, collect logs and stdout/stderr, index evidence, summarize a diff, list files, produce a fixed-format table or Markdown report, and perform repetitive rule checks.

- Delegate to `worker` using GPT-5.6 Luna with medium reasoning.
- On the current collaboration interface, pass `model = "gpt-5.6-luna"` and `reasoning_effort = "medium"` explicitly when spawning the worker.
- Prefer `fork_turns = "none"` and provide the working directory, exact commands or files, and required output.
- The worker may report failures but must return to the root instead of changing Story Engine protocols or business logic to make a check pass.

### LEVEL 1 — Standard Engineering

Examples: repository exploration, call-path tracing, a single-file bug, a clear interface, ordinary unit tests, a simple adapter, or non-core refactoring analysis.

- Delegate bounded exploration and analysis to `explorer` using GPT-5.6 Terra with medium reasoning.
- On the current collaboration interface, pass `model = "gpt-5.6-terra"` and `reasoning_effort = "medium"` explicitly when spawning the explorer.
- Prefer two to four recent turns of context. Use no inherited turns when a complete standalone task contract is sufficient.
- The root remains responsible for decisions that cross modules or alter authority, state, recovery, or concurrency semantics.

### LEVEL 2 — Complex or Core Engineering

Examples: cross-module root cause, pipeline state machines, Canon Authority, Knowledge Boundary, Character Agent or Character Autonomy, World Arbiter, Recovery, Proposal Isolation, Human Approval, Evidence Authority, concurrency, races, and core pipeline protocols.

- Keep primary ownership in the root using GPT-5.6 Sol with high reasoning.
- Use `explorer` for bounded read-heavy investigation and `worker` for mechanical verification.
- Do not accept a subagent conclusion as final evidence; the root must inspect and integrate it.

### LEVEL 3 — Difficult Core Problem

Escalate only when Sol/high has failed to resolve the issue, runtime evidence conflicts with unit tests, several root causes remain plausible, a race is difficult to locate, recovery consistency is unclear, authority boundaries intersect, or high-risk data contamination is possible.

- First raise the root to GPT-5.6 Sol with xhigh reasoning for the difficult phase.
- If the current root cannot change effort safely, use an independent or forked task for the deep analysis.
- Record the reason for escalation.

### LEVEL 4 — Exceptional Review

Use `reviewer` with GPT-6 Astra/high only for an independent review of core changes, repeated failed fixes, unresolved Sol/xhigh analysis, exceptionally complex cross-system behavior, or a potentially false architecture assumption. Use Astra/xhigh only when concrete evidence justifies it.

On the current collaboration interface, pass `model = "gpt-6-astra"` and `reasoning_effort = "high"` explicitly when spawning the reviewer. Pass `xhigh` only for a separately recorded escalation.

Reviewer-triggering areas include:

- Canon Authority
- Character Agent or Character Autonomy
- Knowledge Boundary
- World Arbiter
- Pipeline state machines
- Recovery
- Proposal Isolation
- Human Approval
- data migration
- concurrency and race-sensitive code
- fail-closed behavior and core protocol changes

Do not use the reviewer for ordinary documentation, fixtures, formatting, small bugs, or simple adapters.

## Agent Responsibilities

### Current host compatibility

The current `spawn_agent` contract exposes `task_name`, `fork_turns`, `model`, `reasoning_effort`, and the task message, but it does not expose a custom-agent selector. A task name that matches `.codex/agents/<name>.toml` does not by itself apply that file's model or effort.

Therefore every routed spawn must:

1. use the matching task name (`explorer`, `worker`, or `reviewer`);
2. pass the model and reasoning effort explicitly;
3. include the relevant responsibility and escalation boundary in the task message;
4. verify the actual child model and effort from runtime thread metadata when producing an important routing report.

The project TOML files remain the canonical reusable agent definitions for Codex clients that expose custom-agent selection. Do not report their configured values as runtime evidence unless the runtime metadata agrees.

### explorer

Use for codebase exploration, call-chain tracing, file and symbol location, impact analysis, test discovery, schema/API analysis, ordinary implementation options, routine refactoring analysis, and non-core review.

The explorer must return evidence to the root and must not independently approve Canon Authority, Character Autonomy, Knowledge Boundary, World Arbiter, Recovery, Proposal Isolation, Human Approval, or core pipeline state-machine decisions.

### worker

Use for existing test, `go test`, `go vet`, `go build`, race-test execution, stdout/stderr collection, log organization, evidence indexing, diff and file summaries, fixed-template transformations, and repetitive verification.

When a failure requires complex code interpretation or business-logic changes, the worker must report `RETURN TO ROOT` with the failing command and evidence.

### reviewer

Use as an independent high-risk reviewer. Give it the task contract, relevant diff, tests, runtime evidence, and known risks. Ask it to reach its own conclusion and actively search for counterexamples, authority leaks, special-case patches, missing cases, and test blind spots. Do not prime it by claiming the implementation is already correct.

## Escalation and De-escalation

- Luna to Terra or root: unclear test failure, conflicting evidence, complex code explanation, or any required business-logic change.
- Terra to root Sol: cross-module behavior, authority, state machine, Canon, Knowledge, Recovery, race, or uncertain root cause.
- Sol/high to Sol/xhigh: first repair failed, runtime and tests conflict, or a high-risk protocol still has several plausible interpretations.
- Sol/xhigh to Astra: only for a genuinely difficult independent review.
- After a high-capability model makes the risky decision, delegate mechanical follow-up back to Terra or Luna.

## Concurrency and Shared Files

The normal capacity is the root plus three concurrent subagents. Do not fill every slot without useful independent work.

- Parallel read-only exploration and evidence organization are allowed.
- Judge whether parallel tests contend for ports, databases, caches, or generated artifacts before starting them.
- Parallel edits to the same file are prohibited.
- Parallel edits to different files require explicit ownership boundaries and root integration.
- Run the reviewer after implementation and primary verification unless an earlier independent risk assessment is needed.

## Important Task Reporting

Include this block in reports for material engineering work:

```text
MODEL ROUTING

Task Level:
Primary:
Delegated:
Escalations:
Reviewer:
Reason:
```

Report the models and reasoning efforts actually used. Distinguish configured intent from observed runtime evidence.
