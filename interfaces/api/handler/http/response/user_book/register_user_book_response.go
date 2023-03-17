package user_book

import (
	"github.com/ryota1116/stacked_books/usecase/book"
	"github.com/ryota1116/stacked_books/usecase/userbook"
)

type RegisterUserBookResponseGenerator struct {
	BookDto     book.BookDto
	UserBookDto userbook.UserBookDto
}

type RegisterUserBookResponse struct {
	GoogleBooksId  string `json:"google_books_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Isbn10         string `json:"isbn_10"`
	Isbn13         string `json:"isbn_13"`
	PageCount      int    `json:"page_count"`
	PublishedYear  int    `json:"published_year"`
	PublishedMonth int    `json:"published_month"`
	PublishedDate  int    `json:"published_date"`
	Status         int    `json:"status"`
	Memo           string `json:"memo"`
}

// Execute : 書籍登録用のレスポンス構造体を生成する
func (rubrg RegisterUserBookResponseGenerator) Execute() RegisterUserBookResponse {
	return RegisterUserBookResponse{
		GoogleBooksId:  rubrg.BookDto.GoogleBooksId,
		Title:          rubrg.BookDto.Title,
		Description:    rubrg.BookDto.Description,
		Isbn10:         *rubrg.BookDto.Isbn_10,
		Isbn13:         *rubrg.BookDto.Isbn_13,
		PageCount:      *rubrg.BookDto.PageCount,
		PublishedYear:  *rubrg.BookDto.PublishedYear,
		PublishedMonth: *rubrg.BookDto.PublishedMonth,
		PublishedDate:  *rubrg.BookDto.PublishedDate,
		Status:         rubrg.UserBookDto.Status,
		Memo:           rubrg.UserBookDto.Memo,
	}
}
