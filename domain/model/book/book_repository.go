package book

type BookRepository interface {
	FindOrCreateByGoogleBooksId(GoogleBooksId string) Book
	FindAllByUserId(userId int) ([]Book, error)
}
