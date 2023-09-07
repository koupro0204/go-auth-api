開発終わるまではざっくりとメモ程度に残しておきます。
環境変数を設定
自分はWindowsなので
PowerShell
```
$ $env:MYSQL_USER = "root"
$ $env:MYSQL_PASSWORD = "rootpassword"
$ $env:MYSQL_HOST = "127.0.0.1"
$ $env:MYSQL_PORT = "3306"
$ $env:MYSQL_DATABASE = "auth-db"
```