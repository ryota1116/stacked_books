package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ryota1116/stacked_books/server"
)

func main() {
	router := server.HandleFunc()
	server.StartWebServer(&router)
}
