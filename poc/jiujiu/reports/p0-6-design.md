# P0-6 Proposal Isolation Design

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Date: 2026-09-17

## Contract

```text
AI proposal
    ↓ typed StoryProposalV1
PROPOSAL / PENDING
NOT ACCEPTED CANON
NOT HUMAN APPROVED
    ↓ discussion through dedicated capability only
no fact authority
    ↓ separate future Host/Human approval contract
Accepted Canon (not implemented by P0-6)
```

V1 deliberately contains no model-writable approved state. Proposal registration and approval are separate authorities.

## Enforcement

- `meta/story_proposals.json` is a content-addressed pending registry.
- Host derives the subject/predicate/object claim identity from the statement; model-supplied identity, status or approval fields are rejected.
- Proposal material is available through dedicated submit/list tools and is absent from shared `novel_context`.
- `source_kind=proposal` is design-only and cannot enter normal fact recall or `RAGFactReceipt`.
- Pending queues, old indexes and receipt replay are revalidated against the current registry.
- Foundation, summaries, world tick, Character observation/decision, World arbitration, simulation, Planner, commit/rewrite recovery and Project-All bundle/recovery use the shared Host guard.
- `SourceBindingV2.Authority` distinguishes `PROPOSAL`, `REFERENCE`, `SIMULATION_FACT`, `DERIVED_PRESENTATION` and `ACCEPTED_CANON`. An explicit proposal kind/ref cannot downgrade or self-label as Canon.
- Project-All source identity includes the registry, and workspace manifests bind both live and shadow registry bytes.

## Compatibility

- A missing registry means no registered proposal and preserves historical hashes because the new snapshot entry is conditional.
- Historical `SourceBindingV2` JSON remains byte-compatible through `authority,omitempty`; explicit proposal sources must use `PROPOSAL`.
- Existing Character/World/Planner prompts and producer protocol digests are unchanged.

## Text evidence boundary

The deterministic matcher supports the tested subject-predicate-object forms and structured attribute objects. It distinguishes negative/TBD fields, unrelated characters, fire objects and separate clauses. It does not claim universal natural-language semantic equivalence; existing simulation authority and Planner Grounding remain required backstops.
