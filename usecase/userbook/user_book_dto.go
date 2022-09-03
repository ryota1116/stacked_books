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

	for _, b := range dto.Books {
		dto := UserBookDto{
			ID:             b.Id,
			GoogleBooksId:  b.GoogleBooksId,
			Title:          b.Title,
			Description:    b.Description,
			Isbn10:         b.Isbn_10,
			Isbn13:         b.Isbn_13,
			PageCount:      b.PageCount,
			PublishedYear:  b.PublishedYear,
			PublishedMonth: b.PublishedMonth,
			PublishedDate:  b.PublishedDate,
		}

		booksDto = append(booksDto, dto)
	}

	return booksDto
}
