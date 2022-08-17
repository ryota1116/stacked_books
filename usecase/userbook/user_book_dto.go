package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
)

type DtoGenerator struct {
	Books []book.Book
}

type UserBookDto struct {
	ID             int
	GoogleBooksId  string
	Title          string
	Description    string
	Isbn10         string
	Isbn13         string
	PageCount      int
	PublishedYear  int
	PublishedMonth int
	PublishedDate  int
	//Status int
	//Memo   string
}

func (dto DtoGenerator) Execute() []UserBookDto {
	var booksDto []UserBookDto

	for _, book := range dto.Books {
		dto := UserBookDto{
			ID:             book.Id,
			GoogleBooksId:  book.GoogleBooksId,
			Title:          book.Title,
			Description:    book.Description,
			Isbn10:         book.Isbn_10,
			Isbn13:         book.Isbn_13,
			PageCount:      book.PageCount,
			PublishedYear:  book.PublishedYear,
			PublishedMonth: book.PublishedMonth,
			PublishedDate:  book.PublishedDate,
		}

		booksDto = append(booksDto, dto)
	}

	return booksDto
}
