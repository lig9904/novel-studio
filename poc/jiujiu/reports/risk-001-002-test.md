# RISK-001 and RISK-002

## RISK-001 — Foreshadow Rewrite

Result: **CONFIRMED in TEST_SCAFFOLD only**.

`foreshadow_ledger.initial.json` kept seed ID `ch01-opening-hook` but its description changed after authorized rebase, resource repair, outline-all, and zero-init rebuild. Before SHA-256: `7266bf472e5ecb14ecf4a40a146fa53be34eba2bf1ee66881c0a1ab28d8545c9`. After SHA-256: `51f2ff85f1449707261b2ea9f74fb86c75c55d042413edff62d50433da511027`.

Boundary: accepted chapter remained 0 and the ledger is TEST_SCAFFOLD. This confirms rewrite behavior during chapter-zero rebuild, not corruption of official Accepted Canon.

## RISK-002 — Candidate State / Accepted Canon

Result: **INCONCLUSIVE; official pollution NOT OBSERVED**.

- outline-all candidates remained in separate candidate directories and published only through directory transaction receipts;
- failed generations were archived through all-chapter rebase with equal source/archive content roots;
- source input SHA-256 remained unchanged;
- accepted/current chapter remained 0;
- no `project-all` proposal, arbitration, Character Agent canonical memory, promote receipt, rendered chapter, or accepted episode exists.

The requested proposal-candidate-to-accepted-canon pollution test cannot be completed because `project-all` never started. No pollution was observed, but absence of the candidate layer is not a PASS.
