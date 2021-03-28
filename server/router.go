package server

import (
	"github.com/gorilla/mux"

	"github.com/ryota1116/stacked_books/handler"
	"github.com/ryota1116/stacked_books/infra/persistence"
	"github.com/ryota1116/stacked_books/usecase"
)

// webサーバーに接続する
func HandleFunc() mux.Router {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	router := mux.NewRouter().StrictSlash(true)
	// エンドポイント(リクエストを処理して、レスポンスを返す)
	router.HandleFunc("/signup", userHandler.SignUp).Methods("POST")
	router.HandleFunc("/signin", userHandler.SignIn).Methods("POST")
	router.HandleFunc("/user/{userId:[0-9]+}", userHandler.ShowUser).Methods("GET")
	router.HandleFunc("/user/authenticate", handler.VerifyToken).Methods("POST")

	return *router
}
