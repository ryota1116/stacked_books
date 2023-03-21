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

func (s *password) Value() string {
	return s.value
}
