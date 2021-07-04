package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/usecase"
	"io/ioutil"
	"net/http"
)

type SearchWord struct {
	Title	string	`json:"title"`
}

// TODO: GoogleBooksAPI用のファイルに移す
// GoogleBooksAPIを叩いた時のレスポンスに合わせた構造体
type ResponseByGoogleBooksAPI struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []struct {
		Kind       string `json:"kind"`
		ID         string `json:"id"`
		Etag       string `json:"etag"`
		SelfLink   string `json:"selfLink"`
		VolumeInfo struct {
			Title               string   `json:"title"`
			Subtitle            string   `json:"subtitle"`
			Authors             []string `json:"authors"`
			PublishedDate       string   `json:"publishedDate"`
			Description         string   `json:"description"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
			ReadingModes struct {
				Text  bool `json:"text"`
				Image bool `json:"image"`
			} `json:"readingModes"`
			PageCount        int    `json:"pageCount"`
			PrintType        string `json:"printType"`
			AverageRating    int    `json:"averageRating"`
			RatingsCount     int    `json:"ratingsCount"`
			MaturityRating   string `json:"maturityRating"`
			AllowAnonLogging bool   `json:"allowAnonLogging"`
			ContentVersion   string `json:"contentVersion"`
			ImageLinks       struct {
				SmallThumbnail string `json:"smallThumbnail"`
				Thumbnail      string `json:"thumbnail"`
			} `json:"imageLinks"`
			Language            string `json:"language"`
			PreviewLink         string `json:"previewLink"`
			InfoLink            string `json:"infoLink"`
			CanonicalVolumeLink string `json:"canonicalVolumeLink"`
		} `json:"volumeInfo"`
		SaleInfo struct {
			Country     string `json:"country"`
			Saleability string `json:"saleability"`
			IsEbook     bool   `json:"isEbook"`
		} `json:"saleInfo"`
		AccessInfo struct {
			Country                string `json:"country"`
			Viewability            string `json:"viewability"`
			Embeddable             bool   `json:"embeddable"`
			PublicDomain           bool   `json:"publicDomain"`
			TextToSpeechPermission string `json:"textToSpeechPermission"`
			Epub                   struct {
				IsAvailable bool `json:"isAvailable"`
			} `json:"epub"`
			Pdf struct {
				IsAvailable bool `json:"isAvailable"`
			} `json:"pdf"`
			WebReaderLink       string `json:"webReaderLink"`
			AccessViewStatus    string `json:"accessViewStatus"`
			QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
		} `json:"accessInfo"`
		SearchInfo struct {
			TextSnippet string `json:"textSnippet"`
		} `json:"searchInfo"`
	} `json:"items"`
}

// 検索結果レスポンスの構造体(これをjson形式に変換する)
type SearchBookResult struct {
	GoogleBooksId	string	`json:"google_books_id"`
	Title	string			`json:"title"`
	Authors	[]string				`json:"authors"`
	Description	string		`json:"description"`
	Isbn10 string			`json:"isbn_10"`
	Isbn13 string			`json:"isbn_13"`
	PageCount int 			`json:"page_count"`
	RegisteredAt string	`json:"registered_at"`
}

type SearchBookResults []SearchBookResult

const URLForGoogleBooksAPI = "https://www.googleapis.com/books/v1/volumes?q="

// booksを参照→同じのあればそれを使って、user_booksを作成
func RegisterUserBook(w http.ResponseWriter, r *http.Request) {
	//
	book := model.UserBookParameter{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println(err)
	}

	//認証
	//if VerifyToken(w, r) {
	//}

	dbBook := usecase.RegisterUserBook(book)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbBook)
}

// 書籍を検索するメソッド
func SearchBooks(w http.ResponseWriter, r *http.Request)  {
	var searchWord SearchWord
	responseBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// byte型をデコードしてSearchWord構造体に格納
	if err := json.Unmarshal(responseBodyBytes, &searchWord); err != nil {
		panic(err)
	}

	// 外部APIで書籍を検索
	searchBooksResult, err := SearchBooksByGoogleBooksAPI(searchWord.Title)
	if err != nil {
		json.NewEncoder(w).Encode("検索に失敗しました")
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	json.NewEncoder(w).Encode(searchBooksResult)
	w.Header().Set("Content-Type", "application/json")
}

func SearchBooksByGoogleBooksAPI(searchWord string) (SearchBookResults, error) {
	// TODO: 文字列の結合は処理遅いから、100byteキャパシティ与える方法に変える
	// 文字列を連結してURLを生成
	searchURL := URLForGoogleBooksAPI
	searchURL += searchWord
	fmt.Println(searchURL)

	// APIを叩く
	res, err := http.Get(searchURL)

	if err != nil {
		fmt.Println(err)
		return SearchBookResults{}, err
		// fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}

	defer res.Body.Close()//関数終了時にクローズ

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return SearchBookResults{}, err
	}

	var resByGoogleBooksAPI ResponseByGoogleBooksAPI
	if err := json.Unmarshal(body, &resByGoogleBooksAPI); err != nil {
		return SearchBookResults{}, err
	}

	// TODO: 関数切り分け
	var searchBookResults SearchBookResults
	for _, item := range resByGoogleBooksAPI.Items {
		searchBookResult := SearchBookResult{
			GoogleBooksId:	item.ID,
			Title:			item.VolumeInfo.Title,
			Authors:		item.VolumeInfo.Authors,
			Description:	item.VolumeInfo.Description,
			PageCount:		item.VolumeInfo.PageCount,
			RegisteredAt:	item.VolumeInfo.PublishedDate,
		}

		for _, isbn := range item.VolumeInfo.IndustryIdentifiers {
			switch isbn.Type {
			case "ISBN_10":
				searchBookResult.Isbn10 = isbn.Identifier
			case "ISBN_13":
				searchBookResult.Isbn13 = isbn.Identifier
			}
		}

		// 内部で配列の参照を持つ可変長のリストsliceに要素を追加する
		searchBookResults = append(searchBookResults, searchBookResult)
	}
	return searchBookResults, nil
}

//func ConvertToSearchBookResults(bookResponse ResponseByGoogleBooksAPI) SearchBookResults {
//	var searchBookResults = SearchBookResults{}
//	for _, item := range bookResponse.Items {
//		searchBookResult := SearchBookResult{
//			Title:			item.VolumeInfo.Title,
//			Authors:		item.VolumeInfo.Authors,
//			Description:	item.VolumeInfo.Description,
//			Isbn10:			"",
//			Isbn13:			"",
//			PageCount:		item.VolumeInfo.PageCount,
//			RegisteredAt:	item.VolumeInfo.PublishedDate,
//		}
//
//		for _, isbn := range item.VolumeInfo.IndustryIdentifiers {
//			switch isbn.Type {
//			case "ISBN_10":
//				searchBookResult.Isbn10 = isbn.Identifier
//			case "ISBN_13":
//				searchBookResult.Isbn13 = isbn.Identifier
//			}
//		}
//
//		a = append(searchBookResults, searchBookResult)
//		return a
//	}
//	return a
//}
//
