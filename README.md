> [!WARNING]
> This repository is currently under active development.
> Contents may change without notice.
>
> This is a self-directed learning project focused on practical Go observability and
SRE workflows.

# Go Observability

Minimal Step 1 setup for a Go service flow:

- `gateway` (Gin) on `:8080`
- `worker` (`net/http`) on `:8081`
- `gateway -> worker` HTTP delegation via `POST /work`

## Quick Start

Run from repository root.

### Docker Compose

```bash
docker compose up --build
```

Gateway uses `WORKER_BASE_URL=http://worker:8081` in Compose so it can reach
worker over the container network.

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

## Verification

### Docker Compose path

curl -i -X POST http://localhost:8080/work

Expected:

- Status: 200 OK
- Body: {"result":"done"}

### Normal path

curl -i -X POST http://localhost:8080/work

Expected:

- Status: 200 OK
- Body: {"result":"done"}

### Failure path (worker stopped)

Stop worker, then run:

curl -i -X POST http://localhost:8080/work

Expected:

- Status: 502 Bad Gateway
- Body: {"error":"worker unavailable"}

## Optional Health Checks

curl -i http://localhost:8080/health
curl -i http://localhost:8081/health

## Project Links

- Portfolio: [PORTFOLIO.md](./PORTFOLIO.md)
- Issue Index: [doc/issues/INDEX.md](./doc/issues/INDEX.md)
- Issue Logs: [doc/issues](./doc/issues/)
