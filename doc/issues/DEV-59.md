# DEV-59: Create base Docker Compose setup

## Date
2026-03-06

## Context
Step 2 needs a minimal Docker Compose setup so gateway and worker can start together as the base for later observability services.

## Changes
- Added compose.yaml at repository root.
- Defined worker service using docker/worker/Dockerfile.
- Defined gateway service using docker/gateway/Dockerfile.
- Published ports 8081:8081 for worker and 8080:8080 for gateway.
- Added depends_on from gateway to worker.
- Set WORKER_BASE_URL=http://worker:8081 and HTTP_TIMEOUT_MS=2000 for gateway.
- Added Docker Compose startup and verification steps to README.md.

## Validation
- Validated Compose configuration:
- docker compose config
- Started both services:
- docker compose up --build
- Confirmed both containers were running:
- docker compose ps
- Result: gateway and worker were both Up
- Confirmed gateway startup configuration:
- docker compose logs gateway --tail=50
- Result: gateway start addr=:8080 worker=http://worker:8081 timeout_ms=2000
- Confirmed end-to-end connectivity:
- curl -i -X POST http://localhost:8080/work
- Result: HTTP/1.1 200 OK with {"result":"done"}

## Evidence
- compose.yaml
- README.md
- docker compose ps
- docker compose logs gateway --tail=50
- curl -i -X POST http://localhost:8080/work

## Notes
- Initial startup failed because host port 8080 was already allocated by an older container.
- After stopping the conflicting container, the Compose stack started normally.

Reviewer note: Final wording and formatting were explicitly reviewed by a human.
