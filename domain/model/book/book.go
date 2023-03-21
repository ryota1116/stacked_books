package book

import (
	"time"
)

type BookInterface interface {
	Id() IdInterface
	GoogleBooksId() GoogleBooksIdInterface
	Title() TitleInterface
	Description() DescriptionInterface
	Image() *string
	Isbn10() Isbn10Interface
	Isbn13() Isbn13Interface
	PageCount() PageCountInterface
	PublishedYear() PublishedYearInterface
	PublishedMonth() PublishedMonthInterface
	PublishedDate() PublishedDateInterface
	CreatedAt() *time.Time
}

// Book : 本のドメインモデル
// TODO: ドメインモデルをORMのEntityの用に使ってしまっているから、datasource/userbook/entity.go?を作成する。
// NOTE: Isbn_10カラムを取得する場合フィールドをIsbn_10にする必要がある(=>Isbn10では取得できない)
type book struct {
	id             IdInterface
	googleBooksId  GoogleBooksIdInterface
	title          TitleInterface
	description    DescriptionInterface
	image          *string
	isbn10         Isbn10Interface
	isbn13         Isbn13Interface
	pageCount      PageCountInterface
	publishedYear  PublishedYearInterface
	publishedMonth PublishedMonthInterface
	publishedDate  PublishedDateInterface
	createdAt      *time.Time
}

func NewBook(
	id *int,
	googleBooksId string,
	title string,
	description *string,
	image *string,
	isbn10 *string,
	isbn13 *string,
	pageCount int,
	publishedYear *int,
	publishedMonth *int,
	publishedDate *int,
	createdAt *time.Time,
) (BookInterface, error) {
	return &book{
		NewId(id),
		NewGoogleBooksId(googleBooksId),
		NewTitle(title),
		NewDescription(description),
		image,
		NewIsbn10(isbn10),
		NewIsbn13(isbn13),
		NewPageCount(pageCount),
		NewPublishedYear(publishedYear),
		NewPublishedMonth(publishedMonth),
		NewPublishedDate(publishedDate),
		createdAt,
	}, nil
}

func (b *book) Id() IdInterface {
	return b.id
}

func (b *book) GoogleBooksId() GoogleBooksIdInterface {
	return b.googleBooksId
}

func (b *book) Title() TitleInterface {
	return b.title
}

func (b *book) Description() DescriptionInterface {
	return b.description
}

func (b *book) Image() *string {
	return b.image
}

func (b *book) Isbn10() Isbn10Interface {
	return b.isbn10
}

func (b *book) Isbn13() Isbn13Interface {
	return b.isbn13
}

func (b *book) PageCount() PageCountInterface {
	return b.pageCount
}

func (b *book) PublishedYear() PublishedYearInterface {
	return b.publishedYear
}

func (b *book) PublishedMonth() PublishedMonthInterface {
	return b.publishedMonth
}

func (b *book) PublishedDate() PublishedDateInterface {
	return b.publishedDate
}

func (b *book) CreatedAt() *time.Time {
	return b.createdAt
}
