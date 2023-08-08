package book

type BookRepository interface {
	FindListByUserId(userId int) ([]BookInterface, error)
	FindOneByGoogleBooksId(GoogleBooksId string) (BookInterface, error)
	SaveOne(book BookInterface) error
	FindOneById(bookId int) (BookInterface, error)
}
