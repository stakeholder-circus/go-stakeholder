> [!NOTE]
> Missing or deferred behavior must fail fast and be tracked explicitly. No placeholder behavior should mask absent parity work.

# Gaps

## Landed
- Typed config and value enums are implemented.
- Seeded scheduler and registry-based renderer dispatch are implemented.
- Normalized JSON output is implemented.
- The full 2026+ family registry is exposed through `--list-values`.
- Dedicated evidence exists for the classic-six tranche:
  - `code_analyzer`
  - `data_processing`
  - `jargon`
  - `metrics`
  - `network_activity`
  - `system_monitoring`
- Dedicated evidence exists for the modern-core tranche:
  - `agent_workflows`
  - `platform_engineering`
  - `observability_ai_runtime`
  - `delivery_preview_ops`
  - `supply_chain_security`

## Remaining blocker
- Experimental provider runtime is a current open gap in the future live-provider lane.
- The CLI fails fast with an explicit not-implemented message when experimental provider inputs are supplied.

## Decision rules
- Deterministic mode stays the default.
- Grouped fallback renderers stay in place for non-smoke families.
- No repo-tracked secrets.
