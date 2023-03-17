package book

type BookRepository interface {
	FindAllByUserId(userId int) ([]Book, error)
	FindOneByGoogleBooksId(GoogleBooksId string) (Book, error)
	Save(book Book) error
}
