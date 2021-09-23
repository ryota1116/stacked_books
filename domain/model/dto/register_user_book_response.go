package dto

import "github.com/ryota1116/stacked_books/domain/model"

type RegisterUserBookResponse struct {
	Book `json:"book"`
	UserBook `json:"user_book"`
}

type Book struct{
	GoogleBooksId string    `json:"google_books_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Isbn_10        string    `json:"isbn_10"`
	Isbn_13        string    `json:"isbn_13"`
	PageCount     int       `json:"page_count"`
	PublishedYear   int 	`json:"published_year"`
	PublishedMonth   int 	`json:"published_month"`
	PublishedDate   int 	`json:"published_date"`
}

type UserBook struct {
	Status        int       `json:"status"`
	Memo          string    `json:"memo"`
}

// BuildRegisterUserBookResponse : RegisterUserBookResponse構造体を生成する
func BuildRegisterUserBookResponse(book model.Book, userBook model.UserBook) RegisterUserBookResponse {
	userBookResponse := RegisterUserBookResponse{
		Book:     Book{
			GoogleBooksId:  book.GoogleBooksId,
			Title:          book.Title,
			Description:    book.Description,
			Isbn_10:        book.Isbn_10,
			Isbn_13:        book.Isbn_13,
			PageCount:      book.PageCount,
			PublishedYear:  book.PublishedYear,
			PublishedMonth: book.PublishedMonth,
			PublishedDate:  book.PublishedDate,
		},
		UserBook: UserBook{
			Status: userBook.Status,
			Memo:   userBook.Memo,
		},
	}

	return userBookResponse
}
