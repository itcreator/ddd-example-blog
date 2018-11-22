package entity

import "github.com/google/uuid"

type User interface {
	GetUUID() uuid.UUID
	GetUserName() string
}

type user struct {
	uuid     uuid.UUID
	userName string
}

func NewUser(userName string) User {
	return &user{
		uuid:     uuid.New(),
		userName: userName,
	}
}

func (u *user) GetUUID() uuid.UUID {
	return u.uuid
}

func (u *user) GetUserName() string {
	return u.userName
}
