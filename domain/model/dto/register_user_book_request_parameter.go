package dto

type RegisterUserBookRequestParameter struct {
	Book struct  {
		GoogleBooksId string    `json:"google_books_id"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		Isbn_10       string    `json:"isbn_10"`
		Isbn_13       string    `json:"isbn_13"`
		PageCount     int       `json:"page_count"`
		PublishedYear   int 	`json:"published_year"`
		PublishedMonth   int 	`json:"published_month"`
		PublishedDate   int 	`json:"published_date"`
	} `json:"book"`
	UserBook struct {
		Status        int       `json:"status"`
		Memo          string    `json:"memo"`
	}`json:"user_book"`
}
