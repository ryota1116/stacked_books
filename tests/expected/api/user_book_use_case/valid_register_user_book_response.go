package user_book_use_case

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
)

// ExpectedBookStructForRegisterUserBook userBookUseCase.RegisterUserBookの戻り値で期待するBook構造体
// 構造体を定数constに格納することは出来ないので、変数宣言している
var ExpectedBookStructForRegisterUserBook = book.Book{
	GoogleBooksId:  "Wx1dLwEACAAJ",
	Title:          "リーダブルコード",
	Description:    "読んでわかるコードの重要性と方法について解説",
	Isbn10:         "4873115655",
	Isbn13:         "9784873115658",
	PageCount:      237,
	PublishedYear:  2012,
	PublishedMonth: 6,
	PublishedDate:  0,
}

// ExpectedUserBookStructForRegisterUserBook userBookUseCase.RegisterUserBookの戻り値で期待するUserBook構造体
// 構造体を定数constに格納することは出来ないので、変数宣言している
var ExpectedUserBookStructForRegisterUserBook = userbook.UserBook{
	Id:     1,
	UserId: 1,
	BookId: 1,
	Status: 1,
	Memo:   "メモメモメモ",
}
