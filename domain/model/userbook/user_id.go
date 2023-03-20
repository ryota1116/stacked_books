package userbook

type UserIdInterface interface {
	Value() int
}

type userId struct {
	value int
}

func NewUserId(value int) UserIdInterface {
	return &userId{value}
}

func (s *userId) Value() int {
	return s.value
}
