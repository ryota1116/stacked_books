package user

type RoleInterface interface {
	Value() int
}

type role struct {
	value int
}

func NewRole(value int) RoleInterface {
	return &role{value}
}

func (s *role) Value() int {
	return s.value
}
