# Observability / SRE 学習ノート

## 1. まず用語を固定する
目的:
- 同じ言葉を同じ意味で扱える状態にして、学習や実装時の認識ズレを防ぐ。

- `SRE (Site Reliability Engineering)`: 信頼性を設計・運用で担保する実践。
- `Observability`: メトリクス/ログ/トレースで状態を観測し、障害原因を特定するための考え方。
- `OpenTelemetry (OTel)`: 計測データの共通規格 + 実装（SDK/Collector）。
- `Prometheus`: メトリクス収集・保存システム。
- `PromQL`: Prometheusのクエリ言語。
- `Grafana`: 可視化・アラートのUI基盤。
- `Datadog`: 商用の統合Observability基盤。

## 2. 学習の軸（この順で進める）
目的:
- 「運用の型」を先に掴み、その後にベンダー非依存スキルへ展開するための学習順を固定する。

1. `Datadog`で監視運用の流れを掴む（収集→可視化→アラート→調査）。
2. `OpenTelemetry`でアプリ計測を実装する（metrics/traces/logs）。
3. `Prometheus + PromQL + Grafana`で同じ概念をOSSで再現する。

補足:
- DatadogとPromQLはクエリ言語が別。  
  目的は「同じ運用概念を別ツールで再現できること」。

## 3. 優先度マップ
目的:
- 何から学ぶべきかを明確化し、時間配分を最適化する。

凡例:
- `★★★★★`: 必須
- `★★★★⯪`: ほぼ必須
- `★★★★☆`: 高優先
- `★★★☆☆`: 案件次第で重要
- `★★⯪☆☆`: 補助

| 優先度 | 技術・ツール | 必要度 |
|---|---|---|
| 最優先 | Linux | ★★★★★ |
| 最優先 | Git / GitHub | ★★★★★ |
| 最優先 | Bash | ★★★★★ |
| 最優先 | Terraform | ★★★★★ |
| 高優先 | Python | ★★★★⯪ |
| 高優先 | Docker | ★★★★⯪ |
| 高優先 | Prometheus | ★★★★⯪ |
| 高優先 | Grafana | ★★★★⯪ |
| 高優先 | OpenTelemetry | ★★★★⯪ |
| 高優先 | Kubernetes | ★★★★☆ |
| 高優先 | PromQL | ★★★★☆ |
| 高優先 | Datadog | ★★★★☆ |
| 高優先 | GitHub Actions / GitLab CI | ★★★★☆ |
| 中優先 | Ansible | ★★★☆☆ |
| 中優先 | PagerDuty / Opsgenie | ★★★☆☆ |
| 中優先 | OpenSearch / Elasticsearch | ★★★☆☆ |
| 中優先 | k6 / JMeter | ★★★☆☆ |
| 補助 | Mermaid | ★★⯪☆☆ |
| 補助 | draw.io | ★★⯪☆☆ |
| 補助 | Tera Term（SSHクライアント） | ★★☆☆☆ |

## 4. 実践課題（Bash + Python → Terraform/PromQL/Grafana → Go）
目的:
- 監視を「見るだけ」で終わらせず、実装・自動化・復旧運用まで一連で再現する。

課題名:
- `SRE Mini Platform: SLO監視 + 障害対応自動化`

最終ゴール:
1. Go APIが`/metrics`を公開している。
2. Prometheusで収集できる。
3. Grafanaで可視化・アラートできる。
4. TerraformでGrafana設定を再現できる。
5. Bash/Pythonで運用作業を自動化できる。

### 4.1 ディレクトリを一気に作る（手打ち削減）
目的:
- 作業開始時の手作業ミスを減らし、標準構成を即座に用意する。

```bash
mkdir -p sre-mini-platform/{app-go,monitoring/prometheus,monitoring/grafana,infra/terraform-grafana,scripts,python,runbooks,evidence,reports}
cd sre-mini-platform
git init
```

### 4.2 作るもの
目的:
- どの成果物が揃えば「SREっぽい運用」が成立するかを先に固定する。

1. `Go API`
- エンドポイント: `/healthz`, `/work`, `/metrics`
- `/work`は`delay_ms`と`fail_rate`で遅延/エラーを注入

2. `Prometheus`
- Go APIをscrape
- ルール: `HighErrorRate`, `HighLatencyP95`, `TargetDown`

3. `Grafana`
- パネル: `RPS`, `Error Rate`, `p95 Latency`
- アラート通知先を1つ設定（ローカルWebhookでも可）

4. `Terraform`
- Grafana datasource/dashboard/alertをコード化

5. `Bash + Python`
- Bash: `deploy.sh`, `smoke.sh`, `rollback.sh`
- Python: `load.py`（負荷注入）, `sli_report.py`（SLIレポート）

## 5. 実行手順の正本
目的:
- 手順の重複管理をやめ、更新漏れを防ぐ。

実行手順は以下を正本とする:
- `todo.md`

運用ルール:
- 実装順、コマンド、チェックリストの更新は `todo.md` のみで行う。
- このドキュメント（`spec.md`）には方針・学習軸・完了基準のみを記載する。

## 6. 成果物方針
目的:
- 何を提出すれば完了かを固定し、やりっぱなしを防ぐ。

- `README.md`: 起動手順
- `runbooks/*.md`: 障害対応手順
- `reports/sli-*.md`: SLI実績
- `evidence/`: 証跡一式
- `FINAL_REPORT.md`: 最終まとめ

## 7. 合格条件（Definition of Done）
目的:
- 完了判定を曖昧にせず、必須要件を満たした状態で区切る。

1. `/metrics`がPrometheusで収集されている。
2. Grafanaに3パネル以上ある。
3. 3種類のアラートが期待通りに発報する。
4. Terraform applyでGrafana設定が再現できる。
5. Bash/Pythonで運用作業が自動化されている。
6. runbook/postmortem/evidenceが揃っている。
