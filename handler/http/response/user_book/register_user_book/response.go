package register_user_book

import "github.com/ryota1116/stacked_books/domain/model"

// Response : 書籍登録用のレスポンス構造体
type Response struct {
	Book     `json:"book"`
	UserBook `json:"user_book"`
}

type Book struct{
	GoogleBooksId string    `json:"google_books_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Isbn10        string    `json:"isbn_10"`
	Isbn13        string    `json:"isbn_13"`
	PageCount     int       `json:"page_count"`
	PublishedYear   int 	`json:"published_year"`
	PublishedMonth   int 	`json:"published_month"`
	PublishedDate   int 	`json:"published_date"`
}

type UserBook struct {
	Status        int       `json:"status"`
	Memo          string    `json:"memo"`
}

// BuildResponse : 書籍登録用のレスポンス構造体を生成する
func BuildResponse(book model.Book, userBook model.UserBook) Response {
	userBookResponse := Response{
		Book:     Book{
			GoogleBooksId:  book.GoogleBooksId,
			Title:          book.Title,
			Description:    book.Description,
			Isbn10:         book.Isbn_10,
			Isbn13:         book.Isbn_13,
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
