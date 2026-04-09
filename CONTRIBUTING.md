# Contributing

Do not replace the fail-fast scaffold with runtime parity until the Rust audit and traceability matrix are in place.

## Working commands
- `gofmt -l $(find . -name '*.go' -not -path './.git/*')`
- `golangci-lint run`
- `go build ./...`
- `go test ./...`
- `docker build -t go-stakeholder .`
- `docker run --rm go-stakeholder --list-values`
