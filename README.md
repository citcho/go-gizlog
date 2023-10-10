# go-gizlog

## 環境構築

### ①Docker
```bash
# ローカル用コンテナイメージ作成
$ make build-local

# コンテナ起動
$ make up
```
### ②DBマイグレーション
```bash
# Goコンテナに接続
$ docker compose exec gizlog bash

# 初回だけ（マイグレーションテーブルの作成）
$ go run /app/cmd/migrate/main.go db init

# マイグレーション
$ make migrate

# Seederはまだ無い🙇‍♂️
```

## テスト関連

### テスト実行コマンド
```bash
$ make test
```
### テストファイル雛形作成コマンド
```bash
# example
# ※テストファイルを作成したいディレクトリに移動する必要があります。
$ gotests -w -all greeter.go
```

### モックファイル作成コマンド
```bash
# モックが必要なインターフェースには、以下のように定義の上にgenerate文が記述します。
# //go:generate mockgen -source=./usecase.go -destination=./mock/usecase.go
make generate
```

## デバッグ関連
基本的には`log`パッケージでデバッグします。
