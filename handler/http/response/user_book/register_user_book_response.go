package user_book

import "github.com/ryota1116/stacked_books/domain/model"

type RegisterUserBookResponseGenerator struct {
	model.Book
	model.UserBook
}

type RegisterUserBookResponse struct {
	Book     Book     `json:"userbook"`
	UserBook UserBook `json:"userbook"`
}

type Book struct {
	GoogleBooksId  string `json:"google_books_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Isbn10         string `json:"isbn_10"`
	Isbn13         string `json:"isbn_13"`
	PageCount      int    `json:"page_count"`
	PublishedYear  int    `json:"published_year"`
	PublishedMonth int    `json:"published_month"`
	PublishedDate  int    `json:"published_date"`
}

type UserBook struct {
	Status int    `json:"status"`
	Memo   string `json:"memo"`
}

// Execute : 書籍登録用のレスポンス構造体を生成する
func (rubrg RegisterUserBookResponseGenerator) Execute() RegisterUserBookResponse {
	return RegisterUserBookResponse{
		Book: Book{
			GoogleBooksId:  rubrg.Book.GoogleBooksId,
			Title:          rubrg.Book.Title,
			Description:    rubrg.Book.Description,
			Isbn10:         rubrg.Book.Isbn_10,
			Isbn13:         rubrg.Book.Isbn_13,
			PageCount:      rubrg.Book.PageCount,
			PublishedYear:  rubrg.Book.PublishedYear,
			PublishedMonth: rubrg.Book.PublishedMonth,
			PublishedDate:  rubrg.Book.PublishedDate,
		},
		UserBook: UserBook{
			Status: rubrg.UserBook.Status,
			Memo:   rubrg.UserBook.Memo,
		},
	}
}
