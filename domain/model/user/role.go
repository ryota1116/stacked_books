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

func (r *role) Value() int {
	return r.value
}
