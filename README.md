# go-ja-holidays

[![Create api image](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-api.yml/badge.svg)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-api.yml) [![Create updater image](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-updater.yml/badge.svg?branch=main)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-updater.yml) [![Create Key Manager image](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-key-manager.yml/badge.svg)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-key-manager.yml) [![codecov](https://codecov.io/gh/kynmh69/go-ja-holidays/graph/badge.svg?token=1OTK685UWI)](https://codecov.io/gh/kynmh69/go-ja-holidays)

日本の祝日を返却するAPIを提供します。

## 技術スタック

| 言語/フレームワーク/ミドルウェア | バージョン | 
|-------------------|-------| 
| Go                | 1.22  | 
| gin               | v1.10 | 
| goqu              | v9    | 
| postgresSQL       | 16    | 

## 機能

### 祝日更新

cron等で定期実行する想定で作成しています。
[国民の祝日について - 内閣府](https://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html)のCSVファイルをダウンロードして、DBを更新します。

詳細は[こちら](https://github.com/kynmh69/go-ja-holidays/blob/main/src/updater/README.md)を参照してください。

以下のように更新します。

#### 新規作成の場合

データをすべてDBに格納します。

#### すでにデータが存在する場合

差分のみを更新します。

### API

日本の祝日を取得します。
APIの制限は今のところありません。
また、認証機能もありません。

詳細は[こちら](https://github.com/kynmh69/go-ja-holidays/blob/main/src/api/README.md)を参照しください。
