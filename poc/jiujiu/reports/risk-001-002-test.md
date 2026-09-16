# RISK-001 and RISK-002

## RISK-001 — Foreshadow Rewrite

Result: **CONFIRMED in TEST_SCAFFOLD only**.

`foreshadow_ledger.initial.json` kept seed ID `ch01-opening-hook` but its description changed after authorized rebase, resource repair, outline-all, and zero-init rebuild. Before SHA-256: `7266bf472e5ecb14ecf4a40a146fa53be34eba2bf1ee66881c0a1ab28d8545c9`. After SHA-256: `51f2ff85f1449707261b2ea9f74fb86c75c55d042413edff62d50433da511027`.

Boundary: accepted chapter remained 0 and the ledger is TEST_SCAFFOLD. This confirms rewrite behavior during chapter-zero rebuild, not corruption of official Accepted Canon.

## RISK-002 — Candidate State / Accepted Canon

Updated result: **PASS within the exercised non-promotion boundary**.

Project-all created generation `pg2_378bb2c691e5a07d2c2b8d02` with four real activation cycles, proposals, observations, arbitrations, candidate/projected memories, physical roots, readiness receipts, and usage rows. All remained under the isolated `.project-all` workspace. Live accepted/current chapter stayed 0; no sealed bundle, promote receipt, rendered chapter, or accepted episode was created; the official input SHA-256 remained unchanged.

This proves the engine did not automatically promote Candidate state or interpret AI results as human approval. An explicitly invoked Promote path was not authorized and is not claimed tested.
