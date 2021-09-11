package handler

import "github.com/ryota1116/stacked_books/domain/model/googleBooksApi"

// BookUseCaseMock : BookUseCaseInterfaceを実装しているモック
type BookUseCaseMock struct {}

func (bu BookUseCaseMock) SearchBooks(requestParameter googleBooksApi.RequestParameter) (googleBooksApi.SearchBooksResponses, error) {
	return googleBooksApi.SearchBooksResponses{
		{
			Title:        "リーダブルコード",
			Authors:      []string{"Dustin Boswell", "Trevor Foucher"},
			Description:  "読んでわかるコードの重要性と方法について解説",
			Isbn10:       "4873115655",
			Isbn13:       "9784873115658",
			PageCount:    0237,
			RegisteredAt: "2012-06",
		},
		{
			Title:        "ExcelVBAを実務で使い倒す技術",
			Authors:      []string{"高橋宣成"},
			Description:  "本書では、VBAを実務の現場で活かすための知識(テクニック)と知恵(考え方とコツ)を教えます!",
			Isbn10:       "4798049999",
			Isbn13:       "9784798049991",
			PageCount:    289,
			RegisteredAt: "2017-04",
		},
	}, nil
}
