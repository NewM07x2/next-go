# 概要

このアプリケーションは、タスクを管理するためのシステム(next-go-task)です。  
ユーザーは新規登録、ログイン、タスクの登録・閲覧・編集・削除が可能です。

---

## 開発環境

バックエンド: Go  
フロントエンド: Next.js  
データベース: PostgreSQL  
コンテナ管理: Docker  

---

## 環境構築

以下の手順で環境を構築してください。

### 1. リポジトリのクローン

```bash
git clone <リポジトリのURL>
cd next-go-task
```

### 2. Dockerコンテナの起動

```bash
docker-compose up --build -d
```

### 3. サーバーサイドの初期設定

```bash
docker-compose exec backend go mod init next-go-task
docker-compose exec backend go mod tidy
```

### 4. フロントエンドのセットアップ

```bash
docker-compose exec react-app npm install
docker-compose exec react-app npm start
exit
```

### 5. データベースの初期化

PostgreSQLの初期化スクリプトを実行します。

```bash
docker-compose exec db psql -U postgres -d task_db -f /docker-entrypoint-initdb.d/init.sql
```

### 6. キャッシュのクリア

以下はコンテナの停止や Docker のキャッシュ（イメージ／コンテナ／ボリューム）をクリアする際に便利なコマンド例です。破壊的な操作が含まれるため、実行前に必ず影響範囲を確認してください。

#### コンテナの停止

```powershell
# docker-compose で立ち上げたサービス群を停止（コンテナは停止するが残る）
docker-compose stop

# docker-compose で停止・削除（ネットワークは削除されるがボリュームは残る）
docker-compose down
```

```powershell
# 全ての実行中コンテナを停止（PowerShell）
docker stop (docker ps -q)
```

```bash
# 全ての実行中コンテナを停止（bash/zsh）
docker stop $(docker ps -q)
```

#### 完全に削除（クリーン）

```powershell
# コンテナ・ネットワークに加え、関連ボリュームも削除（永続データが失われます）
docker-compose down --volumes
# 短縮形: docker-compose down -v

# Composeで作成されたイメージも削除する場合
docker-compose down --rmi all

# ボリュームとイメージの両方を削除する完全クリーン
docker-compose down --rmi all --volumes
```

```powershell
# Docker が使う未使用リソース（イメージ・コンテナ・ネットワーク等）を一括削除
# ※ デフォルトではボリュームは削除されません。ボリュームも含める場合は --volumes を追加します。
docker system prune -a --volumes
```

- 注意:
  - `--volumes` を付けると Compose の named volumes や自動作成されたボリュームが削除され、データは復元できません。
  - ホストの bind マウント（ホストパスをマウントしている場合）は削除されません。
  - 実行前に `docker ps -a`, `docker images`, `docker volume ls` などで影響範囲を確認してください。

#### アプリケーションのキャッシュクリア（コンテナ内部で実行）

- Next.js / npm のキャッシュをクリアする例:

```bash
# react-app コンテナ内で npm キャッシュを強制クリア
docker-compose exec react-app npm cache clean --force
```

- Go のビルドキャッシュをクリアする例:

```bash
# backend コンテナ内で Go ビルドキャッシュをクリア
docker-compose exec backend go clean -cache
```

- PostgreSQL のデータを初期化したい場合は、初期化スクリプトを再実行するか、該当ボリュームを削除して再作成してください（*データが失われます*）。

---

これにより、古い設定やキャッシュが原因で発生する問題を防ぐことができます。

---

## ドキュメント

各構成に関する詳細な情報は以下のドキュメントを参照してください。

---

## 注意事項

- `.env`ファイルの設定が必要です。詳細は各ドキュメントを参照してください。
- Dockerが正しくインストールされていることを確認してください。

---

## コマンド チートシート

起動

```bash
# docker-composeで起動（バックグラウンド、必要に応じてビルド）
docker-compose up --build -d

# 単一イメージから起動の例
docker run -d --name myapp -p 3000:3000 <イメージ名>
```

停止・削除

```bash
# 実行中のコンテナを停止
docker stop <コンテナID|名前>

# docker-composeで停止・削除（ネットワークを削除、ボリュームは残る）
docker-compose down

# ボリュームも削除（データ消失に注意）
docker-compose down --volumes
```

一覧表示

```bash
# 実行中のコンテナ一覧
docker ps

# 停止済みも含む全コンテナ一覧
docker ps -a

# ローカルイメージ一覧
docker images

# ボリューム一覧
docker volume ls
```

削除 / クリーン

```bash
# 停止済みコンテナを削除
docker rm <コンテナID>

# イメージを削除
docker rmi <イメージID>

# 未使用ボリュームをまとめて削除
docker volume prune

# 全体クリーン（ボリューム含める場合は --volumes）
docker system prune -a --volumes
```

補助（ログ、シェル等）

```bash
# ログを追う
docker logs -f <コンテナ>

# コンテナ内でシェルを開く
docker exec -it <コンテナ> /bin/sh

# ディスク使用状況の確認
docker system df
```
---
