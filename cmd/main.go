package main

import (
	"../server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	server.StartWebServer()

}
