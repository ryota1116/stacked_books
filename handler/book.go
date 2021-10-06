package handler

import (
	"encoding/json"
	"errors"
	"github.com/ryota1116/stacked_books/handler/http/request"
	"github.com/ryota1116/stacked_books/handler/http/request/book"
	httpResponse "github.com/ryota1116/stacked_books/handler/http/response"
	"github.com/ryota1116/stacked_books/handler/http/response/book"
	"github.com/ryota1116/stacked_books/usecase"
	"io/ioutil"
	"net/http"
)

type BookHandlerInterface interface {
	SearchBooks(w http.ResponseWriter, r *http.Request)
}

type bookHandler struct {
	bookUseCase usecase.BookUseCaseInterface
}

func NewBookHandler(bu usecase.BookUseCaseInterface) BookHandlerInterface {
	return &bookHandler{
		bookUseCase: bu,
	}
}

// SearchBooks : 外部APIを用いた書籍検索のエンドポイント
func (bh bookHandler) SearchBooks(w http.ResponseWriter, r *http.Request)  {
	var requestBody searchBooksRequest.RequestBody
	requestBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(requestBodyBytes, &requestBody); err != nil {
		panic(err)
	}

	// リクエストボディのバリデーション
	isValid, errMsg := request.BookHandlerFormValidator{
		GoogleBooksApiRequestBody: requestBody}.Validate()
	if !isValid {
		// クライアントにHTTPレスポンスを返す
		response := httpResponse.Response{
			StatusCode:   http.StatusUnprocessableEntity,
			ResponseBody: errMsg,
		}
		response.ReturnResponse(w)
		return
	}

	// 外部APIで書籍を検索
	responseFromGoogleBooksAPI, err := bh.bookUseCase.SearchBooks(requestBody)
	// 外部APIリクエストでエラーが発生した場合
	if err != nil {
		httpResponse.Return500Response(w, errors.New("検索に失敗しました"))
		return
	}

	// GoogleBooksAPIのJSONレスポンスの構造体から、 書籍検索用のHTTPレスポンスボディ構造体を生成する
	searchBooksResponse := book.SearchBooksResponseGenerator{
		ResponseBodyFromGoogleBooksAPI: responseFromGoogleBooksAPI,
	}.Execute()

	// 正常なレスポンス
	response := httpResponse.Response{
		StatusCode:   http.StatusOK,
		ResponseBody: searchBooksResponse,
	}
	response.ReturnResponse(w)
}
