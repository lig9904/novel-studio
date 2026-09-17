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

## Phase Task GitHub Archival Policy

Apply this policy after a Phase task reaches its declared terminal state and before work starts on the next Phase task. Archival preserves the verified implementation and evidence; it does not reopen the task, rerun story generation, or change any Canon authority.

### Audit baseline and branch

- Build each Phase audit branch from the latest verified Evidence commit for that Phase chain. For 九九 Phase 0, the unified branch is `codex/jiujiu-phase0-audit` and its initial verified baseline is P0-5 Evidence commit `92e9b5767aefa8b1e987677acb44fb87e58312af`.
- Prefer an isolated clean worktree for archival. Preserve development branches, dirty worktrees, untracked runtime evidence, and their provenance.
- Before creating or pushing an audit branch, inspect local worktrees, branch heads, base SHAs, commit graph, remotes, tracked changes, untracked files, and existing remote refs.
- If the target remote branch already exists and was not created by the current archival chain, or its head is not the expected ancestor, stop and report the actual graph. Never overwrite it or force push.
- Keep the audit history linear when the verified sources permit it. Do not rebase, squash, or discard frozen evidence merely to make the graph look linear.

### Commit separation

For every Phase task, preserve reviewable commit boundaries in this order:

```text
verified prior Evidence
        ↓
Phase Task CORE      (when generic code changed)
        ↓
Phase Task EVIDENCE
```

- `CORE` contains only generic Story Engine implementation, protocols, schemas, migrations, and their generic tests.
- `EVIDENCE` contains PoC fixtures, manifests, reports, controlled runtime results, project-specific configuration, and project-specific test scaffolds.
- Never place a 九九-specific live/runtime scaffold in a generic Core commit.
- Do not manufacture an empty Core commit. Record `CORE: NO CODE CHANGE REQUIRED` when the task changed no generic code.
- Do not squash Core and Evidence together. Use concise commit messages that identify the task and layer.

### Explicit staging and scope control

- Never use `git add .` or `git add -A` for Phase archival.
- Stage only an explicit reviewed file list. Inspect `git diff --cached --name-status`, `git diff --cached --stat`, and `git diff --cached --check` before every commit.
- Exclude unrelated work, runtime workspaces, archives, session transcripts, locks, caches, headless logs, build outputs, and temporary generated files.
- If source worktrees contain mixed or unsafe changes, construct the audit chain in a separate worktree by copying only reviewed files from the verified sources. Do not mutate or clean the source worktrees to simplify archival.

### Evidence authority

Every PoC evidence set and its reports must preserve this authority boundary:

```text
TEST_SCAFFOLD
NOT ACCEPTED CANON
NOT HUMAN APPROVAL
```

- Git commit, GitHub push, fixture retention, model verdict, or test PASS does not promote story content into official Canon.
- Human approval and the normal Canon promotion path remain separately required.
- Preserve content-addressed fixture hashes. If public archival requires a path-only privacy redaction, retain the original content digests, document the normalization, recompute only the enclosing checksum manifest, and do not rerun an evidence-producing process.

### Security and privacy gate

Before every archival commit and push, verify that the staged range contains no:

- API keys, access tokens, refresh tokens, passwords, credentials, private keys, or Keychain contents;
- secret `.env` values or provider configuration with literal credentials;
- unnecessary personal information in absolute local paths;
- Codex session transcripts, local thread/session identifiers, or rollout files;
- unrelated logs, databases, locks, caches, runtime archives, or build artifacts.

Environment-variable names, explicit test-only placeholders, and documented redacted examples are allowed only when they cannot authenticate to any service.

### Verification and push

- Do not rerun paid models, Character Agents, Planner, Grounding, Story Simulation, rendering, or media generation merely to archive Git history.
- Record the latest valid `go test ./...`, `go vet ./...`, and `go build ./...` results for the exact Core tree being pushed. Run only deterministic checks needed for changes made after that verification.
- Run race checks only when the archived Core changes concurrency or race-sensitive behavior; otherwise report `RACE: NOT REQUIRED` with the reason.
- Verify frozen fixture checksums and the official project-input digest when those artifacts are in scope.
- Push only to the authorized `origin`. Do not create an upstream PR unless the user separately requests it.
- Use a normal non-force push. After pushing, read the remote ref and require its SHA to equal the local audit branch head.

### Required archival report and stop boundary

Report the verified baseline, each Core/Evidence commit SHA, audit branch, remote URL, remote head SHA, deterministic check results, race disposition, secrets check, actual changed files, and whether an upstream PR was created. Show the real commit graph when it differs from the preferred linear form.

After the remote SHA is verified, report:

```text
PHASE TASK GITHUB ARCHIVE READY
WAITING FOR HUMAN REVIEW
```

Then stop. Do not begin the next Phase task, another chapter, rendering, promotion, or an upstream PR as part of archival.
