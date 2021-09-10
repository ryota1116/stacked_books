package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartWebServer(router *mux.Router) error {
	log.Println("サーバー起動 : 3000 port で受信")
	return http.ListenAndServe(fmt.Sprintf(":%d", 3000), router)
	//srv := http.Server{
	//	Addr: fmt.Sprintf("127.0.0.1:8080"),
	//	//Addr: fmt.Sprintf("%s:%d", os.Getenv("HOSTNAME"), os.Getenv("PORT")),
	//}

	//return http.ListenAndServe(fmt.Sprintf("127.0.0.1:8080", nil), router)
}