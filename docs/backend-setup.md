# バックエンド構築手順 (Go)

このドキュメントでは、Goを使用したバックエンドの構築手順を説明します。

---

## 必要なツール

以下のツールがインストールされていることを確認してください：

- [Go](https://go.dev/dl/)（バージョン 1.20 以上推奨）
- ターミナル（PowerShell, bash など）

---

## 環境セットアップ

1. **リポジトリのクローン**

   プロジェクトをクローンし、`go-app`ディレクトリに移動します。

   ```bash
   git clone <リポジトリのURL>
   cd next-go/go-app
   ```

2. **Go モジュールの初期化**

   プロジェクトがすでに初期化されている場合はスキップしてください。

   ```bash
   go mod init next-go-sample
   ```

3. **依存関係のインストール**

   必要なパッケージをインストールします。

   ```bash
   go get github.com/labstack/echo/v4
   go mod tidy
   ```

---

## サーバーの起動と動作確認

1. **サーバーの起動**

   以下のコマンドを実行してサーバーを起動します。

   ```bash
   go run src/main.go
   ```

   起動時に以下のようなメッセージが表示されます：

   ```
   ⇨ http server started on [::]:8080
   ```

2. **動作確認**

   ブラウザまたはAPIクライアント（Postman, curl など）を使用して、以下のエンドポイントにアクセスしてください。

   - **URL**: `http://localhost:8080/health`
   - **HTTP メソッド**: GET

   **期待されるレスポンス**:

   ```text
   API is running
   ```

---

## 環境セットアップ（Docker コンテナ）

1. **Dockerfile と docker-compose.yml の確認**

   - `docker/Dockerfile.backend` を使用して Go アプリケーションをビルドします。
   - `docker-compose.yml` を使用してコンテナを管理します。

2. **コンテナのビルドと起動**

   以下のコマンドを実行してコンテナをビルドし、起動します：

   ```bash
   docker-compose up --build
   ```

3. **動作確認**

   ブラウザまたはAPIクライアント（Postman, curl など）を使用して、以下のエンドポイントにアクセスしてください。

   - **URL**: `http://localhost:8080/health`
   - **HTTP メソッド**: GET

   **期待されるレスポンス**:

   ```text
   API is running
   ```

---

## トラブルシューティング

- **ポート競合エラー**:

  - 他のプロセスがポート `8080` を使用している場合、`docker-compose.yml` 内の `ports` セクションを変更してください（例: `- "3000:8080"`）。

- **依存関係のエラー**:

  - `go.mod` や `go.sum` が正しく設定されているか確認してください。
  - 必要に応じて以下を実行してください：

    ```bash
    go mod tidy
    ```

---

これでバックエンドの構築手順は完了です。