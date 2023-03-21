package user

type AvatarInterface interface {
	Value() *string
}

type avatar struct {
	value *string
}

func NewAvatar(value *string) AvatarInterface {
	return &avatar{value}
}

func (s *avatar) Value() *string {
	return s.value
}
