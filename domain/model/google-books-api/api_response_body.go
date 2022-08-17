package google_books_api

// ResponseBodyFromGoogleBooksAPI : GoogleBooksAPIを叩いた時のJSONレスポンスを格納する構造体
type ResponseBodyFromGoogleBooksAPI struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID         string     `json:"id"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title               string               `json:"title"`
	Authors             []string             `json:"authors"`
	PublishedDate       string               `json:"publishedDate"`
	Description         string               `json:"description"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
	PageCount           int                  `json:"pageCount"`
}

type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}
