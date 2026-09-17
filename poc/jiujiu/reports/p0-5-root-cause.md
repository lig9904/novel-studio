# P0-5 Simulation → Plan Projection Root Cause

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Result: **CONFIRMED**

Primary classification: **PLANNER MODEL-FACING AUTHORITY CONTRACT GAP**

Runtime-contributing defects: **GROUNDING POSITIVE-CHANNEL COVERAGE GAP** and **STALE GROUNDING RECEIPT REUSE GAP**

Domain conclusion: **existing simulation, projection, evidence and receipt structures are sufficient; no parallel Authority system is required.**

## Required answers

### 1. Planner's actual authority

The intended contract is projection-only. Final `ChapterWorldSimulation` and arbitration decide what happened. Planner may select, order, group, emphasize, express and organize those facts into scenes. It may add removable sensory or atmospheric rendering, but it may not create story reality.

### 2. Was adding a Story Fact explicitly forbidden?

Only partially before P0-5. Project-Arc called simulation the unique causal input, activation planning said to organize actual events, and Grounding said Soft Outline could not authorize unoccurred events. The Planner-facing prompt and fact-bearing schema descriptions did not define an operational Story Fact / Derived Presentation boundary, causal-persistence deletion test, or explicit categories for new resources, entities, knowledge, events and obligations.

### 3. Origin of the three unsupported facts

`浮标绳`, `水尺` and `现场来客` first appeared in the pre-simulation Soft Outline. They are absent from the official user input and final simulation. Readiness closed that pressure through `REJECTED_WITH_CONSEQUENCE`; it did not make the rejected candidates occur. Planner later re-promoted them into a plan.

### 4. Why Grounding found them while Planner produced them

The historical P0-4 Grounding packet had an explicit rule against treating Soft Outline as occurrence evidence. Planner instead received a large creation-oriented context and schemas asking for environment state, causal beats, dialogue, scene actions and consequences without an equally concrete authority boundary.

P0-5 runtime exposed a second path: after the three headline facts were removed from ordinary plan beats, unsupported rope/gauge/channel facts still survived in positive `external_reference_plan` fields. The old Grounding prompt did not explicitly enumerate every positive Drafter-consumed side channel, so it returned a historical `pass=true`. Under the strengthened current protocol, a direct replay rejects the same plan with five exact findings.

### 5. Prompt, Schema, Domain or Authority encoding?

- Primary cause: Planner prompt/model-facing authority encoding.
- Contributing cause: fact-bearing tool-schema descriptions lacked the boundary.
- Runtime P1: Grounding omitted positive RAG/reader/longform/literary plan surfaces from its explicit checklist.
- Runtime P1: Project-All could reuse a passing receipt created under an older Grounding protocol.
- Not the cause: Domain already defines simulation as source of truth, POV projection as the renderable slice, and literary rendering as non-obligation-creating presentation.

### 6. Minimal generic fix

1. Define final simulation/arbitration as Story Authority.
2. Demote unmatched Soft Outline content to scope/pressure/candidate status.
3. Permit selection, ordering, grouping, emphasis, expression and removable non-causal presentation.
4. Require an authority source for a new entity, resource, character, knowledge transfer, event, state, result or future obligation.
5. Apply the causal-persistence deletion test before submission.
6. Reinforce the same rule in fact-bearing tool descriptions.
7. Make Grounding inspect every positive Drafter-consumed plan channel while excluding source and negative metadata from occurrence claims.
8. Reject passing receipts whose protocol identity does not match the current reviewer protocol.

The fix is project-neutral and model-neutral. It does not modify Character decisions, World arbitration, Canon or the frozen Story Simulation; it does not weaken Grounding or turn Planner into a World Simulator.

## Runtime disposition

The Root Cause is fixed at code/protocol level and independently reviewed. Runtime acceptance remains **BLOCKED** because the allowed one normal Planner run plus one confirmation did not yield a plan that passes the strengthened current Grounding protocol. No Chapter 1 bundle was created.
