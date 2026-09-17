# P0-5R Final Deterministic GitHub Archive

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

```text
BASELINE/AUDIT HEAD BEFORE TASK: 2a22da48ea6fc96224772b2b28a100b4e6caf6a2
CORE: 215228e1c521eea6965c9c892a034cad4223604c
EVIDENCE: this report's commit; exact SHA reported after commit/push
DEVELOPMENT BRANCH: codex/p0-5r-final-deterministic-closure
AUDIT BRANCH: codex/jiujiu-phase0-audit
REMOTE: https://github.com/lig9904/novel-studio.git
UPSTREAM PR: NOT CREATED
```

The Core commit contains only generic deterministic tests for bundle-tail
authority and activation Grounding surface/receipt wiring. The P0-specific
controlled harness update, reports, and machine-readable audit records remain
in Evidence.

Exact Core verification:

```text
go test ./... -count=1: PASS
go vet ./...: PASS
go build ./...: PASS
targeted cmd/domain race: PASS
```

Combined Evidence working-tree verification before commit:

```text
go test ./... -count=1: PASS
go vet ./...: PASS
go build ./...: PASS
targeted cmd/domain race: PASS
git diff --check: PASS
independent review: PASS
```

Archived change scope:

- generic bundle-tail authority matrix test;
- generic activation Grounding offscreen surface/receipt test;
- P0-5 controlled harness reuse of `pipelineProjectAllTail`;
- deterministic bundle-genesis, Grounding-coverage, fox-authority,
  independent-review, routing, and archival reports;
- machine-readable audit JSON and checksum manifest.

No frozen Story Simulation, Character activation evidence, outline, layered
outline, official input, or retained model result was modified. No paid Planner,
Grounding, Character Agent, World Arbiter, or Story Simulation was run.

```text
LOCAL_VERIFIED: YES
REMOTE_CI_VERIFIED: NO
PAID MODEL RUNTIME: NO
SECRETS CHECK: PASS
```

Remote development and audit refs are verified after the Evidence commit is
pushed; their exact shared SHA is returned in the final task result.
