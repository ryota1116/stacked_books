package main

import (
	"../domain/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Hello World\n")

}

func SignUpHandler(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("Hello World\n")
}

func main()  {
	model.StartWebServer()

}

