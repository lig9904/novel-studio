# Recovery test

Result: **PASS** for checkpoint/recovery integrity; no production chapter recovery was claimed.

Project-specific evidence:

- outline-all operation receipts persisted across process failures and resumed from the pending operation;
- after an Architect rehearsal succeeded and the Arbiter MCP inventory timed out, the next run reused the stored Architect draft and dispatched only World Arbiter;
- an interrupted DeepSeek Pro attempt left recoverable state and the later Flash run created a new model-bound rehearsal without corrupting prior reports;
- two `all_chapter_rebase` operations archived the failed generation before creating a new generation;
- the current rebase receipt has identical `source_root` and `archive_root`;
- final accepted/current chapter is 0, no prose exists, and no pending usage call remains;
- Store/Race recovery suites also passed in the baseline test phase.

Boundary: this proves chapter-zero candidate/checkpoint/rebase recovery, not recovery of an accepted rendered chapter.
