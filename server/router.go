package server

import (
	"github.com/gorilla/mux"
	book2 "github.com/ryota1116/stacked_books/infra/datasource/book"
	user2 "github.com/ryota1116/stacked_books/infra/datasource/user"
	userbook2 "github.com/ryota1116/stacked_books/infra/datasource/userbook"
	"github.com/ryota1116/stacked_books/infra/externalapi/google-books-api"
	book3 "github.com/ryota1116/stacked_books/interfaces/api/handler/http/request/book"
	user3 "github.com/ryota1116/stacked_books/interfaces/api/handler/http/request/user"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/http/request/user_book"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/middleware"
	"github.com/ryota1116/stacked_books/usecase/book"
	"github.com/ryota1116/stacked_books/usecase/user"
	"github.com/ryota1116/stacked_books/usecase/userbook"
)

// webサーバーに接続する
func HandleFunc() mux.Router {
	userPersistence := user2.NewUserPersistence()
	userUseCase := user.NewUserUseCase(userPersistence)
	userHandler := user3.NewUserHandler(userUseCase)

	bookPersistence := book2.NewBookPersistence()
	userBookPersistence := userbook2.NewUserBookPersistence()
	userBookUseCase := userbook.NewUserBookUseCase(bookPersistence, userBookPersistence)
	userSessionHandlerMiddleWare := middleware.NewUserSessionHandlerMiddleWare()
	userBookHandler := user_book.NewUserBookHandler(userBookUseCase, userSessionHandlerMiddleWare)

	bookUseCase := book.NewBookUseCase(
		bookPersistence,
		google_books_api.NewClient(),
	)
	bookHandler := book3.NewBookHandler(bookUseCase)

	router := mux.NewRouter().StrictSlash(true)
	// エンドポイント(リクエストを処理して、レスポンスを返す)
	router.HandleFunc("/signup", userHandler.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/signin", userHandler.SignIn).Methods("POST", "OPTIONS")

	// 外部APIを用いた書籍検索のエンドポイント
	router.HandleFunc("/books/search", bookHandler.SearchBooks).Methods("GET", "OPTIONS")

	router.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET", "OPTIONS")

	// ユーザーと書籍を紐付ける
	router.HandleFunc("/register/userbook", userBookHandler.RegisterUserBook).Methods("POST")
	// ログイン中のユーザーが登録している本の一覧を取得する
	router.HandleFunc("/user/books", userBookHandler.FindUserBooks).Methods("GET")

	router.HandleFunc("/user/books/status", userBookHandler.SearchUserBooksByStatus).Methods("GET")
	return *router
}
