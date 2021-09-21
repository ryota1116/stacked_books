package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/handler/middleware"
	"github.com/ryota1116/stacked_books/usecase"
	"net/http"
)

type UserBookHandler interface {
	RegisterUserBook(w http.ResponseWriter, r *http.Request)
	ReadUserBooks(w http.ResponseWriter, r *http.Request)
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
	//
	bookParams := model.UserBookParameter{}
	err := json.NewDecoder(r.Body).Decode(&bookParams)
	if err != nil {
		fmt.Println(err)
	}

	currentUser := middleware.CurrentUser(r)

	dbBook := ubh.userBookUseCase.RegisterUserBook(currentUser.Id,bookParams)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbBook)
}

func (ubh userBookHandler) ReadUserBooks(w http.ResponseWriter, r *http.Request) {
	// セッション情報からUserを取得
	user := middleware.CurrentUser(r)
	userBooks := ubh.userBookUseCase.ReadUserBooks(user.Id)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(userBooks)
	if err != nil {
		return 
	}
}
