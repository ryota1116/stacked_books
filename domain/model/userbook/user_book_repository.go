package userbook

type UserBookRepository interface {
	// SaveOne : UserBooksレコードを作成する
	SaveOne(userBook UserBookInterface) error
	// FindListByStatus : ユーザーの対象ステータスの本を一覧取得する
	FindListByStatus(userID int, status StatusInterface) ([]UserBookInterface, error)
}
