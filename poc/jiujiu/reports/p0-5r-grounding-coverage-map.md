# P0-5R Grounding Coverage Map

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Inventory unit: the 65 direct JSON field surfaces declared by `ChapterPlan`,
`ChapterContract`, and `ChapterCausalSimulation`. Nested members are classified
with their owning surface from schema, consumer, validator, and bundle use. This
is a complete direct-surface inventory, not a claim that every nested fact has
an independent per-leaf source binding.

| Plan Surface | Fact-bearing? | Grounding Input | Host Validator | Source Binding | Coverage |
|---|---:|---|---|---|---|
| `/chapter` | NO | FULL_PLAN | exact chapter checks | NONE | **NOT_REQUIRED** |
| `/title` | YES | FULL_PLAN | shape only | NONE | **FULL** |
| `/goal` | YES | FULL_PLAN | semantic model review | NONE | **FULL** |
| `/conflict` | YES | FULL_PLAN | semantic model review | NONE | **FULL** |
| `/hook` | YES | FULL_PLAN | semantic model review | NONE | **FULL** |
| `/emotion_arc` | YES | FULL_PLAN | shape only | NONE | **FULL** |
| `/notes` | YES | FULL_PLAN | proposal/contamination checks | NONE | **FULL** |
| `/advance_events` | YES | FULL_PLAN | event/weave validation | NONE | **FULL** |
| `/grounding_review` | NO | STRIPPED_TO_AVOID_SELF_ATTESTATION | exact digest/protocol validation | NONE | **NOT_REQUIRED** |
| `/contract/required_beats` | YES | FULL_PLAN | contract/shape plus semantic Grounding | NONE | **FULL** |
| `/contract/forbidden_moves` | NO | FULL_PLAN | contract/shape plus semantic Grounding | NONE | **NOT_REQUIRED** |
| `/contract/continuity_checks` | YES | FULL_PLAN | contract/shape plus semantic Grounding | NONE | **FULL** |
| `/contract/evaluation_focus` | YES | FULL_PLAN | contract/shape plus semantic Grounding | NONE | **FULL** |
| `/contract/emotion_target` | NO | FULL_PLAN | contract/shape plus semantic Grounding | NONE | **NOT_REQUIRED** |
| `/contract/payoff_points` | YES | FULL_PLAN | contract/shape plus semantic Grounding | NONE | **FULL** |
| `/contract/hook_goal` | YES | FULL_PLAN | contract/shape plus semantic Grounding | NONE | **FULL** |
| `/contract/scene_anchors` | YES | FULL_PLAN | contract/shape plus semantic Grounding | NONE | **FULL** |
| `/causal_simulation/world_simulation_id` | NO | FULL_PLAN | exact equality | NONE | **NOT_REQUIRED** |
| `/causal_simulation/protagonist_decision` | YES | FULL_PLAN | exact equality | NONE | **FULL** |
| `/causal_simulation/project_promise` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/chapter_function` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/context_sources` | NO | FULL_PLAN | exact current tokens | NONE | **NOT_REQUIRED** |
| `/causal_simulation/writing_norms_applied` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/anti_ai_execution_plan` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/external_reference_plan` | YES | FULL_PLAN | current RAG/craft receipt validators | POST_GROUNDING_PER_ENTRY_TRANSFORMATION_AND_RECEIPT | **FULL** |
| `/causal_simulation/trend_language_plan` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/reader_entertainment_plan` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/grounding_details` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | POST_GROUNDING_RECEIPT_LEVEL_ONLY | **FULL** |
| `/causal_simulation/offscreen_character_stage` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/longform_opening` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/character_arc_tests` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/reader_reward_plan` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/reader_retention_plan` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/render_capacity` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/arc_transition_contract` | YES | FULL_PLAN | exact predecessor/outgoing validation | NONE | **FULL** |
| `/causal_simulation/evidence_return_chains` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/ending_consequence_contract` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/dormant_character_policy` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/reality_support_plan` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | POST_GROUNDING_RECEIPT_LEVEL_ONLY | **FULL** |
| `/causal_simulation/emotional_logic` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/relationship_emotion_arcs` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/visual_design` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/character_kit` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/world_background_layers` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/information_asymmetry` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/hidden_rule_pressure` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/social_mood_rumors` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/ritual_calendar` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/structural_resources` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/cosmology_checks` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/conflict_web` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/narrative_tension_matrix` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/initial_state` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/voice_logic` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/dialogue_scene_blueprints` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/literary_rendering_plan` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/crowd_roles` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/review_refinement` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/environment_state` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/world_rules_in_force` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/information_gaps` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/causal_beats` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/decision_points` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/outcome_shift` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |
| `/causal_simulation/scene_constraints` | YES | FULL_PLAN | shape/domain validators + semantic Grounding | NONE | **FULL** |

## Coverage conclusion

```text
GROUNDING FACT-SURFACE INVENTORY: COMPLETE
GROUNDING COVERAGE: PASS — NO INPUT COVERAGE DEFECT
```

The exact Grounding packet includes every Plan field; activation mode also includes the verified per-cycle decisions, POV before/after state and new POV knowledge. It does not carry every actor's original observation, every physical-state record, or a per-leaf authority binding. The prompt's general rule still requires every actual action, meeting, location, duration, known fact and result to be supported, so offscreen actions and knowledge are within its review scope even when not repeated in the later illustrative field list.

`SourceBindingV2` is generated after Grounding during bundle construction. External references receive per-entry transformation bindings; RAG/craft receive receipt-level bindings. `grounding_details` and `reality_support_plan` do not receive independent per-field bindings. These bindings provide bundle-time postconditions and do not authorize or expand the earlier Grounding packet, so coverage here does not depend on them or treat them as a second authority system.

The substantive fox false pass is classified separately as reviewer judgment failure: the field and source were both present, but the probabilistic verdict missed an unsupported causal-risk fragment.
