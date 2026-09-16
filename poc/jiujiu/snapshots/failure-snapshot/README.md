# Failure snapshot

The blocker was detected before any story-generation call or Novel Studio run state was created. No rollback or reset was performed.

Evidence:

- `--pipeline --help` states that `--init-only` completes world/characters/full-book navigation and zero-init.
- `parsePipelineFlags` expands `--init-only` to `architect,outline-all,zero-init`.
- The pipeline help defines `outline-all` as the complete volume/arc/chapter contract created in a chapter-zero isolated workspace.
- Direct `--zero-init` acquisition fails closed unless a published `outline-all` exists.
- The CLI exposes no proposal-only, character-state-only, Knowledge Boundary, Hard Canon, Character Logic, or proposal-isolation command.

Running `--init-only` would therefore violate the Phase 0 prohibition on Season Planning. Bypassing this dependency would require a new core entrypoint or manual authoritative-state fabrication, both outside the authorized scope.
