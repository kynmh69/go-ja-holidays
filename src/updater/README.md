# 祝日データベース更新スクリプト

[![Create updater image](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-updater.yml/badge.svg)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-updater.yml) [![Go Test](https://github.com/kynmh69/go-ja-holidays/actions/workflows/go.yml/badge.svg)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/go.yml) [![CodeQL](https://github.com/kynmh69/go-ja-holidays/actions/workflows/codeql.yml/badge.svg)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/codeql.yml)

このディレクトリは以下のスクリプトは祝日データベース更新スクリプトです。
定期的に取得し、DBとの差分を更新します。

## データベースのマイグレーション

マイグレーションはgooseを利用します。

### データベースの起動

リポジトリ配下で実行します。

```bash
docker compose up -d database 
```

### gooseのインストール

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### マイグレーションの実行

```bash
cd src/updater/database/migrations
goose postgres "host=localhost user=app password=password dbname=holidays sslmode=disable" up
```

### マイグレーションの削除

```bash
cd src/updater/database/migrations
goose postgres "host=localhost user=app password=password dbname=holidays sslmode=disable" down
```

## 実行方法

コンテナイメージになっているので、コンテナイメージをrunします。

docker composeで実行します。

```bash
docker compose up -d
```

## 設定できる環境変数

- `PSQL_HOSTNAME`：データベースのホスト名
- `PSQL_PORT`：データベースのポート
- `PSQL_USERNAME`：データベースのユーザ名
- `PSQL_PASSWORD`：データベースのパスワード
