# GraphQL API プロジェクト構造

このドキュメントは、GraphQL + REST API のプロジェクト構造を説明します。

## 📁 プロジェクト構造

```
go-app/
├── cmd/
│   └── api/
│       └── main.go              # エントリーポイント
├── internal/                    # 内部パッケージ（外部から利用不可）
│   ├── config/                  # 設定管理
│   │   └── config.go
│   ├── domain/                  # ドメインモデル（ビジネスルール）
│   │   └── task.go
│   ├── handler/                 # REST APIハンドラー
│   │   ├── health.go
│   │   ├── health_test.go
│   │   ├── task.go
│   │   └── task_test.go
│   ├── repository/              # データアクセス層
│   │   ├── task_repository.go            # インターフェース
│   │   └── task_repository_memory.go     # In-Memory実装
│   ├── service/                 # ビジネスロジック層
│   │   └── task_service.go
│   ├── router/                  # ルーティング設定
│   │   └── router.go
│   └── infrastructure/          # 外部依存（DB等）
│       └── database/
│           └── postgres.go      # PostgreSQL接続（今後実装）
├── pkg/                         # 外部公開可能なパッケージ
│   └── graph/                   # GraphQL関連
│       ├── gqlgen.yml           # gqlgen設定
│       ├── schema/
│       │   └── schema.graphql   # GraphQLスキーマ定義
│       ├── resolver/
│       │   ├── resolver.go      # リゾルバー基底
│       │   └── schema.resolvers.go  # 生成されるリゾルバー
│       ├── model/
│       │   ├── model.go         # カスタムモデル
│       │   └── models_gen.go    # 生成されるモデル
│       └── generated/
│           └── generated.go     # gqlgen生成コード
├── docs/                        # ドキュメント
│   └── ARCHITECTURE.md          # このファイル
├── go.mod
├── go.sum
├── .air.toml                    # ホットリロード設定
├── .env.example
├── .dockerignore
├── .gitignore
└── README.md
```

---

## 🏗️ アーキテクチャ概要

### レイヤー構成

```
┌─────────────────────────────────────────────┐
│         Presentation Layer                  │
│  ┌──────────────┐     ┌──────────────┐     │
│  │  REST API    │     │   GraphQL    │     │
│  │  (Handler)   │     │  (Resolver)  │     │
│  └──────────────┘     └──────────────┘     │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│           Service Layer                     │
│         (Business Logic)                    │
│        ┌──────────────┐                     │
│        │ TaskService  │                     │
│        └──────────────┘                     │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│          Domain Layer                       │
│        (Business Rules)                     │
│        ┌──────────────┐                     │
│        │    Task      │                     │
│        └──────────────┘                     │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│         Repository Layer                    │
│        (Data Access)                        │
│    ┌──────────────────────┐                │
│    │  TaskRepository      │                │
│    │  - Memory (current)  │                │
│    │  - PostgreSQL (TODO) │                │
│    └──────────────────────┘                │
└─────────────────────────────────────────────┘
```

---

## 📦 各層の責務

### 1. **cmd/api**
- アプリケーションのエントリーポイント
- 依存関係の初期化と注入
- サーバーの起動

### 2. **internal/domain**
- ビジネスルールを含むドメインモデル
- フレームワークに依存しない純粋なビジネスロジック
- エンティティとバリューオブジェクト

### 3. **internal/service**
- ユースケースの実装
- ドメインモデルとリポジトリを組み合わせたビジネスロジック
- トランザクション管理

### 4. **internal/repository**
- データ永続化の抽象化
- インターフェースと実装を分離
- 現在: In-Memory実装
- 将来: PostgreSQL実装

### 5. **internal/handler** (REST API)
- HTTPリクエストのハンドリング
- リクエスト/レスポンスの変換
- エラーハンドリング

### 6. **pkg/graph** (GraphQL)
- GraphQLスキーマの定義
- リゾルバーの実装
- gqlgenによるコード生成

