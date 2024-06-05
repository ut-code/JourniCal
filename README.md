# JourniCal

JourniCal はカレンダーアプリとジャーナルアプリを組み合わせたアプリです。

## docker-compose

```
docker-compose up --build
```

で PostgreSQL サーバー・バックエンド・(フロントエンド;TODO!)がすべて起動できます。

## frontend

以下のコマンドを実行する際には frontend ディレクトリに移動してください。

### 環境構築

```
npm ci
```

### 開発用サーバーの起動

```
npm run dev
```

### ビルド

```
npm run build
```

## backend

backend ディレクトリに移動し、以下のコマンドを実行してください。

### パッケージのダウンロード

```sh
go mod download
```

### サーバーの起動

```sh
go run .
```

### ビルド

```sh
go build .
```

実行可能バイナリが `./JourniCalBackend` という名前で生成されます。
