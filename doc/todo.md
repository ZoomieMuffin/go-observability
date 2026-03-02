# Go Observability ハンズオン TODO

最終更新: 2026-02-27

このファイルを「実行手順の正本」とする。
方針・学習軸は `spec.md` を参照。

## Step 1: Go + net/http で最小 API を作る
- [ ] gateway サービスを作る
- [ ] worker サービスを作る
- [ ] gateway から worker を HTTP で呼ぶ

## Step 2: Docker Compose でローカル環境を作る
- [ ] gateway コンテナ
- [ ] worker コンテナ
- [ ] OpenTelemetry Collector コンテナ
- [ ] Jaeger コンテナ
- [ ] Prometheus コンテナ
- [ ] Grafana コンテナ
- [ ] `docker compose up` で全部起動できることを確認

## Step 3: OpenTelemetry の初期化を入れる
- [ ] `service.name` を設定する
- [ ] TracerProvider を初期化する
- [ ] MeterProvider を初期化する
- [ ] gateway と worker から Collector に送れることを確認
- [ ] shutdown 処理を入れる

## Step 4: HTTP 自動計測を入れる (otelhttp)
- [ ] サーバ側に otelhttp middleware を入れる
- [ ] クライアント側に otelhttp transport を入れる
- [ ] Jaeger で gateway -> worker のトレースが 1 本につながることを確認

## Step 5: 手動 span を入れる
- [ ] `validate` span を追加する
- [ ] `db.query` span を追加する
- [ ] `external.api` span を追加する
- [ ] span に Attributes を付ける
- [ ] span に Events を付ける

## Step 6: 失敗と遅延を再現する
- [ ] worker でランダム sleep を入れる
- [ ] worker で一定確率で 500 エラーを返す
- [ ] gateway の timeout を短くしてタイムアウトを再現する

## Step 7: Jaeger でトレースを確認する
- [ ] 遅い span がどれか見る
- [ ] エラーがどこで出たか見る
- [ ] gateway と worker が正しくつながっているか確認する

## Step 8: メトリクスを入れる
- [ ] requests counter を出す
- [ ] errors counter を出す
- [ ] duration histogram を出す
- [ ] `/metrics` エンドポイントを公開する（Prometheus Exporter 経由 or Collector の prometheus exporter）

## Step 9: Prometheus でメトリクスを確認する
- [ ] リクエスト数を見る
- [ ] エラー数を見る
- [ ] レイテンシ分布を見る
- [ ] PromQL を `monitoring/prometheus/promql.md` に保存する

## Step 10: Grafana で可視化する
- [ ] ダッシュボード（RPS / Error Rate / p95 Latency）を作る
- [ ] JSON export を `monitoring/grafana/dashboard.json` に保存する

## Step 11: ログ相関を入れる
- [ ] structured logging を入れる
- [ ] ログに `trace_id` と `span_id` を出す
- [ ] ログから該当トレースを追えることを確認する

## Step 12: 負荷をかける
- [ ] k6 または hey で複数リクエストを流す
- [ ] 負荷時の遅延と失敗を観測する
- [ ] 例:
```bash
python3 python/load.py --rps 5 --duration 300 --delay-ms 0 --fail-rate 0.0
python3 python/load.py --rps 20 --duration 300 --delay-ms 700 --fail-rate 0.03
```

## Step 13: Postgres を入れる
- [ ] worker から DB に INSERT / SELECT する
- [ ] 疑似 db.query を本物の DB アクセスに置き換える
- [ ] DB の遅さを調べる

## Step 14: DB パフォーマンスを調べる
- [ ] インデックスなしで遅くする
- [ ] `EXPLAIN ANALYZE` で実行計画を見る
- [ ] インデックスを貼って速くなることを確認する

## Step 15: Collector 設定を触る
- [ ] サンプリングを入れる (tail sampling)
- [ ] 属性を整理する
- [ ] 不要データを落とす

## Step 16: Terraform で Grafana 設定をコード化する
- [ ] datasource をコード化
- [ ] dashboard をコード化
- [ ] alert rule をコード化
- [ ] alert の通知先（contact point / webhook 等）を設定する
- [ ] 例:
```bash
cd infra/terraform-grafana
terraform init
terraform plan
terraform apply
```

## Step 17: CI を入れる (GitHub Actions)
- [ ] push / PR で `gofmt` を回す
- [ ] push / PR で `golangci-lint` を回す
- [ ] push / PR で `go test` を回す
- [ ] push / PR で `go build` を回す

## Step 18: テストを入れる
- [ ] ロジックのユニットテストを書く
- [ ] API の簡単な疎通テストを書く
- [ ] E2E テストを 1 本書く

## Step 19: 証跡を残す
- [ ] 証跡ディレクトリを作る
```bash
mkdir -p evidence/$(date +%Y%m%d)/step-{01,02,03,04,05,06}
```
- [ ] 保存ルールを守る
  - 画面: `screenshot-*.png`
  - 実行ログ: `command.log`
  - APIレスポンス: `response-*.json`
  - アラート: `alert-*.json`
  - メモ: `notes.md`
- [ ] 作業ログを残す
```bash
script -a evidence/$(date +%Y%m%d)/step-01/command.log
# このあと作業
exit
```

## Step 20: 最終提出物を揃える
- [ ] `README.md`: 起動手順
- [ ] `runbooks/*.md`: 障害対応手順
- [ ] `reports/sli-*.md`: SLI実績
- [ ] `evidence/`: 証跡一式
- [ ] `FINAL_REPORT.md`: 最終まとめ
- [ ] `FINAL_REPORT.md` に以下を記載
  - 構成図（Mermaidかdraw.io）
  - SLO定義（例: p95 < 300ms, error rate < 1%）
  - 主要PromQL
  - アラート発報実績
  - 障害訓練の時系列（検知->復旧）
  - 改善アクション3件

## Definition of Done

- [ ] `/metrics` が Prometheus で収集されている
- [ ] Grafana に 3 パネル以上ある
- [ ] 3 種類のアラートが期待通りに発報する
- [ ] Terraform apply で Grafana 設定が再現できる
- [ ] Bash/Python で運用作業が自動化されている
- [ ] runbook/postmortem/evidence が揃っている
