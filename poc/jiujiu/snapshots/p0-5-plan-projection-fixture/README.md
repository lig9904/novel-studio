# P0-5 Frozen Simulation → Plan Projection Fixture

Status: **IMMUTABLE FOR P0-5 DIAGNOSIS AND CONTROLLED REPLAY**

Generation: `pg2_b2354a98e8714bba2baca89a`

Chapter / simulation: `1` / `ch001-4db45dbca11c`

This fixture binds the retained P0-2 Character decisions, World Arbitration, Readiness, Story Facts and World State to the rejected Planner plan and both Host-valid P0-4 Grounding verdicts.

The exact historical `<host_prefetched_novel_context>` byte packet was not persisted by P0-2. `manifest.json` therefore records a durable component binding instead of claiming byte-identical reconstruction. The bound components include the Soft Outline, final simulation/evidence, projected state, RAG/craft receipts, planning access receipt and Planner prompt used to construct that packet.

The fixture must not be edited during Planner/Grounding comparison. Runtime outputs belong outside this directory.
