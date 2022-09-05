package find_user_books

import (
	"github.com/ryota1116/stacked_books/usecase/book"
)

type FindUserBooksResponseGenerator struct {
	BooksDto []book.BookDto
}

type FindUserBooksResponse struct {
	Books []Book `json:"books"`
}

// TODO
type Book struct {
	ID             int    `json:"id"`
	GoogleBooksId  string `json:"google_books_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Isbn10         string `json:"isbn_10"`
	Isbn13         string `json:"isbn_13"`
	PageCount      int    `json:"page_count"`
	PublishedYear  int    `json:"published_year"`
	PublishedMonth int    `json:"published_month"`
	PublishedDate  int    `json:"published_date"`
	//Status         int    `json:"status"`
	//Memo           string `json:"memo"`
}

func (fubrg FindUserBooksResponseGenerator) Execute() FindUserBooksResponse {
	var books []Book

	for _, bookDto := range fubrg.BooksDto {
		userBook := Book{
			ID:             bookDto.Id,
			GoogleBooksId:  bookDto.GoogleBooksId,
			Title:          bookDto.Title,
			Description:    bookDto.Description,
			Isbn10:         *bookDto.Isbn_10,
			Isbn13:         *bookDto.Isbn_13,
			PageCount:      *bookDto.PageCount,
			PublishedYear:  *bookDto.PublishedYear,
			PublishedMonth: *bookDto.PublishedMonth,
			PublishedDate:  *bookDto.PublishedDate,
		}

		books = append(books, userBook)
	}

	return FindUserBooksResponse{Books: books}
}
