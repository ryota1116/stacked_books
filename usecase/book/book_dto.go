package book

import (
	model "github.com/ryota1116/stacked_books/domain/model/book"
)

type BookDtoGenerator struct {
	Book model.BookInterface
}

type BookDto struct {
	Id             int
	GoogleBooksId  string
	Title          string
	Description    *string
	Image          *string
	Isbn10         *string
	Isbn13         *string
	PageCount      int
	PublishedYear  *int
	PublishedMonth *int
	PublishedDate  *int
}

func (dtog BookDtoGenerator) Execute() BookDto {
	var bookDto = BookDto{
		Id:             *dtog.Book.Id().Value(),
		GoogleBooksId:  dtog.Book.GoogleBooksId().Value(),
		Title:          dtog.Book.Title().Value(),
		Description:    dtog.Book.Description().Value(),
		Image:          dtog.Book.Image(),
		Isbn10:         dtog.Book.Isbn10().Value(),
		Isbn13:         dtog.Book.Isbn13().Value(),
		PageCount:      dtog.Book.PageCount().Value(),
		PublishedYear:  dtog.Book.PublishedYear().Value(),
		PublishedMonth: dtog.Book.PublishedMonth().Value(),
		PublishedDate:  dtog.Book.PublishedDate().Value(),
	}

	return bookDto
}
