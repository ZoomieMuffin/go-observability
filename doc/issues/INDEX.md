# Go-Observability Issue Index

Last updated: 2026-03-04 (JST)
Source of truth: Linear project `Go-Observability`

## Current Snapshot
- In progress: `DEV-55`
- Done: `DEV-51`, `DEV-52`, `DEV-53`, `DEV-54`, `DEV-56`
- Backlog: `DEV-30` and `DEV-31` to `DEV-50`

## Step 1 Rollup (Parent: DEV-30)

| Issue | Title | Linear Status | Local Log |
|---|---|---|---|
| DEV-30 | Step 1: Go + net/http で最小 API を作る | Backlog | [DEV-30](./DEV-30.md) |
| DEV-51 | gateway skeleton を作る (gin) | Done | [DEV-51](./DEV-51.md) |
| DEV-52 | worker skeleton を作る (net/http) | Done | [DEV-52](./DEV-52.md) |
| DEV-53 | gateway から worker を HTTP 呼び出しする | Done | [DEV-53](./DEV-53.md) |
| DEV-54 | 疎通テスト（手動/E2E最小）を追加する | Done | [DEV-54](./DEV-54.md) |
| DEV-55 | README に起動方法と検証手順を追記する | In Progress | [DEV-55](./DEV-55.md) |

## Parent Roadmap (DEV-30 to DEV-50)

| Issue | Summary | Linear Status |
|---|---|---|
| DEV-30 | Step 1: Minimal API (gateway/worker/http) | Backlog |
| DEV-31 | Step 2: Docker Compose local stack | Backlog |
| DEV-32 | Step 3: OpenTelemetry initialization | Backlog |
| DEV-33 | Step 4: HTTP auto instrumentation (otelhttp) | Backlog |
| DEV-34 | Step 5: Manual spans | Backlog |
| DEV-35 | Step 6: Failure and latency simulation | Backlog |
| DEV-36 | Step 7: Trace validation in Jaeger | Backlog |
| DEV-37 | Step 8: Metrics instrumentation | Backlog |
| DEV-38 | Step 9: Metrics validation in Prometheus | Backlog |
| DEV-39 | Step 10: Visualization in Grafana | Backlog |
| DEV-40 | Step 11: Log correlation | Backlog |
| DEV-41 | Step 12: Load testing | Backlog |
| DEV-42 | Step 13: Postgres integration | Backlog |
| DEV-43 | Step 14: DB performance analysis | Backlog |
| DEV-44 | Step 15: Collector tuning | Backlog |
| DEV-45 | Step 16: Grafana as code (Terraform) | Backlog |
| DEV-46 | Step 17: CI (GitHub Actions) | Backlog |
| DEV-47 | Step 18: Add tests | Backlog |
| DEV-48 | Step 19: Collect evidence | Backlog |
| DEV-49 | Step 20: Final deliverables | Backlog |
| DEV-50 | Definition of Done | Backlog |

## Additional Completed Work

| Issue | Title | Linear Status |
|---|---|---|
| DEV-56 | CI導入 (GitHub Actions) | Done |
