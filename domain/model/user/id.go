package user

type IdInterface interface {
	Value() *int
}

type id struct {
	value *int
}

func NewId(value *int) IdInterface {
	return &id{value}
}

func (s *id) Value() *int {
	return s.value
}
