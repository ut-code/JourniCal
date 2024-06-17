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
- `cd backend; ECHO_SERVES_FRONTEND_TOO=true ./backend`

## docker-compose

```
docker compose up --build
```

で PostgreSQL サーバー・バックエンド・(フロントエンド;TODO!)がすべて起動できます。

## 本番環境

本番環境で実行するには、以下のことをしてください。

### 事前準備

- backend/credentials.json を用意する
- 環境変数 DSN を設定する

### ビルド

```sh
docker build -f Dockerfile.prod -t journical-full .
```

### 実行

```sh
docker run \
  -e DSN \ # inherit DSN from its env
  -p ${PORT:-3000}:3000 \ # run at $PORT, default to 3000 if $PORT is not set
  journical-full
```

## Guidelines

### Style Guidelines

- <https://google.github.io/styleguide/go/>
- <https://google.github.io/styleguide/go/decisions>
- <https://rakyll.org/style-packages/>

### Project Layout Standard(s)

- <https://github.com/golang-standards/project-layout/blob/master/README_ja.md>
