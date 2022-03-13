package handler

import (
	"encoding/json"
	"github.com/ryota1116/stacked_books/domain/model/UserBook"
	RegisterUserBooks "github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
	"github.com/ryota1116/stacked_books/handler/http/request/user_book/search_user_books_by_status"
	httpResponse "github.com/ryota1116/stacked_books/handler/http/response"
	"github.com/ryota1116/stacked_books/handler/http/response/user_book/register_user_book"
	"github.com/ryota1116/stacked_books/handler/middleware"
	"github.com/ryota1116/stacked_books/usecase"
	"io/ioutil"
	"net/http"
)

type UserBookHandler interface {
	RegisterUserBook(w http.ResponseWriter, r *http.Request)
	FindUserBooks(w http.ResponseWriter, r *http.Request)
	SearchUserBooksByStatus(w http.ResponseWriter, r *http.Request)
}

type userBookHandler struct {
	userBookUseCase usecase.UserBookUseCase
	userSessionHandlerMiddleWare middleware.UserSessionHandlerMiddleWareInterface
}

func NewUserBookHandler(
	ubu usecase.UserBookUseCase,
	ushmw middleware.UserSessionHandlerMiddleWareInterface) UserBookHandler {
	return &userBookHandler{
		userBookUseCase: ubu,
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
	response := register_user_book.BuildResponse(book, userBook)

	httpResponse.Response{
		StatusCode:   http.StatusOK,
		ResponseBody: response,
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

func (ubh userBookHandler) SearchUserBooksByStatus(w http.ResponseWriter, r *http.Request) {
	// ログイン中のユーザーを取得する
	currentUser := ubh.userSessionHandlerMiddleWare.CurrentUser(r)

	// リクエストボディを構造体に変換する
	requestBody := SearchUserBooksByStatus.RequestBody{}
	requestBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(requestBodyBytes, &requestBody); err != nil {
		panic(err)
	}

	// リクエストボディ構造体のバリデーションを実行
	isValid, errMsg := SearchUserBooksByStatus.FormValidator{
		SearchUserBooksByStatusRequestBody: requestBody}.Validate()
	if !isValid {
		// クライアントにHTTPレスポンスを返す
		response := httpResponse.Response{
			StatusCode:   http.StatusUnprocessableEntity,
			ResponseBody: errMsg,
		}
		response.ReturnResponse(w)
		return
	}

	// 読書ステータスのオブジェクトを生成
	status := UserBook.Status{Value: requestBody.Status}

	// 書籍ステータスから本を検索する
	searchUserBooksByStatusResponse := ubh.userBookUseCase.
		SearchUserBooksByStatus(currentUser.Id, status)

	// 正常なレスポンス
	response := httpResponse.Response{
		StatusCode:   http.StatusOK,
		ResponseBody: searchUserBooksByStatusResponse,
	}
	response.ReturnResponse(w)
}
