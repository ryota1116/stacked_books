package book

type BookRepository interface {
	FindAllByUserId(userId int) ([]BookInterface, error)
	FindOneByGoogleBooksId(GoogleBooksId string) (BookInterface, error)
	Save(book BookInterface) error
	FindOneById(bookId int) (BookInterface, error)
}
