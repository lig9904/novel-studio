# P0-2 Character Readiness Freeze-order Root Cause

Date: 2026-09-17

Result: **CONFIRMED — WIRING / FREEZE-ORDER DEFECT**

Classification: **TEST_SCAFFOLD / NOT ACCEPTED CANON / NOT HUMAN APPROVAL**

## Confirmed runtime evidence

Generation `pg2_f2fc66b794ea7dd681ea0ec9` persisted a first-cycle stimulus containing:

- `character-soft-event-readiness:v1`;
- producer `character-agent-protocol:sha256:3a49b7bd92e39427a395234e6f11c7a98139b9622591c4b25a502d12f5ca596e`.

The same chapter persisted a readiness context/audit using:

- context version `chapter-readiness:arbitrated-events.v1`;
- input policy `chapter-readiness:arbitrated-events.v1`;
- model view `readiness-model-view.requirements-once-aliases.v1`;
- receipt `character-chapter-readiness.v2`;
- `soft_event=null`.

This proves that producer assembly succeeded for the activation input while readiness froze an earlier, incomplete opening stimulus.

## Current order

```text
runCharacterActivationChapter
  |
  +-- resolve activation policy + producer digest
  |
  +-- buildWorldStimulusDraft(opening)
  |     sources do not yet contain the selected producer policy inventory
  |
  +-- buildCharacterReadinessContext(opening)
  |     HasCharacterSoftEventReadinessPolicyV1 == false
  |     context.version defaults to chapter-readiness:arbitrated-events.v1
  |
  +-- SaveCharacterReadinessContext       <-- FREEZE TOO EARLY
  +-- NewCharacterActivationSession binds chapter_context_digest
  |
  +-- loadOrPrepareCharacterActivationInputs
        +-- add character-agent-protocol:<producer>
        +-- prepareCharacterActivationChronology
              +-- characterActivationV3PoliciesForProducer(producer)
              +-- append character-soft-event-readiness:v1
        +-- finalize stimulus / observations

Character Agent -> World Arbiter -> Readiness
                                      |
                                      +-- reads frozen v1 context
                                      +-- selects v1 model view/schema
                                      +-- emits v2 receipt, soft_event=null
```

## Expected order

```text
resolve Character Activation Policy
  |
resolve exact frozen Producer
  |
assemble the complete producer policy inventory
  |
derive Soft Event Readiness contract from that exact inventory
  |
construct final opening semantics for Readiness
  |
build + finalize CharacterReadinessContext
  |
SaveCharacterReadinessContext          <-- FREEZE HERE
  |
create session bound to context digest
  |
build cycle stimulus using the same producer/policy inventory
  |
Character Agent -> World Arbiter -> matching Readiness model view/evaluation
```

## Code-path answers

1. **Producer Policy generation** — `characterActivationProtocolForPolicy`, `characterActivationProtocolV3SoftEventReadinessDigest`, and `characterActivationV3SoftEventReadinessPolicies` select the current executable producer and its complete policy inventory.
2. **Character Activation Policy generation** — `characterActivationPolicyForBoundary` selects v1/v2/v3 from the frozen projected boundary.
3. **Soft Event Contract attachment** — `characterActivationV3SoftEventReadinessPolicies` appends `CharacterSoftEventReadinessPolicyV1`; today it reaches the actual stimulus through `prepareCharacterActivationChronology`.
4. **Readiness Context construction** — `buildCharacterReadinessContext` derives v2 only when its input stimulus already contains `CharacterSoftEventReadinessPolicyV1`.
5. **Freeze point** — `SaveCharacterReadinessContext` runs before the first `loadOrPrepareCharacterActivationInputs`; the session immediately binds the resulting context digest.
6. **Receipt generation** — `FinalizeCharacterReadinessReview` emits v3 only when the frozen input policy is `CharacterReadinessReviewPolicyV2`; otherwise it emits the historical v2 receipt.
7. **Model View generation** — `NewCharacterReadinessModelCodecV1` selects v2 view/schema only from `input.Policy`, which is derived from the frozen context.
8. **Evaluation version source** — `runCharacterChapterReadiness` and `runVerifiedCharacterChapterReadiness` treat the stored context version as authoritative when hashing the review protocol and constructing model input.
9. **Why `soft_event=null`** — the stored context was v1, so the grouped v1 schema did not request `soft_event`, and the canonical v1 finalizer correctly rejected/omitted the new classification.
10. **Why stimulus was new while context was old** — the producer policy inventory was appended later inside `prepareCharacterActivationChronology`, after the context had already been finalized, saved and bound into the activation session.

## Required fix boundary

The fix must assemble one exact producer-policy inventory before context construction and reuse that inventory for activation input preparation. Context/source mismatch must fail closed. No character-, scene-, model- or PoC-specific condition is needed.

P0-3 protected canon / derived coherence behavior is unrelated to this defect and must remain unchanged.
