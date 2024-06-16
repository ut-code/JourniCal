# JourniCal

JourniCal はカレンダーアプリとジャーナルアプリを組み合わせたアプリです。

## 0 からサーバーを起動するまで

### 環境構築: 1 回のみ

- Node.js をインストールする。
- Go >=1.22 をインストールする。
- Docker をインストールする。

### ファイルの準備: 1 回のみ

- backend/.env.sample を backend/.env にコピーする。
  - docker compose を使う場合は、コメントに従って編集する。
- credentials.json を探してきて、 backend/credentials.json にコピーする。

---
以下、 docker compose を使う場合は docker compose が代わりにやってくれます。

### 依存関係の解決

- `(cd frontend; npm ci)`
- `(cd backend; go mod download)`

### 開発環境で実行

- `backend/run-postgresql-at-localhost.sh`
- `(cd backend; go run .)`
- `(cd frontend; npm run dev)`

### ビルド

- `(cd backend; go build .)`
- `(cd frontend; npm run build)`

### ビルドされたファイルを実行

- `backend/run-postgres-at-localhost.sh`
- `cp -r ./frontend/dist ./backend/static`
- `ECHO_SERVES_FRONTEND_TOO=true backend/JourniCalBackend`

## docker-compose

```
docker compose up --build
```

で PostgreSQL サーバー・バックエンド・(フロントエンド;TODO!)がすべて起動できます。

起動時間の関係で、初回起動時は何回かサーバーが落ちてからの起動になるかもしれません。
二回目以降の起動では落ちないので大丈夫です。(大丈夫ではない) (解決方法不明: 解決できたらしてください)

## Style Guidelines as a reference

- <https://rakyll.org/style-packages/>
