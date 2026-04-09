# Example outputs

## Dedicated classic-six smoke examples

### `code_analyzer`
```text
[session.start] deterministic session started
[activity] code_analyzer smoke evidence: stakeholder kept the audited contract aligned for backend family=code_analyzer group=classic-six renderer=dedicated:code_analyzer trace=rust-stakeholder->stakeholder-core
[session.end] deterministic session finished
```

### `data_processing`
```text
[session.start] deterministic session started
[activity] data_processing smoke evidence: stakeholder kept the audited contract aligned for data-science family=data_processing group=classic-six renderer=dedicated:data_processing trace=rust-stakeholder->stakeholder-core
[session.end] deterministic session finished
```

### `jargon`
```text
[session.start] deterministic session started
[activity] jargon smoke evidence: stakeholder kept the audited contract aligned for backend family=jargon group=classic-six renderer=dedicated:jargon trace=rust-stakeholder->stakeholder-core
[session.end] deterministic session finished
```

### `metrics`
```text
[session.start] deterministic session started
[activity] metrics smoke evidence: stakeholder kept the audited contract aligned for dev-ops family=metrics group=classic-six renderer=dedicated:metrics trace=rust-stakeholder->stakeholder-core
[session.end] deterministic session finished
```

### `network_activity`
```text
[session.start] deterministic session started
[activity] network_activity smoke evidence: stakeholder kept the audited contract aligned for fullstack family=network_activity group=classic-six renderer=dedicated:network_activity trace=rust-stakeholder->stakeholder-core
[session.end] deterministic session finished
```

### `system_monitoring`
```text
[session.start] deterministic session started
[activity] system_monitoring smoke evidence: stakeholder kept the audited contract aligned for systems-programming family=system_monitoring group=classic-six renderer=dedicated:system_monitoring trace=rust-stakeholder->stakeholder-core
[session.end] deterministic session finished
```

## Modern-core dedicated example

### `agent_workflows`
```text
[session.start] deterministic session started
[activity] agent_workflows smoke evidence: stakeholder handoff loop completed under 17 seed family=agent_workflows group=modern-core renderer=dedicated:agent_workflows trace=rust-stakeholder->stakeholder-core
[session.end] deterministic session finished
```

## Grouped fallback examples

### `platform_engineering`
```text
[activity] modern-core fallback: platform_engineering handled by group:modern-core renderer for workflow and platform sweep family=platform_engineering group=modern-core renderer=group:modern-core trace=rust-stakeholder->stakeholder-core
```

### `supply_chain_security`
```text
[activity] modern-core fallback: supply_chain_security handled by group:modern-core renderer for workflow and platform sweep family=supply_chain_security group=modern-core renderer=group:modern-core trace=rust-stakeholder->stakeholder-core
```
