開発終わりましたが、ちょっと研究が忙しいのでPythonに戻ります。
このコードは以下の記事で作成したものです。

[Golangで簡易的認証APIを開発](https://zenn.dev/koupro0204/books/39a8bcdd1cab87)

Dockerを立ち上げてGoでサーバーを立てたら使えるようになります。
環境変数を設定して使ってください。

自分はWindowsなので
PowerShell
```
$ $env:MYSQL_USER = "root"
$ $env:MYSQL_PASSWORD = "rootpassword"
$ $env:MYSQL_HOST = "127.0.0.1"
$ $env:MYSQL_PORT = "3306"
$ $env:MYSQL_DATABASE = "auth-db"
```
## Docker立ち上げ
```
$ docker compose up
```
## APIローカル起動方法
```
$ go run ./cmd/main.go
```
