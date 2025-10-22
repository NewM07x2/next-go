.PHONY: help build run test clean docker-up docker-down

help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Goアプリケーションをビルド
	cd go-app && go build -o bin/main cmd/api/main.go

run: ## Goアプリケーションを実行
	cd go-app && go run cmd/api/main.go

dev: ## ホットリロードで開発サーバーを起動 (Air必要)
	cd go-app && air

test: ## テストを実行
	cd go-app && go test -v ./...

test-cover: ## カバレッジ付きでテストを実行
	cd go-app && go test -cover ./...

clean: ## ビルド成果物を削除
	cd go-app && rm -rf bin/ tmp/

docker-up: ## Dockerコンテナを起動
	docker-compose up --build

docker-down: ## Dockerコンテナを停止・削除
	docker-compose down

docker-logs: ## Dockerログを表示
	docker-compose logs -f

install-air: ## Air (ホットリロードツール) をインストール
	go install github.com/cosmtrek/air@latest

tidy: ## go mod tidyを実行
	cd go-app && go mod tidy
