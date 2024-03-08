# 祝日データベース更新スクリプト

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
goose mysql "root:secret@tcp(localhost:3306)/holidays?parseTime=true" up
```

## 実行方法

コンテナイメージになっているので、コンテナイメージをrunします。

docker composeで実行します。

```bash
docker compose up -d
```
