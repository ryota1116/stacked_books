package user

// interfaceを定義し、技術的関心事を扱うinfra層がrepositoryの処理を実装する（依存性逆転）
type UserRepository interface {
	Create(user UserInterface) (UserInterface, error)
	FindOneByEmail(email string) (UserInterface, error)
	FindOne(userId int) (UserInterface, error)
}
