package book

import (
	"errors"
	search_books2 "github.com/ryota1116/stacked_books/interfaces/api/handler/http/request/book/search_books"
	httpResponse "github.com/ryota1116/stacked_books/interfaces/api/handler/http/response"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/http/response/book"
	book2 "github.com/ryota1116/stacked_books/usecase/book"
	"net/http"
)

type BookHandlerInterface interface {
	SearchBooks(w http.ResponseWriter, r *http.Request)
}

type bookHandler struct {
	bookUseCase book2.BookUseCaseInterface
}

func NewBookHandler(bu book2.BookUseCaseInterface) BookHandlerInterface {
	return &bookHandler{
		bookUseCase: bu,
	}
}

// SearchBooks : 外部APIを用いた書籍検索のエンドポイント
func (bh bookHandler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3002")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// リクエストパラメーターを取得
	var requestParameter = search_books2.RequestParameter{
		Title: r.FormValue("title"),
	}

	// リクエストボディのバリデーション
	isValid, validMsg := search_books2.FormValidator{
		GoogleBooksApiRequestParameter: requestParameter,
	}.Validate()

	if !isValid {
		// クライアントにHTTPレスポンスを返す
		response := httpResponse.Response{
			StatusCode:   http.StatusUnprocessableEntity,
			ResponseBody: validMsg,
		}
		response.ReturnResponse(w)
		return
	}

	// 外部APIで書籍を検索
	responseFromGoogleBooksAPI, err := bh.bookUseCase.SearchBooks(requestParameter.Title)
	// 外部APIリクエストでエラーが発生した場合
	if err != nil {
		httpResponse.Return500Response(w, errors.New("検索に失敗しました"))
		return
	}

	httpResponse.Response{
		StatusCode: http.StatusOK,
		// GoogleBooksAPIのJSONレスポンスの構造体から、 書籍検索用のHTTPレスポンスボディ構造体を生成する
		ResponseBody: book.SearchBooksResponseGenerator{
			ResponseBodyFromGoogleBooksAPI: responseFromGoogleBooksAPI,
		}.Execute(),
	}.ReturnResponse(w)
}
