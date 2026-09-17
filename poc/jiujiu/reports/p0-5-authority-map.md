# P0-5 Simulation → Plan Authority Map

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Status: **READ-ONLY DIAGNOSIS**

## MODEL ROUTING

```text
Task Level: LEVEL 2 — Complex / Core Engineering
Primary: /root, GPT-5.6 Sol / High
Delegated: /root/explorer, GPT-5.6 Terra / Medium; /root/worker, GPT-5.6 Luna / Medium
Escalations: none in read-only phase
Reviewer: not used before implementation
Reason: root owns authority semantics; Explorer traced the protocol; Worker indexed frozen evidence.
```

The routed models and efforts were verified from runtime `turn_context` metadata after the read-only phase. The later core-change review used GPT-6 Astra / High and is recorded in `p0-5-summary.md`.

## Authority map

`Planner Modify/Create` below means authority over story reality, not the ability to write a representation into a plan.

| Data | Authority Owner | Planner Read | Planner Modify | Planner Create |
|---|---|---:|---:|---:|
| Canon | Human-approved Canon authority and protected project sources | Yes, projected/scoped view | No | No |
| Character State | Character Agent result + World Arbiter + Host physical-state finalization | Yes, POV projection and bounded continuity view | No | No |
| Character Knowledge | Observation/reception evidence + World Arbiter + Host knowledge timing | Yes, author-side trace may include all actors' scoped facts; POV plan may present only protagonist-known/legally obtained facts | No | No |
| World State | Final `ChapterWorldSimulation` / `WorldPhysicalStateV2` | Yes, projected view | No | No |
| Character Decision | Character Agent proposal accepted by World Arbiter and bound by Host | Yes | No | No |
| Arbitration Result | World Arbiter receipt validated and bound by Host | Yes | No | No |
| Resource | Canon/World State/resource ledger and arbitrated state changes | Yes | No | No |
| Character | Canon/character registry and simulation authority entry | Yes | No | No |
| Event | Final simulation/arbitration for actual events; Soft Outline only proposes scope/pressure | Yes | No actual event | No actual event |
| Presentation Detail | Planner, within the final Story Facts and POV boundary | Yes | Yes, in projection only | Yes, only non-causal derived presentation |

## Existing authority contracts

- `ChapterWorldSimulation` is explicitly the prewriting source of truth; the POV plan derives from `ProtagonistProjection`: `internal/domain/chapter_simulation.go:11-32`.
- `ProtagonistDecisionProjection` is the only simulation slice a POV plan may render directly: `internal/domain/chapter_simulation.go:152-163`.
- The activation Planner boundary says Planner must consume all arbitrated cycles and “只组织已发生的真实内容”: `internal/agents/character_activation_protocol.go:9-10`.
- When an independent simulation exists, Host removes the exact injected Soft Outline beat and binds the exact simulation ID and chosen decision: `internal/tools/plan_chapter.go:496-528`, `internal/tools/plan_chapter_phases.go:212-257`.
- `LiteraryRenderingPlan` already means rendering an established event without adding plot obligations: `internal/domain/writing.go:128-183`.
- `PlanGroundingReceipt` binds the exact plan to the exact simulation/evidence input. Grounding checks time, location, knowledge, intent and outcome while the Domain finalizer validates exact pointers and quotes: `internal/agents/plan_grounding.go:20-203`, `internal/domain/plan_grounding.go:120-205`.

## Final Planner authority contract

Planner may:

```text
SELECT
ORDER
GROUP
EMPHASIZE
EXPRESS
ORGANIZE SCENES
ADD NON-CAUSAL SENSORY / ATMOSPHERIC RENDERING
```

Planner may not decide or introduce:

```text
NEW ACTUAL EVENT
NEW CHARACTER OR ENTITY
NEW RESOURCE OR DEVICE
NEW KNOWLEDGE OR COMMUNICATION
NEW STATE CHANGE
NEW CAUSAL CAPABILITY
NEW FUTURE STORY OBLIGATION
```

Derived presentation is allowed only when deleting it cannot change any later Story Simulation result. It must not persist as an entity/resource/fact, change knowledge/state/result, or create an obligation.

## Source binding

No parallel authority system is required. Existing bindings remain authoritative:

- simulation ID/digest and authority receipt;
- arbitration/evidence digests;
- resource and character IDs;
- planning context source token;
- `context_sources`, `SourceBindingV2`, RAG/craft exact refs;
- exact `PlanGroundingReceipt` over the complete plan/evidence input.

A fact-bearing Plan claim without support in those sources must be rejected or revised. A derived presentation detail may be model-authored only under the non-causal deletion test above.
