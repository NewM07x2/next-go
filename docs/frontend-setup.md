# フロントエンド構築手順

このドキュメントでは、Next.jsを使用したフロントエンドの構築手順を説明します。

---

## 必要なツール

以下のツールがインストールされていることを確認してください：

- [Node.js](https://nodejs.org/)（バージョン 20.17.0 以上推奨）
- npm または yarn

---

## 環境セットアップ

1. **Next.js アプリケーションの作成**

   `next-app` フォルダに移動し、以下のコマンドを実行して TypeScript 対応の Next.js アプリケーションを作成します。

   ```bash
   npx create-next-app@latest . --typescript
   ```

2. **必要なパッケージのインストール**

   GraphQL と urql を使用するために、以下のコマンドを実行して依存関係をインストールします。

   ```bash
   npm install graphql urql
   ```

3. **GraphQL Code Generator の導入**

   GraphQL クエリやスキーマを管理するために、`graphql-codegen` を導入します。

   - 必要なパッケージをインストールします。

     ```bash
     npm install -D @graphql-codegen/cli
     ```

   - 初期設定を行います。

     ```bash
     npx graphql-codegen init
     ```

   - プロンプトに従って以下を設定します：
     - **What type of application are you building?**: `Application built with React`
     - **Where is your schema?**: `src/schema.graphql`
     - **Where are your operations and fragments?**: `src/**/*.tsx`
     - **Where to write the output**: `src/graphql/generated/`
     - **Do you want to generate an introspection file?**: `Yes`
     - **How to name the config file?**: `codegen.ts`
     - **What script in package.json should run the codegen?**: `codegen`

   - 設定後、以下のコマンドでコード生成を実行します。

     ```bash
     npm run codegen
     ```

---

## サーバーの起動と動作確認

1. **サーバーの起動**

   以下のコマンドを実行して開発サーバーを起動します。

   ```bash
   npm run dev
   ```

2. **ブラウザで確認**

   ブラウザを開き、以下の URL にアクセスしてください。

   ```
   http://localhost:3000
   ```

   アプリケーションが正常に動作していることを確認してください。

---

## トラブルシューティング

- **依存関係のエラー**:

  - `node_modules` を削除して再インストールします。

    ```bash
    rm -rf node_modules
    npm install
    ```

- **ポート競合エラー**:

  - 他のプロセスがポート `3000` を使用している場合、`package.json` 内のスクリプトを修正して別のポートを指定してください。

    ```bash
    npm run dev -- -p 4000
    ```

---

これでフロントエンドの構築手順は完了です。