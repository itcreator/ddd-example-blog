package entity

import (
	"github.com/google/uuid"
	"user/model/entity"
)

type Post interface {
	GetUUID() uuid.UUID
	GetTitle() string
	GetBody() string
}

type post struct {
	uuid        uuid.UUID
	author      entity.User
	title, body string
}

func NewPost(author entity.User, title, body string) Post {
	return &post{
		uuid:   uuid.New(),
		author: author,
		title:  title,
		body:   body,
	}
}

func (u *post) GetUUID() uuid.UUID {
	return u.uuid
}

func (u *post) GetBody() string {
	return u.body
}

func (u *post) GetTitle() string {
	return u.title
}
