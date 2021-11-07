package googleBooksApi

// SearchBooksResponseGenerator : 書籍検索用レスポンスボディのジェネレーター
type SearchBooksResponseGenerator struct {
	ResponseBodyFromGoogleBooksAPI ResponseBodyFromGoogleBooksAPI `json:"response_body_from_google_books_api"`
}

// execute : GoogleBooksAPIのJSONレスポンスの構造体から、
// 必要なフィールドだけをセットした書籍検索用のレスポンスボディ構造体を生成する
func (sbrg SearchBooksResponseGenerator) execute() SearchBooksResponses {
	// 書籍検索のレスポンスボディ構造体
	var searchBooksResponses = SearchBooksResponses{}

	// GoogleBooksAPIのJSONレスポンスから、書籍検索用のレスポンスボディ構造体を生成する
	// 検索結果一覧が配列で返ってくるため、slice型に格納して返す
	for _, item := range sbrg.ResponseBodyFromGoogleBooksAPI.Items {
		searchBooksResponse := SearchBooksResponse{
			GoogleBooksId:item.ID,
			Title:        item.VolumeInfo.Title,
			Authors:      item.VolumeInfo.Authors,
			Description:  item.VolumeInfo.Description,
			Isbn10:       "",
			Isbn13:       "",
			PageCount:    item.VolumeInfo.PageCount,
			RegisteredAt: item.VolumeInfo.PublishedDate,
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

		searchBooksResponses = append(searchBooksResponses, searchBooksResponse)
	}

	return searchBooksResponses
}