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

func (a *avatar) Value() *string {
	return a.value
}
