# Docker

The repo builds a static Linux binary and ships it in a distroless runtime image.

## Commands
- `docker build -t go-stakeholder .`
- `docker run --rm go-stakeholder --list-values`

## CI intent
- Build and test in a Go builder image.
- Smoke the runtime image with the CLI contract surface.
