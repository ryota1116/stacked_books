package user

type UserCreateCommand struct {
	UserName string
	Email    string
	Password string
	Avatar   *string
	Role     int
}
