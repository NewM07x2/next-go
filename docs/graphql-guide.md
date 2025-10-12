# GraphQL 使用ガイド

このドキュメントでは、GraphQLの基本的な使用方法や書き方、`graphql-codegen`の使用方法について説明します。

---

## GraphQL の基本

### **GraphQLとは？**

- GraphQLは、クライアントが必要なデータを正確に取得できるように設計されたクエリ言語です。
- REST APIの代替として使用されることが多く、柔軟で効率的なデータ取得を可能にします。

### **基本構文**

1. **クエリ (Query)**:

   データを取得するためのリクエスト。

   ```graphql
   query GetUser($id: ID!) {
     user(id: $id) {
       id
       name
       email
     }
   }
   ```

2. **ミューテーション (Mutation)**:

   データを作成、更新、削除するためのリクエスト。

   ```graphql
   mutation CreateUser($name: String!, $email: String!) {
     createUser(name: $name, email: $email) {
       id
       name
     }
   }
   ```

3. **サブスクリプション (Subscription)**:

   リアルタイムでデータの変更を受け取るためのリクエスト。

   ```graphql
   subscription OnUserCreated {
     userCreated {
       id
       name
     }
   }
   ```

---

## GraphQL Code Generator の使用方法

`graphql-codegen`は、GraphQLスキーマやクエリから型定義やコードを自動生成するツールです。

### **インストール**

以下のコマンドで`graphql-codegen`をインストールします。

```bash
npm install -D @graphql-codegen/cli
```

### **初期設定**

初期設定を行うには、以下のコマンドを実行します。

```bash
npx graphql-codegen init
```

プロンプトに従って以下を設定します：

- **What type of application are you building?**: `Application built with React`
- **Where is your schema?**: `src/schema.graphql`
- **Where are your operations and fragments?**: `src/**/*.tsx`
- **Where to write the output**: `src/graphql/generated/`
- **Do you want to generate an introspection file?**: `Yes`
- **How to name the config file?**: `codegen.ts`
- **What script in package.json should run the codegen?**: `codegen`

### **コード生成の実行**

設定が完了したら、以下のコマンドでコードを生成します。

```bash
npm run codegen
```

### **生成されるファイル**

- **型定義ファイル**:

  GraphQLスキーマやクエリに基づいたTypeScriptの型定義が生成されます。

  例: `src/graphql/generated/graphql.ts`

- **React用のフック**:

  クエリやミューテーションに対応するReactフックが生成されます。

  例:

  ```typescript
  import { useQuery } from '@apollo/client';
  import { GetUserDocument } from './generated/graphql';

  const { data, loading, error } = useQuery(GetUserDocument, {
    variables: { id: '123' },
  });
  ```

---

## GraphQL のベストプラクティス

1. **スキーマの設計**:

   スキーマは、クライアントが必要とするデータを正確に反映するように設計します。

2. **型の活用**:

   TypeScriptと`graphql-codegen`を使用して、型安全なコードを実現します。

3. **クエリの最適化**:

   必要なデータのみを取得するようにクエリを設計します。

4. **エラーハンドリング**:

   クエリやミューテーションのエラーを適切に処理します。

---

これでGraphQLの基本的な使用方法と`graphql-codegen`の導入手順は完了です。