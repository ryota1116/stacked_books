package find_user_books

import (
	"github.com/ryota1116/stacked_books/usecase/userbook"
)

type FindUserBooksResponseGenerator struct {
	UserBooksDto []userbook.UserBookDto
}

type FindUserBooksResponse struct {
	UserBooks []UserBook `json:"user_books"`
}

type UserBook struct {
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
	var userBooks []UserBook

	for _, userBookDto := range fubrg.UserBooksDto {
		userBook := UserBook{
			ID:             userBookDto.ID,
			GoogleBooksId:  userBookDto.GoogleBooksId,
			Title:          userBookDto.Title,
			Description:    userBookDto.Description,
			Isbn10:         userBookDto.Isbn10,
			Isbn13:         userBookDto.Isbn13,
			PageCount:      userBookDto.PageCount,
			PublishedYear:  userBookDto.PublishedYear,
			PublishedMonth: userBookDto.PublishedMonth,
			PublishedDate:  userBookDto.PublishedDate,
		}

		userBooks = append(userBooks, userBook)
	}

	return FindUserBooksResponse{UserBooks: userBooks}
}
