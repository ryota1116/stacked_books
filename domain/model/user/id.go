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

func (i *id) Value() *int {
	return i.value
}
