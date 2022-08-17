package model

import (
	"time"
)

// Book : 本のドメインモデル
// TODO: ドメインモデルをORMのEntityの用に使ってしまっているから、persistence/userbook/entity.go?を作成する。
// NOTE: Isbn_10カラムを取得する場合フィールドをIsbn_10にする必要がある(=>Isbn10では取得できない)
type Book struct {
	Id             int       `json:"id" gorm:"primaryKey;"`
	GoogleBooksId  string    `json:"google_books_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Image          string    `json:"image"`
	Isbn_10        string    `json:"isbn_10"`
	Isbn_13        string    `json:"isbn_13"`
	PageCount      int       `json:"page_count"`
	PublishedYear  int       `json:"published_year"`
	PublishedMonth int       `json:"published_month"`
	PublishedDate  int       `json:"published_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Users          []User    `gorm:"many2many:user_books;"`
}
