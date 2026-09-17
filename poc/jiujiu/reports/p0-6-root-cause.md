# P0-6 Proposal Isolation Root Cause

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Date: 2026-09-17

Result: **CONFIRMED — PRESENCE/PROVENANCE IS NOT DISTINCT FROM FACT AUTHORITY AT ALL INGESTION BOUNDARIES**

## Primary defect

The repository has strong integrity proofs after Character/World/Planning artifacts exist, but no generic first-class contract for an unapproved setting proposal. `RAGChunk.SourceKind`, summary files, context containers and `SourceBindingV2` identify where bytes came from; they do not say whether those bytes are Proposal, Reference, Simulation Fact or Accepted Canon.

## Concrete fail-open path

1. Persist P001 in a RAG chunk with `source_kind=proposal`.
2. `IsDesignOnlySourceKind` does not recognize `proposal`.
3. `activeRAGFactChunks`, BM25/vector selection and `persistChapterRAGFactReceipt` accept every non-forbidden, non-design kind.
4. The receipt calls the selection a fact and supplies an exact source token to Planner.
5. Strong downstream digest and bundle checks can prove the contaminated bytes are unchanged, but cannot prove a human approved P001.

Equivalent contamination can originate in model-authored summary/foundation fields because their save paths do not consult a pending-proposal registry.

## Why existing gates are insufficient

- P0-3 protects established Canon and derived coherence but cannot classify an unregistered proposal.
- P0-4/P0-5 Grounding and projection authority can semantically reject unsupported facts, but model review is not the only Authority gate.
- Character/World receipts prove that models consumed exact inputs and that actions were arbitrated; they cannot repair authority already lost upstream.
- Seal/promotion/CAS prove identity and atomicity. They do not authorize source semantics.

## Minimal generic fix

Introduce a typed pending-proposal registry and a shared Host validator; quarantine proposal RAG from fact recall; expose proposals only in labelled Architect context; reject pending proposal claims at foundation, summary, Character, World, Plan and bundle/recovery boundaries; add explicit Source Binding Authority. Do not implement model-writable approval. Preserve compatibility for historical artifacts without proposal metadata.