### 7. **internal/router**
- ルーティング設定
- ミドルウェアの設定
- REST と GraphQL の統合

### 8. **internal/infrastructure**
- 外部システムとの連携
- データベース接続
- 外部API クライアント

---

## 🔄 データフロー

### REST API の場合
```
HTTP Request
    ↓
Handler (internal/handler)
    ↓
Service (internal/service)
    ↓
Domain Model (internal/domain)
    ↓
Repository (internal/repository)
    ↓
Database
```

### GraphQL の場合
```
GraphQL Query/Mutation
    ↓
Resolver (pkg/graph/resolver)
    ↓
Service (internal/service)
    ↓
Domain Model (internal/domain)
    ↓
Repository (internal/repository)
    ↓
Database
```

---

## 🚀 GraphQL コード生成

### gqlgen のセットアップ

1. **依存関係のインストール**
```bash
go get github.com/99designs/gqlgen
go get github.com/99designs/gqlgen/graphql/handler
go get github.com/99designs/gqlgen/graphql/playground
```

2. **GraphQL コードの生成**
```bash
cd pkg/graph
go run github.com/99designs/gqlgen generate
```

3. **生成されるファイル**
- `pkg/graph/generated/generated.go` - GraphQLサーバーの実装
- `pkg/graph/model/models_gen.go` - GraphQL型のGoモデル
- `pkg/graph/resolver/schema.resolvers.go` - リゾルバーのスタブ

---

## 📝 API エンドポイント

### REST API

#### Health Check
```http
GET /health
GET /api/health
```

#### Tasks
```http
GET /tasks
GET /api/tasks
POST /api/tasks
```

### GraphQL

#### Playground
```
GET /playground
```

#### Query Endpoint
```
POST /query
```

#### サンプルクエリ

**すべてのタスクを取得**
```graphql
query {
  tasks {
    id
    title
    description
    completed
    createdAt
    updatedAt
  }
}
```

**特定のタスクを取得**
```graphql
query {
  task(id: "1") {
    id
    title
    description
    completed
  }
}
```

**タスクを作成**
```graphql
mutation {
  createTask(input: {
    title: "New Task"
    description: "Task description"
  }) {
    id
    title
    description
    completed
    createdAt
  }
}
```

**タスクを更新**
```graphql
mutation {
  updateTask(
    id: "1"
    input: {
      title: "Updated Title"
      completed: true
    }
  ) {
    id
    title
    completed
    updatedAt
  }
}
```

**タスクを削除**
```graphql
mutation {
  deleteTask(id: "1")
}
```

---

## 🔧 開発ガイド

### 新しいエンティティの追加手順

1. **ドメインモデルを定義** (`internal/domain/`)
2. **リポジトリインターフェースを定義** (`internal/repository/`)
3. **サービス層を実装** (`internal/service/`)
4. **GraphQLスキーマを更新** (`pkg/graph/schema/schema.graphql`)
5. **gqlgen でコード生成**
6. **リゾルバーを実装** (`pkg/graph/resolver/`)
7. **(オプション) REST ハンドラーを追加** (`internal/handler/`)

---

## 🧪 テスト戦略

### ユニットテスト
- **Domain**: ビジネスロジックのテスト
- **Service**: モックリポジトリを使用
- **Handler**: HTTPテスト
- **Resolver**: GraphQLリゾルバーのテスト

### 統合テスト
- データベースを含むエンドツーエンドテスト
- テストコンテナ (testcontainers) の使用

---

## 🔐 セキュリティ

### 実装予定

1. **認証**
   - JWT トークンベース認証
   - GraphQL ディレクティブ `@auth`

2. **認可**
   - ロールベースアクセス制御 (RBAC)
   - ミドルウェアでの権限チェック

3. **バリデーション**
   - 入力値の検証
   - GraphQL スキーマレベルでの制約

---

## 📚 参考資料

- [gqlgen Documentation](https://gqlgen.com/)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Domain-Driven Design](https://martinfowler.com/tags/domain%20driven%20design.html)
