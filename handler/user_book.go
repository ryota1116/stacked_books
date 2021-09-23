package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/handler/middleware"
	"github.com/ryota1116/stacked_books/usecase"
	"net/http"
)

type UserBookHandler interface {
	RegisterUserBook(w http.ResponseWriter, r *http.Request)
}

type userBookHandler struct {
	userBookUseCase usecase.UserBookUseCase
}

func NewUserBookHandler(ubu usecase.UserBookUseCase) UserBookHandler {
	return &userBookHandler{
		userBookUseCase: ubu,
	}
}

// RegisterUserBook : booksを参照→同じのあればそれを使って、user_booksを作成
func (ubh userBookHandler) RegisterUserBook(w http.ResponseWriter, r *http.Request) {
	registerUserBookRequestParams := dto.RegisterUserBookRequestParameter{}

	err := json.NewDecoder(r.Body).Decode(&registerUserBookRequestParams)
	if err != nil {
		fmt.Println(err)
	}

	// ログイン中のユーザーを取得する
	ushm := middleware.NewUserSessionHandlerMiddleWare()
	currentUser := ushm.CurrentUser(r)

	// UserBooksレコードを作成する
	dbBook := ubh.userBookUseCase.RegisterUserBook(currentUser.Id, registerUserBookRequestParams)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbBook)
}
