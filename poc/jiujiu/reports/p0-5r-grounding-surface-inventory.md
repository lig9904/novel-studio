# P0-5R Grounding Surface Inventory

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Scope: current `ChapterPlan`, `ChapterContract`, and `ChapterCausalSimulation` JSON surfaces.

The inventory unit is each of the 65 direct fields declared by those three
structs. Nested member shapes and consumers are assessed under the owning
field. `COMPLETE` therefore means no direct schema surface was omitted; it does
not mean every nested factual sentence has an independent per-leaf authority
record.

| Plan path | Classification | Primary consumer |
|---|---|---|
| `/chapter` | `STRUCTURAL_ONLY` | Store/checkpoint/bundle identity |
| `/title` | `MIXED` | Drafter/render metadata |
| `/goal` | `FACT_BEARING` | Drafter/Editor |
| `/conflict` | `FACT_BEARING` | Drafter/Editor |
| `/hook` | `FACT_BEARING` | Drafter/Editor/continuity |
| `/emotion_arc` | `MIXED` | Drafter/Editor |
| `/notes` | `MIXED` | Planner repair context |
| `/advance_events` | `FACT_BEARING` | event weave |
| `/grounding_review` | `STRUCTURAL_ONLY` | receipt/read gate |
| `/contract/required_beats` | `FACT_BEARING` | Drafter/Editor hard results |
| `/contract/forbidden_moves` | `STRUCTURAL_ONLY` | negative constraint |
| `/contract/continuity_checks` | `MIXED` | Editor continuity |
| `/contract/evaluation_focus` | `MIXED` | Editor instruction that may carry positive required outcomes |
| `/contract/emotion_target` | `PRESENTATION_ONLY` | tone target |
| `/contract/payoff_points` | `FACT_BEARING` | Drafter payoff candidate |
| `/contract/hook_goal` | `MIXED` | Drafter ending intent |
| `/contract/scene_anchors` | `FACT_BEARING` | Drafter objects/actions |
| `/causal_simulation/world_simulation_id` | `STRUCTURAL_ONLY` | exact simulation binding |
| `/causal_simulation/protagonist_decision` | `FACT_BEARING` | POV choice |
| `/causal_simulation/project_promise` | `MIXED` | Drafter/longform |
| `/causal_simulation/chapter_function` | `MIXED` | Drafter/arc |
| `/causal_simulation/context_sources` | `STRUCTURAL_ONLY` | receipt/source identity |
| `/causal_simulation/writing_norms_applied` | `MIXED` | Drafter method plus positive chapter_application/proof_targets |
| `/causal_simulation/anti_ai_execution_plan` | `MIXED` | Drafter style guard plus positive counter_moves/dialogue actions |
| `/causal_simulation/external_reference_plan` | `MIXED` | Drafter fact anchors |
| `/causal_simulation/trend_language_plan` | `MIXED` | Drafter dialogue/texture |
| `/causal_simulation/reader_entertainment_plan` | `MIXED` | Drafter positive beats |
| `/causal_simulation/grounding_details` | `FACT_BEARING` | Drafter reality anchors |
| `/causal_simulation/offscreen_character_stage` | `FACT_BEARING` | Drafter + persisted character stage ledger |
| `/causal_simulation/longform_opening` | `MIXED` | Drafter opening/promises |
| `/causal_simulation/character_arc_tests` | `MIXED` | Drafter character actions |
| `/causal_simulation/reader_reward_plan` | `MIXED` | Drafter payoff/cost |
| `/causal_simulation/reader_retention_plan` | `MIXED` | Drafter surface/latent selection |
| `/causal_simulation/render_capacity` | `MIXED` | Drafter scene actions/consequences |
| `/causal_simulation/arc_transition_contract` | `FACT_BEARING` | bundle/next chapter |
| `/causal_simulation/evidence_return_chains` | `FACT_BEARING` | Drafter/future evidence |
| `/causal_simulation/ending_consequence_contract` | `FACT_BEARING` | bundle/next chapter |
| `/causal_simulation/dormant_character_policy` | `FACT_BEARING` | future character continuity |
| `/causal_simulation/reality_support_plan` | `FACT_BEARING` | Drafter reality facts |
| `/causal_simulation/emotional_logic` | `FACT_BEARING` | Drafter action/emotion evidence |
| `/causal_simulation/relationship_emotion_arcs` | `FACT_BEARING` | relationship continuity |
| `/causal_simulation/visual_design` | `FACT_BEARING` | Drafter character state |
| `/causal_simulation/character_kit` | `FACT_BEARING` | Drafter resources/abilities |
| `/causal_simulation/world_background_layers` | `FACT_BEARING` | Drafter world pressures |
| `/causal_simulation/information_asymmetry` | `FACT_BEARING` | knowledge boundaries |
| `/causal_simulation/hidden_rule_pressure` | `FACT_BEARING` | world rules/costs |
| `/causal_simulation/social_mood_rumors` | `FACT_BEARING` | crowd knowledge/actions |
| `/causal_simulation/ritual_calendar` | `FACT_BEARING` | time/deadline constraints |
| `/causal_simulation/structural_resources` | `FACT_BEARING` | resources/access/cost |
| `/causal_simulation/cosmology_checks` | `FACT_BEARING` | world rules/abilities |
| `/causal_simulation/conflict_web` | `FACT_BEARING` | actors/goals/resources |
| `/causal_simulation/narrative_tension_matrix` | `MIXED` | Drafter pressure/POV |
| `/causal_simulation/initial_state` | `FACT_BEARING` | Drafter opening character state |
| `/causal_simulation/voice_logic` | `MIXED` | Drafter knowledge/dialogue |
| `/causal_simulation/dialogue_scene_blueprints` | `FACT_BEARING` | Drafter communication/new knowledge |
| `/causal_simulation/literary_rendering_plan` | `MIXED` | Drafter presentation + state_change |
| `/causal_simulation/crowd_roles` | `FACT_BEARING` | Drafter participants/actions |
| `/causal_simulation/review_refinement` | `MIXED` | repair constraints/targets |
| `/causal_simulation/environment_state` | `FACT_BEARING` | Drafter visible world state |
| `/causal_simulation/world_rules_in_force` | `FACT_BEARING` | world constraint |
| `/causal_simulation/information_gaps` | `MIXED` | knowledge/uncertainty |
| `/causal_simulation/causal_beats` | `FACT_BEARING` | Drafter/bundle causality |
| `/causal_simulation/decision_points` | `FACT_BEARING` | Drafter choices |
| `/causal_simulation/outcome_shift` | `FACT_BEARING` | bundle/state delta |
| `/causal_simulation/scene_constraints` | `MIXED` | Drafter authority boundary |

```text
SURFACES: 65
FACT/MIXED COVERAGE FULL: 59
NOT REQUIRED: 6
PARTIAL: 0
NONE: 0
```

Classification is based on schema, consumers and validators. `MIXED` means the object contains both authority-bearing and presentation/constraint members. `NOT_REQUIRED` is limited to identity, receipt, negative instruction or non-causal presentation fields that cannot independently create Story Truth.

The current activation Grounding packet carries the complete Plan with `grounding_review` removed, plus the verified per-cycle decision trace and POV before/after/new-knowledge state. It does not carry every actor's original observation, every physical-state record, or a per-leaf authority binding. `FULL` means the owning field is model-visible, the applicable authority packet is available as defined by the current protocol, and the broad action/location/knowledge/resource/event/causal rule applies. It does not mean each claim is supported, and it does not claim the probabilistic reviewer is infallible.
