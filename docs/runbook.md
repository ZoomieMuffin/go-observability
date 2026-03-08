# Runbook

## Purpose

This runbook covers the current local Docker Compose stack.

## Start

```bash
docker compose up --build
```

## Stop

```bash
docker compose down
```

## Verify Service Status

```bash
docker compose ps
curl -i http://localhost:8080/health
curl -i http://localhost:8081/health
docker compose logs otel-collector --tail=50
```

## Basic Failure Check

If `POST /work` returns `502`:

1. Check whether `worker` is running with `docker compose ps`.
2. Inspect `gateway` logs with `docker compose logs gateway --tail=50`.
3. Inspect `worker` logs with `docker compose logs worker --tail=50`.
4. Restart the stack with `docker compose down && docker compose up --build`.

## Notes

- No production alerting is configured yet.
- Telemetry export is not implemented yet; the collector is only prepared as local infrastructure.
