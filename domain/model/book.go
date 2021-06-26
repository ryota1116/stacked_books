package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Id            int       `json:"id" gorm:"primaryKey;"`
	GoogleBooksId string    `json:"google_books_id" gorm:"primaryKey"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Image         string    `json:"image"`
	Isbn10        string    `json:"isbn_10"`
	Isbn13        string    `json:"isbn_13"`
	PageCount     int       `json:"page_count"`
	PublishAt     time.Time `json:"publish_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Users         []User    `gorm:"many2many:user_books;"`
}
