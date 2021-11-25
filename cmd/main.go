package main

import (
	_ "github.com/go-sql-driver/mysql"
	config "github.com/ryota1116/stacked_books"
	"github.com/ryota1116/stacked_books/infra"
	"github.com/ryota1116/stacked_books/server"
)

func main() {
	// TODO: main.goでserver, router, dbの関数呼びたい
	router := server.HandleFunc()
	server.StartWebServer(&router)

	// ローカルDBの環境変数を設定する
	config.SetLocalDBConfig()
	// DBサーバーに接続する
	infra.DbConnect()
}
