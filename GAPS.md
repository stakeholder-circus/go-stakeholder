> [!NOTE]
> Missing or deferred behavior must fail fast and be tracked explicitly. No placeholder behavior should mask absent parity work.

# Gaps

## Landed
- Typed config and value enums are implemented.
- Seeded scheduler and registry-based renderer dispatch are implemented.
- Normalized JSON output is implemented.
- The full 2026+ family registry is exposed through `--list-values`.
- Dedicated smoke evidence exists for the classic-six widening tranche:
  - `code_analyzer`
  - `data_processing`
  - `jargon`
  - `metrics`
  - `network_activity`
  - `system_monitoring`
- `agent_workflows` remains dedicated.

## Remaining blocker
- Experimental provider flags are parsed consistently, but the experimental provider runtime is still intentionally not implemented in this follower foundation.
- The CLI fails fast with an explicit not-implemented message when experimental provider inputs are supplied.

## Decision rules
- Deterministic mode stays the default.
- Grouped fallback renderers stay in place for non-smoke families.
- No repo-tracked secrets.
