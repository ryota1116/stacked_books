package handler

import (
	"encoding/json"
	"github.com/ryota1116/stacked_books/domain/model/googleBooksApi"
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
	// 
	var requestParameter googleBooksApi.RequestParameter
	responseBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(responseBodyBytes, &requestParameter); err != nil {
		panic(err)
	}

	// 外部APIで書籍を検索
	searchBooksResult, err := bh.bookUseCase.SearchBooks(requestParameter)

	if err != nil {
		err := json.NewEncoder(w).Encode("検索に失敗しました")
		if err != nil {
			return 
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	err = json.NewEncoder(w).Encode(searchBooksResult)
	if err != nil {
		return 
	}
	w.Header().Set("Content-Type", "application/json")
}
