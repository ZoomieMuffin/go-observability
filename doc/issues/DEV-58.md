# DEV-58: Containerize worker

## Date
2026-03-05

## Context
Step 2 requires worker to run as an isolated container so it can be composed with
gateway and observability services later.

## Changes
- Created docker/worker/Dockerfile for cmd/worker.
- Used a multi-stage build:
- build stage with golang:1.25.5-alpine
- runtime stage with alpine:3.20
- Built worker binary at /out/worker and copied it to /app/worker.
- Configured non-root runtime user appuser (uid 10001).
- Exposed port 8081 and set container entrypoint to /app/worker.

## Validation
- Built image successfully:
- docker build -f docker/worker/Dockerfile -t go-observability-worker:dev58 .
- Ran container successfully:
- docker run -d --rm -p 8081:8081 --name worker-dev58 go-observability-worker:dev58
- Confirmed health endpoint response:
- curl -i http://localhost:8081/health
- Result: HTTP/1.1 200 OK with {"status":"ok"}
- Confirmed work endpoint response:
- curl -i -X POST http://localhost:8081/work
- Result: HTTP/1.1 200 OK with {"result":"done"}
- Confirmed startup log:
- docker logs worker-dev58
- Result: worker start addr=:8081
- Stopped container:
- docker stop worker-dev58

## Evidence
- docker/worker/Dockerfile
- curl output for /health and /work
- container logs from worker-dev58

## Notes
- This issue covers worker containerization only.
- Compose integration is handled in DEV-59 and later issues.

Reviewer note: Final wording and formatting were explicitly reviewed by a human.
