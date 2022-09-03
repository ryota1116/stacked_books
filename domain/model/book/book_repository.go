package book

type BookRepository interface {
	FindOrCreateByGoogleBooksId(GoogleBooksId string) Book
}
