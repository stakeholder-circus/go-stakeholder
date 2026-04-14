> [!IMPORTANT]
> This repository is part of a Codex-assisted rewrite experiment. All changes are manually reviewed, a human remains in the loop, and missing behavior is tracked explicitly rather than hidden. The project exists for fun, research, language learning, AI agent workflow/planning, interop experiments, and code review testing.
# go-stakeholder

Go is the current validated follower baseline through the modern-core wave for the stakeholder rewrite, with the live-provider lane still open for a later phase.

## Implemented surface
- Go modules with standard-library-first runtime code.
- Thin CLI entrypoint in `cmd/stakeholder`.
- Typed config and value enums for dev type, jargon, complexity, and output format.
- Seeded scheduler with current family-selection rules.
- Registry-based renderer dispatch with dedicated classic-six and modern-core renderers plus grouped fallback renderers for the remaining families.
- Normalized JSON session output.
- Full 2026+ family registry for `--list-values`.
- Dedicated evidence for `code_analyzer`, `data_processing`, `jargon`, `metrics`, `network_activity`, `system_monitoring`, `agent_workflows`, `platform_engineering`, `observability_ai_runtime`, `delivery_preview_ops`, and `supply_chain_security`.
- Experimental provider flags are parsed consistently and fail fast as the current open gap in the future live-provider lane.

## Commands
- `gofmt -l $(find . -name '*.go' -not -path './.git/*')`
- `golangci-lint run`
- `go build ./...`
- `go test ./...`
- `docker build -t go-stakeholder .`
- `docker run --rm go-stakeholder --list-values`

## Runtime notes
- Deterministic mode is the default.
- `--seed` and `--output-format` are part of the stable contract.
- `--list-values` emits the family and config registry as JSON.
- Native validation is the fast path on this workstation; Docker remains the portable release gate.
- Experimental provider flags are parsed, but runtime provider execution remains a current open gap in this lane.

## Gaps
- [GAPS.md](GAPS.md)
