# Go Backend - API Server

Echo フレームワークを使用した REST API サーバーです。

---

## 📁 プロジェクト構造

```
go-app/
├── cmd/
│   └── api/
│       └── main.go           # アプリケーションエントリーポイント
├── internal/
│   ├── config/
│   │   └── config.go         # 設定管理
│   ├── handler/
│   │   ├── health.go         # ヘルスチェックハンドラー
│   │   ├── health_test.go    # ヘルスチェックのテスト
│   │   ├── task.go           # タスクハンドラー
│   │   └── task_test.go      # タスクのテスト
│   ├── model/
│   │   └── task.go           # データモデル
│   └── router/
│       └── router.go         # ルーティング設定
├── go.mod
├── go.sum
└── README.md
```

### アーキテクチャの特徴

- **cmd/api**: アプリケーションのエントリーポイント
- **internal/**: 外部に公開しないパッケージ（Go標準のプラクティス）
  - **config/**: 環境変数の管理
  - **handler/**: HTTPリクエストハンドラー（ビジネスロジック）
  - **model/**: データ構造の定義
  - **router/**: ルーティングとミドルウェアの設定

---

## 🚀 ローカル開発

### 前提条件
- Go 1.23 以上

### セットアップ

1. **依存関係のインストール**

```bash
go mod download
```

2. **サーバーの起動**

```bash
go run cmd/api/main.go
```

3. **動作確認**

```bash
curl http://localhost:8080/health
```

---

## 🧪 テスト

### 全テストを実行

```bash
go test ./...
```

### カバレッジ付きで実行

```bash
go test -cover ./...
```

### 特定のパッケージをテスト

```bash
go test ./internal/handler/...
```

### 詳細なカバレッジレポート

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 📝 API エンドポイント

### Health Check
```http
GET /health
```

**レスポンス:**
```json
{
  "status": "ok",
  "message": "API is running"
}
```

### タスク一覧取得
```http
GET /tasks
```

**レスポンス:**
```json
[
  {
    "id": 1,
    "title": "Buy milk"
  },
  {
    "id": 2,
    "title": "Write report"
  }
]
```

### タスク作成
```http
POST /tasks
Content-Type: application/json

{
  "title": "New Task"
}
```

**レスポンス:**
```json
{
  "id": 3,
  "title": "New Task"
}
```

---

## 🔧 環境変数

| 変数名 | デフォルト | 説明 |
|--------|-----------|------|
| `PORT` | 8080 | APIサーバーのポート |
| `DB_HOST` | localhost | データベースホスト |
| `DB_USER` | postgres | データベースユーザー |
| `DB_PASSWORD` | postgres | データベースパスワード |
| `DB_NAME` | next-go-task-db | データベース名 |
| `FRONTEND_URL` | http://localhost:3000 | フロントエンドURL (CORS用) |

---

## 🏗️ 開発時のポイント

### ホットリロード (Air)

開発効率を上げるため、Airを使用してホットリロードを実現できます。

```bash
# Airのインストール
go install github.com/cosmtrek/air@latest

# 起動
air
```

### .air.toml の設定例

```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ./cmd/api"
  bin = "tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_dir = ["assets", "tmp", "vendor"]
  delay = 1000
```

---

## 📦 依存パッケージ

- `github.com/labstack/echo/v4`: 軽量・高速なWebフレームワーク
- `github.com/stretchr/testify`: テストアサーション

---

## 🔐 セキュリティ

- CORS設定: フロントエンドのオリジンのみ許可
- 環境変数: 機密情報はハードコードしない
- バリデーション: 入力値の検証を実施

---

## 📚 参考資料

- [Echo Framework Documentation](https://echo.labstack.com/)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Go Testing Best Practices](https://go.dev/doc/tutorial/add-a-test)

- **ポート競合エラー**:
  - 他のプロセスがポート `8080` を使用している場合、`docker-compose.yml` 内の `ports` セクションを変更してください（例: `- "3000:8080"`）。

- **依存関係のエラー**:
  - `go.mod` や `go.sum` が正しく設定されているか確認してください。
  - 必要に応じて `docker-compose build` を再実行してください。

---

これで Docker コンテナ上での Go 環境のセットアップと API サーバーの動作確認が完了します。
