package server

import (
	"github.com/gorilla/mux"
	"github.com/ryota1116/stacked_books/domain/model/googleBooksApi"
	"github.com/ryota1116/stacked_books/handler"
	"github.com/ryota1116/stacked_books/handler/middleware"
	"github.com/ryota1116/stacked_books/infra/persistence"
	"github.com/ryota1116/stacked_books/usecase"
)

// webサーバーに接続する
func HandleFunc() mux.Router {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	bookPersistence := persistence.NewBookPersistence()
	userBookPersistence := persistence.NewUserBookPersistence()
	userBookUseCase := usecase.NewUserBookUseCase(bookPersistence, userBookPersistence)
	userSessionHandlerMiddleWare := middleware.NewUserSessionHandlerMiddleWare()
	userBookHandler := handler.NewUserBookHandler(userBookUseCase, userSessionHandlerMiddleWare)

	bookUseCase := usecase.NewBookUseCase(googleBooksApi.NewClient())
	bookHandler := handler.NewBookHandler(bookUseCase)

	router := mux.NewRouter().StrictSlash(true)
	// エンドポイント(リクエストを処理して、レスポンスを返す)
	router.HandleFunc("/signup", userHandler.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/signin", userHandler.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/{userId:[0-9]+}", userHandler.ShowUser).Methods("GET")

	// 外部APIを用いた書籍検索のエンドポイント
	router.HandleFunc("/books/search", bookHandler.SearchBooks).Methods("GET", "OPTIONS")

	// ユーザーと書籍を紐付ける
	router.HandleFunc("/register/book", userBookHandler.RegisterUserBook).Methods("POST")
	// ログイン中のユーザーが登録している本の一覧を取得する
	router.HandleFunc("/user/books", userBookHandler.FindUserBooks).Methods("GET")

	return *router
}
