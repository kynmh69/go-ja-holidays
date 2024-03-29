# API Key管理

## TailwindCSSのコンパイル

以下の手順は`key_management`配下で実施すること。

1. TailwindCSSに関連するソースをインストールする

```bash
npm install -D tailwindcss postcss autoprefixer
```

2. コンパイルをする

```bash
npx tailwindcss -i ./static/css/main.css -o ./static/css/main_output.css --watch
```
