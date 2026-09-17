# P0-5 Projection Authority Design

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

## Goal

```text
FINAL STORY SIMULATION
        ↓ authoritative facts / state / knowledge / outcomes
POV PLAN PROJECTION
        ↓ selection / order / grouping / emphasis / expression
DERIVED PRESENTATION
```

## Reused contracts

No new authority enum or ledger is introduced.

- `ChapterWorldSimulation`: chapter source of truth.
- `ProtagonistDecisionProjection`: directly renderable POV slice.
- `CharacterActivationDecisionTrace`: actual multi-cycle chronology.
- `WorldPhysicalStateV2`, resource/character IDs and knowledge receipts: durable authority.
- `LiteraryRenderingPlan`: presentation of established events without new plot obligations.
- `SourceBindingV2`, `context_sources` and planning access receipts: provenance.
- `PlanGroundingReceipt`: exact-plan review bound to exact evidence and reviewer protocol.

## Story Fact / Presentation Detail boundary

Fact-bearing content can persist, be held, used, measured, tracked, serve as evidence, transfer knowledge, trigger action, change a result or create a future obligation. It requires an authoritative source.

Derived presentation is wording, visualization, sensory texture or atmosphere that can be deleted without changing Canon, knowledge, state, entities, resources, events, outcomes or obligations.

Planner keeps creative authority over scene grouping, emphasis, pacing, viewpoint distance and non-causal expression. Multiple real facts may share a scene; Planner may not invent a causal link between them.

## Enforcement layers

1. `assets/prompts/planner.md` defines the authority boundary and deletion test.
2. Project-Arc and activation boundaries demote unmatched Soft Outline content.
3. Fact-bearing `plan_structure` and `plan_details` descriptions require authoritative support.
4. Existing Host anchors bind the exact simulation and protagonist decision.
5. Grounding checks ordinary plan beats and all positive future-Drafter channels, including RAG transformation, reader reward, longform and literary rendering fields.
6. Grounding treats receipts as provenance only; it excludes `source_refs`, `query_or_need`, `do_not_use`, `forbidden_*`, `reveal_budget` and `knowledge_boundary` from occurrence claims.
7. Project-All rejects a saved passing receipt when its review protocol differs from the current Grounding protocol.

## Compatibility

- persisted plan schema is unchanged;
- no migration is required;
- old plans remain readable but cannot be finalized with a stale Grounding receipt;
- model-visible protocol digests change intentionally;
- changes are project-neutral and contain no 九九 or model-specific production rule.
