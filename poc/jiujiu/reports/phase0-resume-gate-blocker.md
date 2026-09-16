# Phase 0 Resume Gate Blocker

## New blocker

`P0-RESUME-001`: chapter-zero outline repair is valid at operation 0 but is overwritten by later outline-all operations before publication.

## Blocking field/state

- operation 0 receipt says all three chapter replacements were applied;
- `after_layered_digest = sha256:afc1ab134d27491c7444684e4138d49b0fd941ce8f34c03db47bae31d1e00a03`;
- operation 3 finishes with `after_layered_digest = sha256:e42cc429dcd10281a7c1f8a3b0de807aaad4327586bd6b795d96598de26b86fc`;
- final published `outline.json` does not contain the manifest replacements.

## Source

TEST_SCAFFOLD repair manifest `phase0-resume-outline-repair.json` supplied complete title/core_event/hook/scenes replacements for chapters 1–3. It explicitly removed messenger contact, credential verification as a selected dependency, next-chapter material promises, and in-story pipeline actors.

## Architect output

Later outline-all model operations regenerated all three chapter bodies. The final output contains:

- an anonymous but concrete messenger-like actor (`岸上有人`, `喊话人`);
- 九九追问来源 and the speaker repeating the claim;
- Chapter 2 choosing that pursuit as an actual event;
- planning/pipeline behavior represented inside Chapter 3 story content.

This output violates the authorized TEST_SCAFFOLD repair even though operation 0 succeeded.

## World Arbiter output

NOT EXECUTED for the corrected attempt. Running rehearsal would evaluate the overwritten outline, not the authorized corrected outline.

## Readiness path

The historical readiness code is unchanged: unresolved/infeasible contract checks and missing/unclear material checks force `ready_for_detail=false`. That path was not reached for the corrected attempt because no legal corrected outline survived publication.

## Classification

- TEST_SCAFFOLD input problem: corrected by the manifest.
- Formal Canon problem: NO; official input remained byte-identical.
- World Arbiter too strict: NOT APPLICABLE; it was not called.
- Product defect: YES. The documented host-only sparse chapter replacement is not preserved through the subsequent outline-all operation chain, and the publication verifier accepts the later drift as a new valid tail.
- Core-source modification required to proceed safely: YES, or an equivalent first-class post-expansion repair boundary. Not authorized in Phase 0.

## Material Check Blocking Semantics gap retained

The prior product gap is recorded but not repaired:

- A. Architect may promote an optional Character action into a material dependency.
- B. Material checks lack a typed `required_for_selected_path` / blocking semantic.
- C. Validator accepts `status=missing` together with a non-blocking narrative.
- D. Arbiter cannot remove or legally downgrade the misclassified check.

This gap was not re-tested because outline repair publication failed earlier.

## Stop condition

Triggered: continuing requires a core product change or manual authoritative-state mutation. Both are prohibited. STOP.
