> [!WARNING]
> This repository is currently under active development.
> Contents may change without notice.
>
> This is a self-directed learning project focused on practical Go observability and
SRE workflows.

# Go Observability

Minimal local service flow for Go observability practice:

- `gateway` (Gin) on `:8080`
- `worker` (`net/http`) on `:8081`
- `gateway -> worker` HTTP delegation via `POST /work`
- `otel-collector` for future OTLP ingest on `:4317` / `:4318`

## Quick Start

Run from repository root.

### Docker Compose

```bash
docker compose up --build
```

Gateway uses `WORKER_BASE_URL=http://worker:8081` in Compose so it can reach
worker over the container network.

Collector is available inside Compose at `http://otel-collector:4318`.

### Run services directly

```bash
# Terminal A
go run ./cmd/worker

# Terminal B
go run ./cmd/gateway

# Terminal B (explicit env values)
WORKER_BASE_URL=http://localhost:8081 HTTP_TIMEOUT_MS=2000 go run ./cmd/gateway
```

## Environment Variables (gateway)

| Name | Description | Default |
|---|---|---|
| WORKER_BASE_URL | Worker base URL used by gateway | http://localhost:8081 |
| HTTP_TIMEOUT_MS | Gateway HTTP timeout to worker (ms) | 2000 |

## Environment Variables (Compose)

| Name | Used by | Description |
|---|---|---|
| OTEL_SERVICE_NAME | gateway, worker | Logical service name for future OTel setup |
| OTEL_EXPORTER_OTLP_ENDPOINT | gateway, worker | Collector OTLP endpoint |
| OTEL_EXPORTER_OTLP_PROTOCOL | gateway, worker | OTLP transport protocol |

## Verification

### Docker Compose path

```bash
curl -i -X POST http://localhost:8080/work
```

Expected:

- Status: 200 OK
- Body: {"result":"done"}

### Normal path

```bash
curl -i -X POST http://localhost:8080/work
```

Expected:

- Status: 200 OK
- Body: {"result":"done"}

### Failure path (worker stopped)

Stop worker, then run:

```bash
curl -i -X POST http://localhost:8080/work
```

Expected:

- Status: 502 Bad Gateway
- Body: {"error":"worker unavailable"}

## Optional Health Checks

```bash
curl -i http://localhost:8080/health
curl -i http://localhost:8081/health
docker compose ps
docker compose logs otel-collector --tail=50
docker compose logs jaeger --tail=50
```

## Jaeger UI

After starting the Compose stack, open:

```text
http://localhost:16686
```

Expected:

- Jaeger UI loads successfully
- Traces will appear after application-side telemetry is added in a later step

## Docs

- [Architecture](./docs/architecture.md)
- [Runbook](./docs/runbook.md)
