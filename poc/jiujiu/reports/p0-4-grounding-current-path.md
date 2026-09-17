# P0-4 Planner Grounding Current Path

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`


Date: 2026-09-17

Status: **READ-ONLY ROOT-CAUSE RECONNAISSANCE COMPLETE / NO STORY ENGINE CODE CHANGED**

Classification: **TEST_SCAFFOLD / P0-2 FROZEN STORY FACTS / CHAPTER 1 NOT PUBLISHED**

## MODEL ROUTING

```text
Task Level: LEVEL 2 — core pipeline protocol diagnosis
Primary: /root, GPT-5.6 Sol / High
Delegated: /root/explorer, GPT-5.6 Terra / Medium; /root/worker, GPT-5.6 Luna / Medium
Escalations: none
Reviewer: not used in the read-only phase
Reason: root owns the protocol conclusion; Terra traced the call path; Luna indexed frozen runtime evidence.
```

The delegated model and effort values above are the explicit runtime spawn metadata observed by the root, not values inferred only from `.codex/agents/*.toml`.

## Current execution path

```text
project-all reuses retained Chapter 1 simulation
  -> Planner builds/retains 01.plan.partial.json
  -> plan_details finalize
  -> deterministic plan checks
  -> reviewChapterPlanGrounding
       -> resolve world_arbiter model snapshot
       -> construct exact activation-trace grounding input
       -> one non-streaming model.Generate call
       -> require exactly one submit_plan_grounding_verdict tool call
       -> strict JSON + evidence-pointer validation
       -> persist immutable grounding audit only after a valid verdict
  -> consume planning-context receipt
  -> save formal chapter plan
  -> start/advance chapter checkpoint
  -> build Chapter 1 bundle
```

The failure occurs inside `reviewChapterPlanGrounding`, before the planning-context receipt is consumed, before the formal plan is saved, before the checkpoint advances and before a bundle can be built.

## Required ten-point reconnaissance

### 1. Planner Grounding Reviewer entry

- `internal/agents/project_all.go:853-929` attaches `NewPlanGroundingReviewer` to the Planner's `plan_details` tool.
- `internal/tools/plan_chapter.go:274-290` calls grounding after deterministic plan checks and before any formal-plan save.
- `internal/tools/plan_grounding.go:46-158` owns exact input construction, audit lookup, model review, receipt finalization and pass/fail handling.

### 2. Model-facing input

The input root is `policy`, `review_protocol`, `plan`, `simulation`, `pov_observation`, `arbitration` and optional `activation` (`internal/domain/plan_grounding.go:30-38`). P0-2 is activation mode, so the reviewer sends one indivisible `ExactAgentPacket` containing the complete JSON input (`internal/agents/plan_grounding.go:128-172`). The source facts are retained under `poc/jiujiu/snapshots/phase0-p0-2-freeze-order/`; the only Planner output is `chapter_001_plan_partial.json`.

### 3. Prompt

`internal/agents/plan_grounding.go:20-24` defines the base classifier prompt. Activation runs append the activation-trace rules at `:38-39`; work-artifact rules are appended only when that policy is present (`:41-42`). The prompt requires one tool submission, pass/fail classification only, exact evidence pointers and no story rewrite.

### 4. Tool definition

`internal/agents/plan_grounding.go:26-36` exposes one read-only capability named `submit_plan_grounding_verdict`.

### 5. JSON Schema

The root requires `pass:boolean` and `findings:array`. Each finding requires `kind`, `plan_path`, `plan_quote`, `source_path`, `source_quote` and `explanation`; `kind` is restricted to `time`, `location`, `knowledge`, `intent` or `outcome`.

### 6. Host rule for one structured verdict

`internal/agents/plan_grounding.go:173-203` performs one non-streaming `Generate` call with one tool spec and `max_tokens=6144`. The response must contain exactly one tool call, named `submit_plan_grounding_verdict`, with parseable arguments. The decoder rejects unknown fields, missing/non-null `pass` or `findings`, and trailing JSON. `internal/domain/plan_grounding.go:120-159` then enforces pass iff findings are empty, fail with 1-8 findings, allowed kinds, bounded text, valid plan/source pointer namespaces and quotes that occur at the cited input pointers.

### 7. Current DeepSeek return

P0-2 retained six `plan_grounding` model usage records for `deepseek/deepseek-flash`. All six ended with the Host error:

```text
plan grounding review did not complete: plan grounding must return exactly one structured verdict
```

The original assistant response blocks were not persisted. Existing evidence therefore cannot distinguish zero tool calls, multiple tool calls, a wrong tool name or `ArgsInvalid`; that detail is `UNKNOWN` until the controlled replay captures the raw response shape.

### 8. Non-submission behavior

- Host-accepted structured verdicts: `0/6`.
- Host result: rejected before receipt creation.
- Formal grounding audit/receipt: absent.
- Partial plan: retained.
- Formal Chapter 1 plan/bundle: absent.

### 9. Retry count and mechanism

There is no dedicated retry loop inside the Grounding Reviewer; each review is one `model.Generate` call. Planner repair handling retried `plan_details finalize` three times per invocation and then cancelled the Planner session while retaining the partial. P0-2 ran two such invocations, producing six Grounding calls in total.

The six recorded input/output token pairs were:

| Call | Input | Output |
|---:|---:|---:|
| 1 | 24,625 | 6,144 |
| 2 | 25,189 | 6,144 |
| 3 | 25,753 | 6,144 |
| 4 | 26,478 | 6,145 |
| 5 | 26,656 | 6,144 |
| 6 | 26,546 | 6,144 |

Aggregate: input `155,247`, output `36,865`, cache read `17,792`, total `192,112`. Recorded cost is `0 USD`, which means cost was unavailable/unpriced in this audit and does not prove the provider charged zero.

### 10. Whether retry used identical input

The historical six calls cannot be proven byte-identical because P0-2 did not persist the canonical grounding input digest or raw response for each attempt. Their different token counts show that the complete provider requests were not token-identical. The frozen P0-4 fixture will remove this ambiguity by persisting one canonical input and using it unchanged for every model-comparison call.

## Why Plan Finalization is blocked

Without a valid unique tool call, the Host cannot construct `PlanGroundingVerdict`; without a valid verdict it cannot finalize or save `PlanGroundingReceipt`; without a passing exact-plan receipt the formal plan fails closed. Therefore the planning-context receipt remains unconsumed, the formal plan is not saved, the chapter checkpoint does not advance and `buildPipelineProjectedChapterBundle` never receives an eligible plan.

## Historical runtime evidence

| Metric | Evidence | Result |
|---|---|---|
| Frozen generation | `pg2_b2354a98e8714bba2baca89a` | confirmed |
| Story Simulation / Character / World / Readiness | retained snapshots; readiness phase `ready` | completed before Planner failure |
| Grounding model | `deepseek/deepseek-flash` via role `world_arbiter` | confirmed |
| P0-2 configured reasoning | `high` | configuration evidence; provider-returned reasoning metadata unavailable |
| Grounding calls | 6 | confirmed |
| Accepted verdicts | 0 | confirmed |
| Whole failed invocation durations | 818.523 s and 178.199 s | confirmed |
| Per-call duration | unavailable historically | UNKNOWN |
| Raw response/tool-call shape | not persisted historically | UNKNOWN |
| Bundle published | no | confirmed |

Primary evidence locations:

- `poc/jiujiu/reports/p0-2-summary.md:3-35`
- `poc/jiujiu/project/test-scaffold-run/output/novel/meta/runtime/usage_audit.jsonl:630-655`
- `poc/jiujiu/project/test-scaffold-run/output/novel/meta/pipeline_timings.jsonl:45-46`
- `~/.novel-studio/last-error.log:59-62`
- `internal/agents/plan_grounding.go:20-203`
- `internal/domain/plan_grounding.go:30-205`
- `internal/tools/plan_grounding.go:43-158`
- `internal/tools/plan_chapter.go:266-306`

## Read-only conclusion before controlled replay

The retained evidence proves repeated Host rejection at the unique structured-verdict boundary. It does not yet separate model/tool compliance from protocol or transport behavior because the raw DeepSeek response and a same-input stronger-model control are missing. No code-fix authorization is implied by this reconnaissance.
