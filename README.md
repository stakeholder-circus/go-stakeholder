> [!IMPORTANT]
> This repository is part of a Codex-assisted rewrite experiment. All changes are manually reviewed, a human remains in the loop, and missing behavior is tracked explicitly rather than hidden. The project exists for fun, research, language learning, AI agent workflow/planning, interop experiments, and code review testing.
# go-stakeholder

Go follower implementation through the modern-core wave for the stakeholder rewrite.

## Implemented surface
- Go modules with standard-library-first runtime code.
- Thin CLI entrypoint in `cmd/stakeholder`.
- Typed config and value enums for dev type, jargon, complexity, and output format.
- Seeded scheduler with current family-selection rules.
- Registry-based renderer dispatch with dedicated classic-six and modern-core renderers plus grouped fallback renderers for the remaining families.
- Normalized JSON session output.
- Full 2026+ family registry for `--list-values`.
- Dedicated evidence for `code_analyzer`, `data_processing`, `jargon`, `metrics`, `network_activity`, `system_monitoring`, `agent_workflows`, `platform_engineering`, `observability_ai_runtime`, `delivery_preview_ops`, and `supply_chain_security`.
- Experimental provider flags are parsed consistently and fail fast with an explicit not-implemented message.

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
- Experimental provider flags are parsed, but they do not enable runtime provider execution in this follower foundation.

## Gaps
- [GAPS.md](GAPS.md)
