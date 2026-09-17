# P0-4 Grounding Fix Validation Fixture

This fixture preserves exactly the same Story Facts as `p0-4-grounding-fixture`. The canonical JSON excluding `review_protocol` has the same SHA-256:

```text
dfec76f217ea6de6eddd87a5c7a251cb01fabad9c075a019103db57fa1f88fcb
```

Only the model-facing pointer instructions, tool description and their derived protocol digest changed. The Host validator, Chapter Plan, simulation, activation evidence, Character decisions, World Arbitration and Readiness evidence are unchanged.

```text
base protocol before: sha256:39aa00184bebe288c0b4ce8e054141b186bc72275830a60dd397defe0e492452
base protocol after:  sha256:41daea71da9f8d7b80900cf633dc908bb1516ab43c57b50d5de7b73390cff4b5
input digest after:   sha256:4d5763e8c237e5f20f274fd478272902ca0e97194fd8109e17601c53b8ab1637
```
