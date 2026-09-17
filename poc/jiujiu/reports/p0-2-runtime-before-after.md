# P0-2 Runtime Before / After

Date: 2026-09-17

Classification: **TEST_SCAFFOLD / NOT ACCEPTED CANON / NOT HUMAN APPROVAL**

| Binding/result | Before: `pg2_f2fc66b794ea7dd681ea0ec9` | After: `pg2_b2354a98e8714bba2baca89a` |
|---|---|---|
| Stimulus Soft Event marker | present | present |
| Readiness context | `arbitrated-events.v1` | `soft-event-outcome.v2` |
| Frozen producer binding | absent | present, digest `c2888d…c0e4` |
| Model view | requirements aliases v1 | soft-event evidence aliases v2 |
| Receipt version | `character-chapter-readiness.v2` | `character-chapter-readiness.v3` |
| `soft_event` | `null` | cycle 1 `DEFERRED`; cycle 2 `REJECTED_WITH_CONSEQUENCE` |
| Cycles before readiness | old run continued toward four-cycle ceiling | 2 cycles, then `ready_for_plan` |
| Character forced to accept | no, but old contract could not classify refusal | no; original refusal/reason preserved |

The after run used the same official input, Canon, Soft Event, DeepSeek Flash role routing and P0-3-published project. It was a single successor-generation run, not repeated sampling for a preferred choice.

Evidence: `poc/jiujiu/snapshots/phase0-p0-3-runtime/` (before) and `poc/jiujiu/snapshots/phase0-p0-2-freeze-order/` (after).
