package user

import (
	"time"
)

type UserInterface interface {
	Id() IdInterface
	UserName() UserNameInterface
	Email() EmailInterface
	Password() PasswordInterface
	Avatar() AvatarInterface
	Role() RoleInterface
	CreatedAt() *time.Time
	UpdatedAt() *time.Time
}

type user struct {
	id        IdInterface
	userName  UserNameInterface
	email     EmailInterface
	password  PasswordInterface
	avatar    AvatarInterface
	role      RoleInterface
	createdAt *time.Time
	updatedAt *time.Time
}

func NewUser(
	id *int,
	userName string,
	email string,
	password string,
	avatar *string,
	role int,
	createdAt *time.Time,
	updatedAt *time.Time,
) (UserInterface, error) {
	return &user{
		NewId(id),
		NewUserName(userName),
		NewEmail(email),
		NewPassword(password),
		NewAvatar(avatar),
		NewRole(role),
		createdAt,
		updatedAt,
	}, nil
}

func (u *user) Id() IdInterface {
	return u.id
}

func (u *user) UserName() UserNameInterface {
	return u.userName
}

func (u *user) Email() EmailInterface {
	return u.email
}

func (u *user) Password() PasswordInterface {
	return u.password
}

func (u *user) Avatar() AvatarInterface {
	return u.avatar
}

func (u *user) Role() RoleInterface {
	return u.role
}

func (u *user) CreatedAt() *time.Time {
	return u.createdAt
}

func (u *user) UpdatedAt() *time.Time {
	return u.updatedAt
}
