# Go Observability Portfolio

## Summary
Self-directed project focused on practical Go observability and SRE workflows.

## What I Have Implemented
- Go services: `gateway` (Gin) and `worker` (`net/http`)
- Service-to-service HTTP delegation (`gateway` -> `worker`) with timeout/error
handling
- Reproducible manual verification flow (normal path `200`, failure path `502`)
- CI baseline with GitHub Actions (`go test`, `go vet`, `go build`)
- Onboarding-oriented startup and verification documentation updates (DEV-55 in
progress)

## What I’m Building (Roadmap Scope)
- Docker Compose local stack (gateway, worker, OTel Collector, Jaeger, Prometheus,
Grafana)
- OpenTelemetry instrumentation (auto + manual spans), log correlation, and metrics
design
- Prometheus/PromQL validation and Grafana dashboards/alerts
- Load testing, Postgres integration, and DB performance investigation
- Terraform-based Grafana provisioning, test expansion, evidence/report packaging, and
final deliverables

## Proof Links
- README: [README.md](./README.md)
- Issue index: [doc/issues/INDEX.md](./doc/issues/INDEX.md)
- Issue logs: [doc/issues](./doc/issues/)

## Working Style & Quality Assurance
I use AI-assisted workflows for speed, but all architecture decisions, testing, and
final quality checks are performed by me.

