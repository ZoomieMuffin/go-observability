  前提: リポジトリルート（`<repo-root>`）で実行する

  1. ターミナルAで worker 起動
  go run ./cmd/worker

  2. ターミナルBで gateway 起動
  HTTP_TIMEOUT_MS=2000 WORKER_BASE_URL=http://localhost:8081 go run ./cmd/gateway

  3. 正常系確認（ターミナルC）
  curl -i -X POST http://localhost:8080/work
  期待: 200 / {"result":"done"}

  4. 異常系確認
  worker を停止してから:
  curl -i -X POST http://localhost:8080/work
  期待: 502 / {"error":"worker unavailable"}

## 正常系疎通確認（curl）

  前提:
  - リポジトリルートで実行する
  - `worker` が起動している
  - `gateway` が起動している（`WORKER_BASE_URL=http://localhost:8081`）

  手順:
  1. `gateway` の `/work` を呼び出す

  ```bash
  curl -i -X POST http://localhost:8080/work

  期待結果:

  - HTTPステータスが 200 OK
  - レスポンスボディが {"result":"done"}

  確認観点:

  - gateway -> worker のHTTP呼び出しが成功していること

## 正常系疎通確認（curl）

  前提:
  - リポジトリルートで実行する
  - `worker` が起動している
  - `gateway` が起動している（`WORKER_BASE_URL=http://localhost:8081`）

  手順:
  1. `gateway` の `/work` を呼び出す

  ```bash
  curl -i -X POST http://localhost:8080/work

  期待結果:

  - HTTPステータスが 200 OK
  - レスポンスボディが {"result":"done"}

  確認観点:

  - gateway -> worker のHTTP呼び出しが成功していること


上記手順と前提条件に従えば、正常系（200）と異常系（502）を誰でも再現可能
