package main

import (
	"github.com/ryota1116/stacked_books/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	server.StartWebServer()

}
