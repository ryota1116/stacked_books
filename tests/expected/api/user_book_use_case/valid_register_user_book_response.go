package user_book_use_case

import "github.com/ryota1116/stacked_books/domain/model/dto"

// ExpectedRegisterUserBookResponse : userBookUseCaseのRegisterUserBookの戻り値で期待する構造体
// 構造体を定数constに格納することは出来ないので、変数宣言している
var ExpectedRegisterUserBookResponse = dto.RegisterUserBookResponse{
	Book: dto.Book{
		GoogleBooksId:  "Wx1dLwEACAAJ",
		Title:          "リーダブルコード",
		Description:    "読んでわかるコードの重要性と方法について解説",
		Isbn_10:        "4873115655",
		Isbn_13:        "9784873115658",
		PageCount:      237,
		PublishedYear:  2012,
		PublishedMonth: 6,
		PublishedDate:  0,
	},
	UserBook: dto.UserBook{
		Status: 0,
		Memo:   "メモメモメモ",
	},
}
