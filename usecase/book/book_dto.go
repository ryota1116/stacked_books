package book

import model "github.com/ryota1116/stacked_books/domain/model/book"

type BookDtoGenerator struct {
	Book model.Book
}

type BookDto struct {
	Id             int
	GoogleBooksId  string
	Title          string
	Description    string
	Image          *string
	Isbn_10        *string
	Isbn_13        *string
	PageCount      *int
	PublishedYear  *int
	PublishedMonth *int
	PublishedDate  *int
}

func (dtog BookDtoGenerator) Execute() BookDto {
	var bookDto = BookDto{
		Id:             dtog.Book.Id,
		GoogleBooksId:  dtog.Book.GoogleBooksId,
		Title:          dtog.Book.Title,
		Description:    dtog.Book.Description,
		Image:          &dtog.Book.Image,
		Isbn_10:        &dtog.Book.Isbn_10,
		Isbn_13:        &dtog.Book.Isbn_13,
		PageCount:      &dtog.Book.PageCount,
		PublishedYear:  &dtog.Book.PublishedYear,
		PublishedMonth: &dtog.Book.PublishedMonth,
		PublishedDate:  &dtog.Book.PublishedDate,
	}

	return bookDto
}
