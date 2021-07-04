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

	db.Model(&model.UserBook{}).Create(map[string]interface{}{
		"UserId": 1,
		"BookId": userBookParameter.Book.Id,
		"status": userBookParameter.Status,
		"memo": userBookParameter.Memo,
	})

	fmt.Println(userBookParameter)
	return userBookParameter
}
