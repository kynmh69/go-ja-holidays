# 祝日取得API

[![Create api image](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-api.yml/badge.svg)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/docker-publish-api.yml) [![Go Test](https://github.com/kynmh69/go-ja-holidays/actions/workflows/go.yml/badge.svg)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/go.yml) [![CodeQL](https://github.com/kynmh69/go-ja-holidays/actions/workflows/codeql.yml/badge.svg)](https://github.com/kynmh69/go-ja-holidays/actions/workflows/codeql.yml)

## 設定できる環境変数

- `PSQL_HOSTNAME`：データベースのホスト名
- `PSQL_PORT`：データベースのポート
- `PSQL_USERNAME`：データベースのユーザ名
- `PSQL_PASSWORD`：データベースのパスワード
- `LOG_DIR`：ログの出力先
- `LOG_LEVEL`：ログレベル
- `LOG_FORMAT`: ログフォーマット

## URI

### 祝日取得

全期間または指定した範囲の祝日を取得します。

| 項目                       | 説明               | 
| -------------------------- | ------------------ | 
| URI                        | /holidays          | 
| HTTPメソッド               | GET                | 
| 指定可能なクエリパラメータ | start-day, end-day | 

#### 例

##### リクエスト

```http
curl HTTP://localhost:8080/holidays?end-day=2024-12-31&start-day=2024-01-01
```

##### レスポンス

```http
[
    {
        "date": "2024-01-01T00:00:00+09:00",
        "name": "元日"
    },
    .....
    {
        "date": "2024-11-04T00:00:00+09:00",
        "name": "休日"
    },
    {
        "date": "2024-11-23T00:00:00+09:00",
        "name": "勤労感謝の日"
    }
]
```


### 祝日確認

指定した日付が祝日であるかどうかを確認します。


| 項目                       | 説明          | 
| -------------------------- | ------------- | 
| HTTPメソッド               | GET           | 
| URI                        | /holidays/day | 
| 指定可能なクエリパラメータ | なし          | 

#### 例

##### リクエスト

```http
curl http://localhost:8080/holidays/2024-03-21
```

##### レスポンス

```http
{
    "is_holiday": false,
    "date": "2024-03-21T00:00:00+09:00",
    "name": ""
}
```

### 祝日カウント

全期間または指定した期間の祝日の数をカウントします。

| 項目                       | 説明               | 
| -------------------------- | ------------------ | 
| HTTPメソッド               | GET                | 
| URI                        | /holidays/count    | 
| 指定可能なクエリパラメータ | start-day, end-day | 

#### 例

##### リクエスト

```http
curl http://localhost:8080/holidays/count?start-day=2023-01-01&end-day=2023-12-31
```
##### レスポンス

```http
{
    "count": 17
}
```
