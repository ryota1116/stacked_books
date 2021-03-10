package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"../handler"
	"../infra/persistence"
	"../usecase"
)

// webサーバーに接続する
func StartWebServer() error {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	router := mux.NewRouter().StrictSlash(true)
	// エンドポイント(リクエストを処理して、レスポンスを返す)
	router.HandleFunc("/signup", userHandler.SignUp).Methods("POST")
	router.HandleFunc("/signin", userHandler.SignIn).Methods("POST")
	router.HandleFunc("/user/{userId:[0-9]+}", userHandler.ShowUser).Methods("GET")
	router.HandleFunc("/user/authenticate", handler.VerifyToken).Methods("POST")

	log.Println("サーバー起動 : 8080 port で受信")
	//log.Fatal(fmt.Sprintf(":%d", 8080), router)
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
