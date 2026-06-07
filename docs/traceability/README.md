# Traceability

Map every future behavior change to audited Rust lines first.

## Current follower foundation
- Dedicated renderers now exist for the classic-six tranche:
  - `code_analyzer`
  - `data_processing`
  - `jargon`
  - `metrics`
  - `network_activity`
  - `system_monitoring`
- Dedicated renderers now exist for the modern-core tranche:
  - `agent_workflows`
  - `platform_engineering`
  - `observability_ai_runtime`
  - `delivery_preview_ops`
  - `supply_chain_security`
- Dedicated modern-core rows now point back to Rust generator files, Java renderer anchors, and the shared `stakeholder-core` contract docs.
- All other families stay on grouped fallback renderers until their tranche lands.
