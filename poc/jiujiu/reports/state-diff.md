# State diff

Comparison: `snapshot-before-simulation` → `snapshot-after-simulation`.

| Domain | Before | After blocker | Change |
|---|---|---|---|
| Official Canon input | SHA-256 `7ef95a...715a` | same SHA-256 | none |
| Accepted chapters | 0 | 0 | none |
| Character | three source-bounded scaffold states | same identities/known facts plus authorized physical resource balances | TEST_SCAFFOLD only |
| Knowledge | fox secret only in fox initial state | final rehearsal packet still exposes secret only to fox | no cross-character leak in input |
| Relationship | chapter 0, `contracts=null` | chapter 0, `contracts=null` | timestamp only |
| Timeline | chapter 0, first generation | chapter 0, new rebase generation | old generations archived |
| Foreshadow | one TEST_SCAFFOLD opening seed | same seed ID but rewritten description | RISK-001 confirmed in scaffold |
| RAG | initial scaffold index | rebuilt scaffold index after authorized rebase | test-only rebuild |
| Proposal candidates | outline candidates only | additional isolated outline candidates; no project-all proposal | no accepted proposal |
| Pending state | none in snapshot | none in final usage snapshot | recovered/cleared |

`characters.json` and `outline.json` changed because the user explicitly authorized TEST_SCAFFOLD resource/outline rebuilds. Those changes remain inside the isolated run and are not official 九九 Canon.

Proposal pollution check: `project-all` produced no character proposal, arbitration, or canonical memory file. Official input and accepted chapter state remain unchanged. This is positive non-mutation evidence, but Proposal Isolation remains INCONCLUSIVE because the required A/B branch test never ran.
