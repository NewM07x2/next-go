# プロジェクト構造

```
go-app/
├── cmd/
│   └── api/
│       └── main.go              # エントリーポイント
│
├── internal/                    # 内部パッケージ（外部から利用不可）
│   ├── config/                  # 設定管理
│   │   └── config.go            # 環境変数の読み込み
│   │
│   ├── domain/                  # ドメインモデル（ビジネスルール）
│   │   └── task.go              # Taskエンティティ
│   │
│   ├── handler/                 # REST APIハンドラー
│   │   ├── health.go            # ヘルスチェック
│   │   ├── health_test.go
│   │   ├── task.go              # タスク関連のREST API
│   │   └── task_test.go
│   │
│   ├── repository/              # データアクセス層
│   │   ├── task_repository.go            # インターフェース定義
│   │   └── task_repository_memory.go     # In-Memory実装
│   │
│   ├── service/                 # ビジネスロジック層
│   │   └── task_service.go      # タスク関連のビジネスロジック
│   │
│   ├── router/                  # ルーティング設定
│   │   └── router.go            # REST + GraphQL ルーティング
│   │
│   └── infrastructure/          # 外部依存（DB、外部API等）
│       └── database/
│           └── postgres.go      # PostgreSQL接続（今後実装）
│
├── pkg/                         # 外部公開可能なパッケージ
│   └── graph/                   # GraphQL関連
│       ├── gqlgen.yml           # gqlgen設定ファイル
│       │
│       ├── schema/
│       │   └── schema.graphql   # GraphQLスキーマ定義
│       │
│       ├── resolver/
│       │   ├── resolver.go      # リゾルバー基底
│       │   └── schema.resolvers.go  # 生成されるリゾルバー
│       │
│       ├── model/
│       │   ├── model.go         # カスタムモデル
│       │   └── models_gen.go    # 生成されるモデル
│       │
│       └── generated/
│           └── generated.go     # gqlgen生成コード
│
├── docs/                        # ドキュメント
│   ├── ARCHITECTURE.md          # アーキテクチャ説明
│   ├── GRAPHQL_SETUP.md         # GraphQLセットアップガイド
│   └── PROJECT_STRUCTURE.md     # このファイル
│
├── go.mod                       # Go モジュール定義
├── go.sum                       # 依存ハッシュ
├── .air.toml                    # ホットリロード設定
├── .env.example                 # 環境変数サンプル
├── .dockerignore                # Docker ビルドの除外設定
├── .gitignore                   # Git除外設定
└── README.md                    # プロジェクトREADME
```

---

## 📁 ディレクトリの役割

### `cmd/`
アプリケーションのエントリーポイント。複数のアプリケーションがある場合、`cmd/api/`, `cmd/worker/` のように分ける。

### `internal/`
Go標準の規約で、外部パッケージから参照できない内部パッケージ。

#### `internal/config/`
環境変数や設定ファイルの管理。

#### `internal/domain/`
ビジネスルールを含むドメインモデル。フレームワークやインフラに依存しない。

#### `internal/service/`
ユースケースの実装。ドメインモデルとリポジトリを組み合わせる。

#### `internal/repository/`
データアクセス層。インターフェースと実装を分離。

#### `internal/handler/`
REST APIのHTTPハンドラー。リクエスト/レスポンスの変換。

#### `internal/router/`
ルーティング設定とミドルウェアの適用。

#### `internal/infrastructure/`
外部システム（DB、外部API等）との連携。

### `pkg/`
外部パッケージから参照可能な公開パッケージ。GraphQLなど。

### `docs/`
プロジェクトのドキュメント。

---

## 🔄 依存関係の方向

```
cmd/api
  ↓
router → handler → service → repository → infrastructure
  ↓                ↓
resolver ────────→ domain
```

**依存のルール:**
- 上位層は下位層に依存できる
- 下位層は上位層に依存しない
- `domain` はどこにも依存しない（純粋なビジネスロジック）

---

## 🚀 開発フロー

### 新機能の追加手順

1. **ドメインモデルを定義** (`internal/domain/`)
2. **リポジトリインターフェースを定義** (`internal/repository/`)
3. **リポジトリ実装を追加** (`internal/repository/*_memory.go`)
4. **サービス層を実装** (`internal/service/`)
5. **GraphQLスキーマを更新** (`pkg/graph/schema/schema.graphql`)
6. **gqlgen でコード生成**
7. **リゾルバーを実装** (`pkg/graph/resolver/`)
8. **(オプション) REST ハンドラーを追加** (`internal/handler/`)
9. **テストを作成**

---

## 🧪 テストの配置

テストファイルは、対象のパッケージと同じディレクトリに配置:
- `internal/handler/task_test.go`
- `internal/service/task_service_test.go`
- `internal/repository/task_repository_test.go`

---

## 📦 生成されるファイル

以下のファイルは gqlgen によって自動生成されるため、手動編集しない:
- `pkg/graph/generated/generated.go`
- `pkg/graph/model/models_gen.go`

手動で実装するファイル:
- `pkg/graph/resolver/resolver.go` (依存注入)
- `pkg/graph/resolver/schema.resolvers.go` (ビジネスロジック)

---

## 🔐 今後の拡張

### データベース統合
`internal/infrastructure/database/` に PostgreSQL接続を実装し、
`internal/repository/task_repository_postgres.go` を追加。

### 認証・認可
`internal/middleware/` を追加して、JWT認証ミドルウェアを実装。

### キャッシュ
`internal/infrastructure/cache/` に Redis接続を実装。
