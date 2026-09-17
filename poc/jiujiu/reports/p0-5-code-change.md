# P0-5 Code Change

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Status: **IMPLEMENTED, TESTED, RUNTIME ACCEPTANCE BLOCKED**

## Planner authority contract

- `assets/prompts/planner.md`
  - final simulation/arbitration owns Story Facts;
  - Soft Outline is scope/pressure/candidate input after simulation;
  - Planner may select/order/group/emphasize/express/organize and add removable non-causal rendering;
  - causal-persistent additions require an authority source;
  - RAG/craft/web may realize existing facts but cannot authorize a new event, person, resource, device, communication, knowledge or result.
- `internal/agents/project_all.go`
  - applies the same authority contract to Project-Arc planning.
- `internal/agents/character_activation_protocol.go`
  - forbids revival of unrealized Soft Outline candidates.
- `internal/tools/plan_chapter_phases.go` and `internal/tools/plan_chapter.go`
  - reinforce authority in fact-bearing structure and detail fields.

## Runtime-discovered hardening

- `internal/agents/plan_grounding.go`
  - checks positive external-reference, grounding, reality-support, reader, longform and literary-rendering channels;
  - says provenance receipts do not prove story occurrence;
  - excludes source/negative metadata from occurrence claims;
  - requires real JSON pointers and exact quotes.
- `internal/agents/project_all.go`
  - verifies a saved passing Grounding receipt against the current reviewer protocol before bundle construction;
  - fails closed when the receipt is absent, failed or stale.

## Tests and scaffold

- Authority contract tests: `internal/agents/plan_projection_authority_test.go`, `internal/tools/plan_projection_authority_test.go`, `assets/load_test.go`.
- Stale receipt gate: `internal/agents/plan_grounding_reuse_test.go`.
- Opt-in frozen runtime and direct Grounding regression: `cmd/novel-studio/pipeline_p0_5_live_test.go`.
- The fixture-specific scanner checks only positive Drafter-consumed fields and excludes source/forbidden/negative metadata.

## Preserved boundaries

- no Domain enum, persistent schema or migration;
- no Character Decision, World Arbitration, Readiness, Canon or Story Simulation change;
- no Host source-binding relaxation;
- no 九九-specific or model-specific production branch;
- no Chapter 2, seal, promote, render or Canon publication.

Protocol digest snapshot updates are intentional. Code review is PASS; runtime acceptance is not.
