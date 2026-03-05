# DEV-57: Containerize gateway

## Date
2026-03-05

## Context
Step 2 requires a containerized gateway service as the first building block of the
Docker Compose local environment.

## Changes
- Created docker/gateway/Dockerfile for cmd/gateway
- Built gateway binary in a multi-stage build and packaged it in a lightweight
runtime image
- Added default runtime environment variables:
- WORKER_BASE_URL=http://localhost:8081
- HTTP_TIMEOUT_MS=2000
- Exposed port 8080 and configured container entrypoint to run gateway

## Validation
- Built image successfully:
- docker build -f docker/gateway/Dockerfile -t go-observability-gateway:dev57 .
- Ran container successfully:
- docker run -d --rm -p 8080:8080 --name gateway-dev57 -e WORKER_BASE_URL=http://
worker:8081 -e HTTP_TIMEOUT_MS=1500 go-observability-gateway:dev57
- Confirmed health endpoint response:
- curl -i http://localhost:8080/health
- Result: HTTP/1.1 200 OK
- Confirmed env injection from logs:
- worker=http://worker:8081 timeout_ms=1500

## Evidence
- docker/gateway/Dockerfile
- Container logs from gateway-dev57
- Health check output from curl -i http://localhost:8080/health

## Notes
- This issue covers gateway containerization only.
- Compose integration is handled in DEV-59 and later issues.

Reviewer note: Formatting cleanup was explicitly requested by a human reviewer.
