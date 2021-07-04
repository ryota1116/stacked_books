package usecase

import (
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/infra/persistence"
)

func RegisterUserBook(userBookParameter model.UserBookParameter) model.UserBookParameter {
	db := persistence.DbConnect()
	//dbBook := model.Book{}
	// Booksが存在すればそのレコードを取得し、存在しなければ新しいレコードを作成する
	db.Where("google_books_id = ?", userBookParameter.GoogleBooksId).FirstOrCreate(&userBookParameter.Book)

	fmt.Println(userBookParameter)
	return userBookParameter
}
