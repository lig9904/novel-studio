# P0-6 Proposal Isolation Authority Audit

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Date: 2026-09-17

Status: **READ-ONLY AUDIT COMPLETE — CONTAMINATION PATH CONFIRMED**

## Controlled proposal

```text
ID: PROPOSAL-P001
Statement: 凤凰具有火属性。
Authority: PROPOSAL
Status: PENDING
Accepted Canon: false
Human approved: false
```

The official source says phoenix fire/rebirth/healing/immortality powers are TBD and must not be completed automatically. P001 is not a new 九九 fact.

## Authority chain

| Stage | Current behavior | Authority finding |
|---|---|---|
| Proposal creation | No generic typed Story Proposal/approval object | Proposal status is not mechanically preserved across stores |
| Persistence | Arbitrary text can be placed in RAG chunks, summaries, planning inputs or model-produced foundation fields | Path/presence can substitute for authority |
| Context assembly | Architect context loads summaries/foundation; chapter context recalls all non-forbidden, non-design RAG kinds | No generic pending-proposal quarantine |
| Retrieval/RAG | `activeRAGFactChunks` rejects forbidden/design-only kinds; unknown kinds remain factual | `source_kind=proposal` is currently fail-open |
| Fact receipt | Receipt binds chunk identity, hash, source path and kind | Provenance is mistaken for fact authority |
| Summary/memory | Arc/volume summary tools persist model-authored claims without proposal checks | Proposal text can become durable planning memory |
| Character observation/decision | Strong observation/proposal/arbitration digest binding exists | It proves source integrity, not that an upstream claim was approved Canon |
| World Arbiter/World State | Final receipts bind actual decisions and mutations | Contaminated inputs can still be faithfully propagated |
| Planner/Grounding | P0-5 forbids RAG as occurrence authority and Grounding checks positive fields | Semantic model gate remains a backstop; Host lacks a generic registered-proposal gate |
| Recovery | Digests/CAS authenticate retained artifacts | Recovery can faithfully replay an already contaminated artifact |
| Bundle | `SourceBindingV2` has kind/provenance/usable facts but no explicit Authority | Presence and receipt identity can still look fact-bearing |
| Accepted Canon | Seal/promote boundaries are strong | They validate the chain, not the missing Proposal→Approval transition |

## Existing protections retained

Character proposal/arbitration receipts, World simulation authority receipts, P0-5 projection authority, Grounding, bundle digests, CAS publication and accepted-outcome promotion remain necessary. None is weakened by P0-6.

## Required contract

- Proposal content must have a typed, durable `PROPOSAL / PENDING / NOT ACCEPTED CANON / NOT HUMAN APPROVED` identity.
- Proposal material may be visible only in a labelled proposal context.
- Proposal RAG chunks must be excluded from ordinary fact recall and fact receipts.
- Fact-bearing Host writes must reject a registered pending proposal claim.
- Source bindings must state whether material is `PROPOSAL`, `REFERENCE`, `SIMULATION_FACT`, `DERIVED_PRESENTATION` or `ACCEPTED_CANON`.
- No model/tool may create approval authority. Approval must be a separate host/human-controlled path.
