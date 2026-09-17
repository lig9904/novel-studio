# P0-2 Code Change

Date: 2026-09-17

Core commit: `1fb49d33d58a2137679fbf4c8f9e280c2bfa532a`

Branch: `codex/character-readiness-freeze-order`

## Change

The readiness context now freezes an explicit Host-only policy binding before the activation session is created:

- exact Character Activation policy;
- exact executable producer digest;
- normalized producer policy inventory;
- selected Readiness policy;
- Soft Event contract policy;
- content digest of the complete binding.

The resolved producer is made explicit on the chapter boundary before context construction. The same producer is then used by activation input preparation.

Before the first Character model call, and again when constructing a Readiness review input, the Host verifies that the actual stimulus contains the exact frozen producer, activation policy, Soft Event marker and policy inventory. `Stimulus=new / Context=old`, producer substitution and policy downgrade fail closed.

The Host binding is excluded from the model-facing readiness view; the model receives the selected v2 schema and evidence aliases, not producer digests or authority metadata.

V3 readiness receipts are accepted by complete activation chapter evidence. Reused continuation proposal digests are disambiguated by the required same-cycle cycle/arbitration evidence before matching the proposal.

## Existing Domain terminology

No duplicate enum was added:

- task `ACCEPTED` → existing `OCCURRED`;
- plain rejection without a proven consequence → existing non-closing `DEFERRED`;
- `REJECTED_WITH_CONSEQUENCE` → unchanged;
- `SUPERSEDED_BY_ACTUAL_CHOICE` → unchanged;
- task `UNRESOLVED` → existing non-closing `DEFERRED`;
- `HARD_CONTRACT_UNSATISFIED` remains the separate hard fail-closed state.

P0-3 protected canon / derived coherence code was not modified.
