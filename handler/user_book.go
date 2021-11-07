package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/UserBook"
	"github.com/ryota1116/stacked_books/handler/http/request/user_book/search_user_books_by_status"
	httpResponse "github.com/ryota1116/stacked_books/handler/http/response"
	"github.com/ryota1116/stacked_books/usecase"
	"io/ioutil"
	"net/http"
)

type UserBookHandler interface {
	RegisterUserBook(w http.ResponseWriter, r *http.Request)
	ReadUserBooks(w http.ResponseWriter, r *http.Request)
	SearchUserBooksByStatus(w http.ResponseWriter, r *http.Request)
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

	//認証
	//if VerifyToken(w, r) {
	//}

	dbBook := ubh.userBookUseCase.RegisterUserBook(bookParams)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbBook)
}

func (ubh userBookHandler) ReadUserBooks(w http.ResponseWriter, r *http.Request) {
	// セッション情報からUserを取得
	user := CurrentUser(r)
	userBooks := ubh.userBookUseCase.ReadUserBooks(user.Id)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(userBooks)
	if err != nil {
		return 
	}
}

func (ubh userBookHandler) SearchUserBooksByStatus(w http.ResponseWriter, r *http.Request) {
	// セッション情報からUserを取得
	user := CurrentUser(r)

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
		SearchUserBooksByStatus(user.Id, status)

	// 正常なレスポンス
	response := httpResponse.Response{
		StatusCode:   http.StatusOK,
		ResponseBody: searchUserBooksByStatusResponse,
	}
	response.ReturnResponse(w)
}
