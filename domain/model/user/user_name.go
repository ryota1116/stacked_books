package user

type UserNameInterface interface {
	Value() string
}

type userName struct {
	value string
}

func NewUserName(value string) UserNameInterface {
	return &userName{value}
}

func (un *userName) Value() string {
	return un.value
}
