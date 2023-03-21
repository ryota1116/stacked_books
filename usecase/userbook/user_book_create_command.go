package userbook

type UserBookCreateCommand struct {
	UserId   int
	Book     Book
	UserBook UserBook
}

type Book struct {
	GoogleBooksId  string
	Title          string
	Description    *string
	Image          *string
	Isbn10         *string
	Isbn13         *string
	PageCount      int
	PublishedYear  *int
	PublishedMonth *int
	PublishedDate  *int
}

type UserBook struct {
	Status int
	Memo   *string
}
