# P0-4 Frozen Grounding Fixture

Status: **IMMUTABLE FOR MODEL COMPARISON**

This fixture is reconstructed from the retained P0-2 runtime generation `pg2_b2354a98e8714bba2baca89a` at core commit `1fb49d33d58a2137679fbf4c8f9e280c2bfa532a` and evidence commit `5850b04023812a200b0c4c0bbdee4962f14ec82d`.

The fixture contains the retained Planner partial, canonical `ChapterPlan`, Chapter 1 simulation, full activation evidence, exact model-facing `PlanGroundingInput`, system prompt, tool schema and serialized model messages. `fixture-metadata.json` binds the Story State, activation evidence, protocol and content hashes.

Canonical comparison input:

```text
input_digest: sha256:5a9a512d20a486d886c67e26e637b4ef56647b0cb5868b7df0da181f2be938c4
simulation_digest: sha256:cd312f3ac76daa340f5bec568c68f5286766775480b1432f06eea57e4fa80642
activation_evidence_digest: sha256:d498f00a4d40f9e635d54b3f5bdfdd968e170b56cff826cef247756c4992ecfa
generation_id: pg2_b2354a98e8714bba2baca89a
simulation_id: ch001-4db45dbca11c
activation_cycles: 2
```

Comparison results are stored outside this directory under `poc/jiujiu/reports/p0-4-grounding-comparison/` so the fixture is never mutated by a model run.
