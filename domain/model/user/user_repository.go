package user

// interfaceを定義し、技術的関心事を扱うinfra層がrepositoryの処理を実装する（依存性逆転）
type UserRepository interface {
	// *型名でポイント型になる
	// *型名でUser型へのポイント型
	Create(user User) error
	FindOneByEmail(email string) (User, error)
	FindOne(userId int) (User, error)
}
