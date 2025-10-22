# GraphQL セットアップガイド

このガイドでは、gqlgenを使用したGraphQL APIのセットアップ手順を説明します。

---

## 📋 前提条件

- Go 1.23 以上
- Git

---

## 🚀 セットアップ手順

### 1. 依存関係のインストール

```bash
cd go-app

# gqlgen と関連パッケージのインストール
go get github.com/99designs/gqlgen@latest
go get github.com/99designs/gqlgen/graphql/handler
go get github.com/99designs/gqlgen/graphql/playground

# その他の依存関係
go mod tidy
```

### 2. GraphQL コードの生成

```bash
# pkg/graph ディレクトリに移動
cd pkg/graph

# gqlgen でコードを生成
go run github.com/99designs/gqlgen generate
```

生成されるファイル:
- `generated/generated.go` - GraphQLサーバーの実装
- `model/models_gen.go` - GraphQL型のGoモデル

### 3. リゾルバーの実装

`pkg/graph/resolver/schema.resolvers.go` を編集して、各クエリとミューテーションのロジックを実装します。

**例: Query.tasks**
```go
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
    tasks, err := r.TaskService.GetAllTasks(ctx)
    if err != nil {
        return nil, err
    }
    
    // domain.Task を model.Task に変換
    result := make([]*model.Task, len(tasks))
    for i, task := range tasks {
        result[i] = &model.Task{
            ID:          task.ID,
            Title:       task.Title,
            Description: task.Description,
            Completed:   task.Completed,
            CreatedAt:   task.CreatedAt.Format(time.RFC3339),
            UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
        }
    }
    
    return result, nil
}
```

### 4. サーバーの起動

```bash
# go-app ディレクトリから実行
cd ../..
go run cmd/api/main.go
```

---

## 🧪 GraphQL Playground で動作確認

サーバー起動後、ブラウザで以下にアクセス:

```
http://localhost:8080/playground
```

### サンプルクエリ

#### すべてのタスクを取得
```graphql
query GetAllTasks {
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

#### タスクを作成
```graphql
mutation CreateTask {
  createTask(input: {
    title: "Buy groceries"
    description: "Milk, bread, eggs"
  }) {
    id
    title
    description
    completed
    createdAt
  }
}
```

#### タスクを更新
```graphql
mutation UpdateTask {
  updateTask(
    id: "1"
    input: {
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

#### タスクを削除
```graphql
mutation DeleteTask {
  deleteTask(id: "1")
}
```

---

## 📝 スキーマの変更手順

1. `pkg/graph/schema/schema.graphql` を編集
2. `go run github.com/99designs/gqlgen generate` を実行
3. 生成された `schema.resolvers.go` のスタブを実装
4. サーバーを再起動して確認

---

## 🔧 開発Tips

### ホットリロード (Air使用)

```bash
# Airのインストール
go install github.com/cosmtrek/air@latest

# go-app ディレクトリで実行
air
```

### GraphQLスキーマのバリデーション

```bash
cd pkg/graph
go run github.com/99designs/gqlgen validate
```

---

## 📚 参考資料

- [gqlgen Documentation](https://gqlgen.com/)
- [GraphQL Spec](https://spec.graphql.org/)
- [GraphQL Best Practices](https://graphql.org/learn/best-practices/)

---

## ⚠️ 注意事項

### 現在の実装

- **データストレージ**: In-Memory (メモリ上)
- **永続化**: なし（サーバー再起動でデータ消失）

### 今後の実装予定

- PostgreSQL との連携
- 認証・認可
- ページネーション
- サブスクリプション (リアルタイム更新)
