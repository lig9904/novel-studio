# P0-5R Planner Routing Decision

> **Authority:** `TEST_SCAFFOLD` · `NOT ACCEPTED CANON` · `NOT HUMAN APPROVAL`

Current Project-All behavior is:

```text
Planner model = models.ForRole("writer")
Planner reasoning/max_turns = writer settings
project_all_planner = accounting/stage name only
```

The isolated run changed only `roles.writer` to Sol/Medium. This did not permanently change the project and did not change Drafter configuration.

There is no independent `planner` role today. P0-5R does **not** add one because the required evidence condition was not met:

```text
DeepSeek historical failure
+
Sol Medium stable positive PASS
```

The second term is absent. Permanent routing remains unchanged.
