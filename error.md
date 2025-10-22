(base) mnitta@MasatonoMacBook-Pro next-go % docker-compose up --build -d               
WARN[0000] /Users/mnitta/devlop/next-go/docker-compose.yml: the attribute `version` is obsolete, it will be ignored, please remove it to avoid potential confusion 
Compose can now delegate builds to bake for better performance.
 To do so, set COMPOSE_BAKE=true.
[+] Building 1.2s (22/25)                                                   docker:desktop-linux
 => [go-app internal] load build definition from Dockerfile.backend                         0.0s
 => => transferring dockerfile: 526B                                                        0.0s
 => [next-app internal] load build definition from Dockerfile.frontend                      0.0s
 => => transferring dockerfile: 226B                                                        0.0s
 => [go-app internal] load metadata for docker.io/library/golang:1.23-alpine                1.0s
 => [go-app internal] load metadata for docker.io/library/alpine:latest                     1.0s
 => [next-app internal] load metadata for docker.io/library/node:20                         1.0s
 => [go-app internal] load .dockerignore                                                    0.0s
 => => transferring context: 335B                                                           0.0s
 => [go-app builder 1/6] FROM docker.io/library/golang:1.23-alpine@sha256:383395b794dffa5b  0.0s
 => [go-app internal] load build context                                                    0.0s
 => => transferring context: 9.01kB                                                         0.0s
 => [go-app stage-1 1/4] FROM docker.io/library/alpine:latest@sha256:4b7ce07002c69e8f3d704  0.0s
 => CACHED [go-app stage-1 2/4] RUN apk --no-cache add ca-certificates                      0.0s
 => CACHED [go-app stage-1 3/4] WORKDIR /root/                                              0.0s
 => CACHED [go-app builder 2/6] WORKDIR /app                                                0.0s
 => CACHED [go-app builder 3/6] COPY go.mod go.sum ./                                       0.0s
 => ERROR [go-app builder 4/6] RUN go mod download                                          0.2s
 => [next-app internal] load .dockerignore                                                  0.0s
 => => transferring context: 2B                                                             0.0s
 => [next-app 1/4] FROM docker.io/library/node:20@sha256:cb61978c7e08f58f6042ae65dd2198183  0.0s
 => [next-app internal] load build context                                                  0.0s
 => => transferring context: 11.39kB                                                        0.0s
 => CACHED [next-app 2/4] WORKDIR /app                                                      0.0s
 => CACHED [next-app 3/4] COPY package*.json ./                                             0.0s
 => CACHED [next-app 4/4] COPY . .                                                          0.0s
 => [next-app] exporting to image                                                           0.0s
 => => exporting layers                                                                     0.0s
 => => writing image sha256:014e9ae161e107fa0915d33b62ea4e872de4b663f18590063fe0450b5b381b  0.0s
 => => naming to docker.io/library/next-go-next-app                                         0.0s
 => [next-app] resolving provenance for metadata file                                       0.0s
------
 > [go-app builder 4/6] RUN go mod download:
0.173 go: go.mod requires go >= 1.24.0 (running go 1.23.12; GOTOOLCHAIN=local)
------
failed to solve: process "/bin/sh -c go mod download" did not complete successfully: exit code: 1
(base) mnitta@MasatonoMacBook-Pro next-go % 

---

## 🔍 エラー原因

**Goバージョンの不一致**

- `go.mod`ファイルで`go 1.24.0`を要求していた
- Dockerイメージ(`golang:1.23-alpine`)にインストールされているGoのバージョンは`1.23.12`
- Go 1.24.0はまだリリースされていない（2024年時点で最新は1.23系）

## ✅ 解決方法

### 方法1: go.modのバージョンを修正（推奨） ✓ **完了**

`go-app/go.mod`を以下のように修正しました:

```diff
- go 1.24.0
+ go 1.23.0
```

その後、依存関係を更新:

```bash
cd go-app
go mod tidy
```

### 方法2: Dockerイメージのバージョンを変更（非推奨）

`docker/Dockerfile.backend`のベースイメージを変更:

```diff
- FROM golang:1.23-alpine AS builder
+ FROM golang:1.24-alpine AS builder  # Go 1.24がリリースされた場合のみ
```

**注意**: Go 1.24.0は現時点で存在しないため、この方法は使用できません。

### 方法3: GOTOOLCHAINを設定（一時的な対処）

環境変数で自動ダウンロードを許可:

```bash
export GOTOOLCHAIN=auto
```

ただし、Dockerビルド時には推奨されません。

## 🧪 検証手順

修正後、以下のコマンドでDockerビルドを確認:

```bash
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

正常に起動することを確認:

```bash
docker-compose ps
curl http://localhost:8080/api/health
```

## 📝 追加の警告について

`docker-compose.yml`の以下の警告も修正することをお勧めします:

```
WARN[0000] the attribute `version` is obsolete
```

**修正方法**: `docker-compose.yml`の最初の行を削除

```diff
- version: "3.8"
  services:
    ...
```

Docker Compose v2では`version`属性は不要になりました。

## 🚀 トラブルシューティング

もしまだエラーが発生する場合:

1. **Dockerキャッシュのクリア**:
   ```bash
   docker system prune -a
   ```

2. **go.sumの削除と再生成**:
   ```bash
   cd go-app
   rm go.sum
   go mod tidy
   ```

3. **Goのバージョン確認**:
   ```bash
   go version
   ``` 