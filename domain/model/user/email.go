package user

type EmailInterface interface {
	Value() string
}

type email struct {
	value string
}

func NewEmail(value string) EmailInterface {
	return &email{value}
}

func (e *email) Value() string {
	return e.value
}
