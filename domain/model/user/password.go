package user

type PasswordInterface interface {
	Value() string
}

type password struct {
	value string
}

func NewPassword(value string) PasswordInterface {
	return &password{value}
}

func (p *password) Value() string {
	return p.value
}
