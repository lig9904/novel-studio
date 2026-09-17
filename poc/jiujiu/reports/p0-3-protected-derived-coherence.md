# P0-3 Protected Canon / Derived Coherence

Date: 2026-09-17

Result: **PASS**

Classification: **TEST_SCAFFOLD / NOT ACCEPTED CANON / NOT HUMAN APPROVAL**

## Contract

`meta/world_coherence_report.json/md` is now treated as source-bound derived evidence, not as source canon.

The authored source/protected root remains responsible for characters, world rules, hard contracts, premise/canon and every other non-authorized source file. A new v3 outline-all receipt separately binds `derived_coherence_evidence_root` to:

- the current deterministic coherence `source_digest`;
- the coherence `report_digest`;
- the exact JSON artifact digest;
- the exact Markdown projection digest.

The Host verifies the report against the current world rules, world codex and book world, requires a ready deterministic audit, and verifies that the Markdown artifact is the exact projection of that report. Candidate completion, publication, published-stage verification and recoverable publication all check this derived root.

Historical v2 outline-all receipts remain readable and cannot claim the new v3 derived authority.

## Required cases

| Case | Result |
|---|---|
| A: authorized world-rule visibility migration invalidates the old report, refreshes it and binds the new source digest | PASS |
| B: unauthorized Character change | FAIL CLOSED |
| B: unauthorized World Rule change | FAIL CLOSED |
| B: unauthorized hard contract / Compass non-negotiable change | FAIL CLOSED |
| B: unauthorized premise/canon change | FAIL CLOSED |
| C: source changed while report retains the old digest | FAIL CLOSED |
| D: forged or detached Markdown/derived evidence | FAIL CLOSED |

## Regression verification

- `go test -p 1 -count=1 ./...`: PASS.
- targeted `go test -race` for outline-all, protected canon, derived coherence and receipt paths: PASS.
- `go vet ./...`: PASS.
- `go build ./...`: PASS.

Core commit: `9d296a832e404ee9fb855b4f178e71e3575fbf52`

Branch: `codex/physical-resource-semantics`

Remote: `lig9904/novel-studio`

No upstream PR was opened. P0-1 Knowledge Boundary, P0-2 Chapter Readiness, Proposal Isolation and Foreshadow/RISK-001 were not modified by this commit.

## Same-PoC native evidence

The native pipeline, without archive-copy or manual state editing, completed:

`rebase → outline-all → zero-init → preplan → rehearse-arc`

The published outline-all receipt is version 3 and binds:

- protected canon root: `sha256:c9b944f17b1e7f03a45e2627de2f38950c41f49002492d6c8b01cceb571405fb`;
- derived coherence evidence root: `sha256:84a35f66b51cfd0ce4db95c8d3cca311e24f57cb39b82fb21506db320cd11a14`;
- receipt digest: `sha256:bbf9a69d9cd15a7d657c07567bb6e922fd182d428a2b0b166e0e06ec16743a4a`.

The deterministic coherence report is ready and binds source digest `sha256:1ca7f00142307b8f0c8a6346ffd26d1dbaa803bbd90f0783ffe7153a78811155`.

The whole-arc rehearsal completed with `ready_for_detail=true`, report digest `sha256:228542551ede63f25cbf9c8ccc5ac1fcd17e9f299782cb3577d0ebd897b08ea0`, zero blocking contracts and zero blocking required materials.

Evidence: `poc/jiujiu/snapshots/phase0-p0-3-runtime/`.
