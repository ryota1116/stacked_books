package book

import (
	model "github.com/ryota1116/stacked_books/domain/model/searched_books/google_books_api"
)

// SearchBooksResponseGenerator : 書籍検索用レスポンスボディのジェネレーター
type SearchBooksResponseGenerator struct {
	ResponseBodyFromGoogleBooksApi model.ResponseBodyFromGoogleBooksApi `json:"response_body_from_google_books_api"`
}

// SearchBooksResponses : 書籍検索用のレスポンスボディ構造体のコレクション
type SearchBooksResponses struct {
	TypeName `json:"books"`
}

type TypeName []SearchBooksResponse

// SearchBooksResponse : 書籍検索用のレスポンスボディ構造体。
// GoogleBooksAPIを叩いた時に取得したJSONレスポンスのうち、
// 必要なフィールドだけをセットしたレスポンスボディの構造体。
type SearchBooksResponse struct {
	GoogleBooksId string   `json:"googleBooksId"`
	Title         string   `json:"title"`
	Authors       []string `json:"authors"`
	Description   string   `json:"description"`
	Isbn10        string   `json:"isbn10"`
	Isbn13        string   `json:"isbn13"`
	PageCount     int      `json:"pageCount"`
	RegisteredAt  string   `json:"publishedAt"`
}

// Execute : GoogleBooksAPIのJSONレスポンスの構造体から、
// 必要なフィールドだけをセットした書籍検索用のレスポンスボディ構造体を生成する
func (sbrg SearchBooksResponseGenerator) Execute() SearchBooksResponses {
	// 書籍検索のレスポンスボディ構造体

	var typeName = TypeName{}

	// GoogleBooksAPIのJSONレスポンスから、書籍検索用のレスポンスボディ構造体を生成する
	// 検索結果一覧が配列で返ってくるため、slice型に格納して返す
	for _, item := range sbrg.ResponseBodyFromGoogleBooksApi.Items {
		searchBooksResponse := SearchBooksResponse{
			GoogleBooksId: item.ID,
			Title:         item.VolumeInfo.Title,
			Authors:       item.VolumeInfo.Authors,
			Description:   item.VolumeInfo.Description,
			Isbn10:        "",
			Isbn13:        "",
			PageCount:     item.VolumeInfo.PageCount,
			RegisteredAt:  item.VolumeInfo.PublishedDate,
		}

		// ISBNが存在すれば構造体にセットする
		for _, isbn := range item.VolumeInfo.IndustryIdentifiers {
			switch isbn.Type {
			case "ISBN_10":
				searchBooksResponse.Isbn10 = isbn.Identifier
			case "ISBN_13":
				searchBooksResponse.Isbn13 = isbn.Identifier
			}
		}

		typeName = append(typeName, searchBooksResponse)
	}

	var searchBooksResponses = SearchBooksResponses{typeName}

	return searchBooksResponses
}
