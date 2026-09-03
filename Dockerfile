FROM golang:1.27 AS build
WORKDIR /src
COPY . .
RUN go test ./...
RUN go build -o /out/stakeholder ./cmd/stakeholder

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/stakeholder /stakeholder
ENTRYPOINT ["/stakeholder"]
