package handler

import (
	"encoding/json"
	RegisterUserBooks "github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
	httpResponse "github.com/ryota1116/stacked_books/handler/http/response"
	"github.com/ryota1116/stacked_books/handler/http/response/user_book"
	"github.com/ryota1116/stacked_books/handler/middleware"
	"github.com/ryota1116/stacked_books/usecase"
	"net/http"
)

type UserBookHandler interface {
	RegisterUserBook(w http.ResponseWriter, r *http.Request)
	FindUserBooks(w http.ResponseWriter, r *http.Request)
}

type userBookHandler struct {
	userBookUseCase              usecase.UserBookUseCase
	userSessionHandlerMiddleWare middleware.UserSessionHandlerMiddleWareInterface
}

func NewUserBookHandler(
	ubu usecase.UserBookUseCase,
	ushmw middleware.UserSessionHandlerMiddleWareInterface) UserBookHandler {
	return &userBookHandler{
		userBookUseCase:              ubu,
		userSessionHandlerMiddleWare: ushmw,
	}
}

// RegisterUserBook : booksを参照→同じのあればそれを使って、user_booksを作成
func (ubh userBookHandler) RegisterUserBook(w http.ResponseWriter, r *http.Request) {
	// JSONのリクエストボディを構造体に変換する
	requestBody := RegisterUserBooks.RequestBody{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		httpResponse.Response{
			StatusCode:   http.StatusInternalServerError,
			ResponseBody: err.Error(),
		}.ReturnResponse(w)
		return
	}

	// リクエストボディ構造体のバリデーションを実行
	isValid, errMsg := RegisterUserBooks.FormValidator{
		RequestBody: requestBody,
	}.Validate()
	if !isValid {
		httpResponse.Response{
			StatusCode:   http.StatusUnprocessableEntity,
			ResponseBody: errMsg,
		}.ReturnResponse(w)
		return
	}

	// ログイン中のユーザーを取得する
	currentUser := ubh.userSessionHandlerMiddleWare.CurrentUser(r)

	// UserBooksレコードを作成する
	book, userBook := ubh.userBookUseCase.RegisterUserBook(
		currentUser.Id,
		requestBody)

	// 書籍登録用のレスポンス構造体を生成する

	httpResponse.Response{
		StatusCode: http.StatusOK,
		ResponseBody: user_book.RegisterUserBookResponseGenerator{
			Book:     book,
			UserBook: userBook,
		}.Execute(),
	}.ReturnResponse(w)
}

// FindUserBooks : ログイン中のユーザーが登録している本の一覧を取得する
func (ubh userBookHandler) FindUserBooks(w http.ResponseWriter, r *http.Request) {
	// セッション情報からUserを取得
	ushm := middleware.NewUserSessionHandlerMiddleWare()
	user := ushm.CurrentUser(r)

	userBooks, err := ubh.userBookUseCase.FindUserBooksByUserId(user.Id)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	// 正常なレスポンス
	response := httpResponse.Response{
		StatusCode:   http.StatusOK,
		ResponseBody: userBooks,
	}
	response.ReturnResponse(w)
}
