package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SearchWord struct {
	Title	string	`json:"title"`
}

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


type SearchBookResult struct {
	Title	string			`json:"title"`
	Authors	[]string				`json:"authors"`
	Description	string		`json:"description"`
	Isbn10 string			`json:"isbn_10"`
	Isbn13 string			`json:"isbn_13"`
	PageCount int 			`json:"page_count"`
	RegisteredAt string	`json:"created_at"`
}

type SearchBookResults []SearchBookResult


// Booksテーブルの構造体
type Book struct {
	Id		int64			`json:"id"`
	Title	string			`json:"title"`
	Description	string		`json:"description"`
	Isbn10 string			`json:"isbn_10"`
	Isbn13 string			`json:"isbn_13"`
	PageCount int 			`json:"page_count"`
	RegisteredAt time.Time	`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
}

type Books []Book

const URLForGoogleBooksAPI = "https://www.googleapis.com/books/v1/volumes?q="

func SearchBooks(w http.ResponseWriter, r *http.Request)  {
	var searchWord SearchWord
	responseBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
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
	//var byteURL = make([]byte, 0, 100) // 100byte のキャパシティを確保
	//byteURL = append(byteURL, []byte(URLForGoogleBooksAPI))
	//byteURL = append(byteURL, searchWord[0])
	//searchURL := string(byteURL)

	// 文字列を連結してURLを生成
	searchURL := URLForGoogleBooksAPI
	searchURL += searchWord
	fmt.Println(searchURL)

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

	var bookResponse ResponseByGoogleBooksAPI
	if err := json.Unmarshal(body, &bookResponse); err != nil {
		return SearchBookResults{}, err
	}

	// TODO: 関数切り分け
	var searchBookResults = SearchBookResults{}
	for _, item := range bookResponse.Items {
		searchBookResult := SearchBookResult{
			Title:			item.VolumeInfo.Title,
			Authors:		item.VolumeInfo.Authors,
			Description:	item.VolumeInfo.Description,
			Isbn10:			"",
			Isbn13:			"",
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
