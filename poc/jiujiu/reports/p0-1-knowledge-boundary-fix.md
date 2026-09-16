# P0-1 Knowledge Boundary Fix

Result: **PASS**

Scope: Character View visibility/enforcement semantics and necessary tests only. Chapter Readiness, Proposal Isolation, RISK-001, seal, promote, render, and Phase 1 were not modified or executed.

## Contract

World rules now carry two independent dimensions:

- `enforcement_scope`: `GLOBAL` or `CHARACTER_SCOPED`, with optional `enforcement_character_ids`;
- `visibility_scope`: `PUBLIC`, `CHARACTER_SCOPED`, or `AUTHOR_ONLY`, with `character_ids` for scoped visibility.

`character_ids` and `enforcement_character_ids` accept a formal character name, alias, or registered AgentID. World Arbiter still receives author rules and their enforcement scope. Character observations receive only PUBLIC rules plus CHARACTER_SCOPED rules matching the exact actor. AUTHOR_ONLY rules never enter a character packet.

PUBLIC views that name a registered character now fail closed before character observation construction and are rejected by `save_foundation`; identity-specific facts must use CHARACTER_SCOPED visibility.

## Same-PoC validation

The TEST_SCAFFOLD rules were migrated without changing official `jiujiu-poc-input.md`:

- shape policy: globally enforced, scoped visibility to the three named actors with generic self-only text;
- autonomy pressure: enforced/visible only for 九九;
- fox tide-mark rule: enforced for 九尾狐, AUTHOR_ONLY visibility;
- source identity/name dossier: globally enforced, AUTHOR_ONLY visibility;
- generic evidence/knowledge/coastal rules remain PUBLIC.

The real host builder generated `phoenix-observation.json` from this same foundation. The packet contains:

- `自己是凤凰`;
- `自己与羽族有关`;
- the observe-only `近岸公开水况` resource and public channel;
- generic source/time/unknown-handling rules.

It contains none of: 九九, 螭吻, 九尾狐, 泼水节, 独角龙, private tide mark, fox secret mechanism, future plot, or other proposal material.

Evidence: `poc/jiujiu/snapshots/phase0-p0-1-character-view/`.

Next: STOP and wait for separate authorization before Chapter Readiness work.
