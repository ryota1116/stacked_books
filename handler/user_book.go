package handler

import (
	"encoding/json"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	RegisterUserBooks "github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
	httpResponse "github.com/ryota1116/stacked_books/handler/http/response"
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
	// JSONのリクエストボディを構造体に変換する
	registerUserBookRequestParams := dto.RegisterUserBookRequestParameter{}

	err := json.NewDecoder(r.Body).Decode(&registerUserBookRequestParams)
	if err != nil {
		httpResponse.Response{
			StatusCode:   http.StatusInternalServerError,
			ResponseBody: err.Error(),
		}.ReturnResponse(w)
		return
	}

	// リクエストボディ構造体のバリデーションを実行
	isValid, errMsg := RegisterUserBooks.FormValidator{
		RegisterUserBookRequestParameter: registerUserBookRequestParams}.
		Validate()
	if !isValid {
		httpResponse.Response{
			StatusCode:   http.StatusUnprocessableEntity,
			ResponseBody: errMsg,
		}.ReturnResponse(w)
		return
	}

	// ログイン中のユーザーを取得する
	ushm := middleware.NewUserSessionHandlerMiddleWare()
	currentUser := ushm.CurrentUser(r)

	// UserBooksレコードを作成する
	book, userBook, err := ubh.userBookUseCase.RegisterUserBook(
		currentUser.Id,
		registerUserBookRequestParams);
	// レコード作成失敗時のレスポンス
	if err != nil {
		httpResponse.Response{
			StatusCode:   http.StatusBadRequest,
			ResponseBody: err.Error(),
		}.ReturnResponse(w)
		return
	}

	// RegisterUserBookResponse構造体を生成する
	registerUserBookResponse := dto.BuildRegisterUserBookResponse(book, userBook)

	httpResponse.Response{
		StatusCode:   http.StatusOK,
		ResponseBody: registerUserBookResponse,
	}.ReturnResponse(w)
}
