# Tooling

## Standard commands
- `gofmt -l $(find . -name '*.go' -not -path './.git/*')`
- `golangci-lint run`
- `go build ./...`
- `go test ./...`
- `docker build -t go-stakeholder .`
- `docker run --rm go-stakeholder --list-values`

## Notes
- `gofmt` is the formatter gate.
- `golangci-lint` is the primary linter; the workflow installs it explicitly.
