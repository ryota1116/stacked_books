package model

import (
	"time"
)

type Book struct {
	Id            int       `json:"id" gorm:"primaryKey;"`
	GoogleBooksId string    `json:"google_books_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Image         string    `json:"image"`
	Isbn_10        string    `json:"isbn_10"`
	Isbn_13        string    `json:"isbn_13"`
	PageCount     int       `json:"page_count"`
	PublishedYear   int 	`json:"published_year"`
	PublishedMonth   int 	`json:"published_month"`
	PublishedDate   int 	`json:"published_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Users         []User    `gorm:"many2many:user_books;"`
}
