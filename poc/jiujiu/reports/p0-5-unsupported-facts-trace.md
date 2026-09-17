# P0-5 Unsupported Facts Trace

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Status: **READ-ONLY TRACE COMPLETE**

## Origin boundary

The official PoC input does not contain a buoy rope, water gauge or on-site visitor. Their first durable appearance is the pre-simulation Soft Outline produced by `outline-all`:

- `source-soft-outline.json:5-9`;
- `source-layered-outline.json:15-19`.

They were then carried into the Character activation `soft_outline` as candidate pressure. This did not make them Story Facts. Cycle 1 Readiness explicitly records that the rope, gauge mismatch and visitor-delivered message “均未成为已记录的真实事件”. Cycle 2 closed the soft pressure through `REJECTED_WITH_CONSEQUENCE`, not by making the rejected outline scene occur.

The Planner later copied and expanded the stale candidates into the rejected formal-plan partial. P0-4 Grounding correctly rejected them against the final activation evidence.

## 浮标绳

| Question | Evidence-based answer |
|---|---|
| First appearance | Soft Outline `core_event`, `hook`, scene 0. Not in the official user input. |
| Upstream resource | `现场警示浮标` exists as `res_2222333344445555`, amount 1. No rope resource/entity exists in final `physical_state.resources`. A buoy does not authorize a detachable rope, its position, condition or use. |
| Planner promotion | Rejected plan makes the rope loose/displaced, uses it as an observed fact, scene anchor, causal trigger and ending consequence. |
| Causal persistence | Yes. It can be inspected, compared, moved, repaired and used as evidence; its unresolved handling creates a next-chapter obligation. It is a causal entity/resource, not presentation texture. |
| Grounding result | Rejected as unsupported outcome in both fixed Sol verdicts. |

## 水尺

| Question | Evidence-based answer |
|---|---|
| First appearance | Same Soft Outline `core_event` and scene 0. Not in the official user input. |
| World State device | No gauge/device/resource exists. Final World State contains only public coastal conditions, one warning buoy, six blank papers and the fox-private tide mark. |
| Qualitative → measurement escalation | Yes. Final evidence supports a bounded qualitative visual observation with no clearly located risk. Planner introduces an instrument, a previous reading and a comparison result. |
| Knowledge/Story Fact change | Yes. It makes 九九 know a measurable discrepancy and creates an unexplained fact requiring later investigation. |
| Grounding result | Rejected together with the rope as an unsupported actual result. |

## 现场来客

| Question | Evidence-based answer |
|---|---|
| First appearance | Soft Outline `core_event` and scene 1. Not in the official user input. |
| Registered character/entity | No corresponding actor/entity exists in final physical state, character registry projection or activation decisions. |
| Dialogue/event creation | Rejected plan adds arrival, gestures, spoken claims, questions and answers. |
| Knowledge change | Yes. The actual decision reason says the prestige lead “没有传递到我这里”. Planner turns an abstract unsourced pressure into delivered dialogue known to 九九. |
| Causal persistence | Yes. The encounter becomes the cause of refusal, dialogue conflict and future investigation pressure. |
| Grounding result | Rejected as unsupported knowledge/event creation in both fixed Sol verdicts. |

## Causal-persistence result

Deleting any of the three additions changes the rejected plan's later event sequence, character knowledge, evidence chain or next-chapter obligation. All three therefore require Story Authority and cannot be classified as `DERIVED_PRESENTATION`.

## Frozen evidence

The full trace is bound by `poc/jiujiu/snapshots/p0-5-plan-projection-fixture/manifest.json` and `SHA256SUMS`. The exact historical Planner packet was not persisted; the fixture explicitly binds its durable components without claiming a byte-identical reconstruction.
