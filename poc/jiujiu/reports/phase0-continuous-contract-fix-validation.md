# Phase 0 — Continuous Contract Core Fix Validation

Status: **CORE FIX VERIFIED / REHEARSAL GATE STILL FAILS**

## Submitted core fix

- Branch: `codex/outline-contract-evidence-modes`
- Commit: `f6d611b` (`fix: distinguish continuous outline contracts`)
- Upstream pull request: `https://github.com/Xiaoyangy/novel-studio/pull/6`

The change implements the direction recorded in the upstream delivery audit: distinguish persistent constraints from terminal chapter payoffs without lowering the existing payoff validator.

### Contract behavior

- omitted/default `evidence_mode=payoff` keeps the original unique arc/chapter realization gate;
- the Host derives `continuous` only for source-bound non-negotiables that express prohibition, secrecy, unknown state, or a maintained boundary;
- continuous refs use `planned_payoff_chapter=0`, empty `planned_resolution`, remain arc-only, and retain the original author source digest as authority;
- ending and open-thread contracts remain payoff-only;
- model-selected mode drift and chapter copies of continuous refs fail closed;
- legacy refs with no `evidence_mode` remain payoff contracts.

## Verification

- `go test -p 1 -count=1 ./...`: PASS.
- targeted Race tests for continuous/payoff coverage: PASS.
- `go vet ./...`: PASS.
- `go build ./cmd/novel-studio`: PASS.
- real corrected-input 九九 outline-all: PASS and published.

## Real outline result

- outline-all status: `complete`;
- completed operations: 3;
- final layered digest: `sha256:0c95cd026388666bf184b9039efa6cbb00a0552eeee2ecbad8a1337875eb9876`;
- all 16 non-negotiables are arc-only `continuous` refs;
- chapters 1 and 2 contain no contract refs;
- chapter 3 contains only 5 payoff open-thread refs;
- no selected path requires locating/contacting a sender or obtaining a credential;
- zero-init: PASS;
- preplan: PASS.

## Rehearsal result

- report digest: `sha256:12888ad3c625ebf136547994bf822b73f882d4699882d165bbebb1c28a4269c4`;
- `ready_for_detail=false`;
- contract checks: 12 plausible, 4 conditional, 1 infeasible_prediction;
- material checks: 4 available, 1 not_required, 4 unclear;
- project-all and Character Agent: NOT RUN.

The original sender/credential material dependency is gone. The new blocking source is separate:

1. TEST_SCAFFOLD created the public-water resource as a counted `1 处` resource, while `operational_observation` requires a qualitative non-document resource (`actual_amount=null`, empty unit).
2. Private fox observation has only a secret mechanism, while the current operational-observation capability accepts public mechanisms only.
3. Phoenix has no public mechanism for non-owner qualitative water observation.
4. The resort notification effect has no addressable recipient actor.
5. Material checks still lack structured blocking semantics, so `unclear` blocks readiness even when the report says the operation is optional/non-blocking.

## Boundary

The contract-kind core fix is successful and submitted upstream. The later Rehearsal failure is not evidence that the fix failed; it is the next independent TEST_SCAFFOLD/capability/material-semantics boundary. No Promote, Render, Accepted Canon write, or Phase 1 feature work ran.

