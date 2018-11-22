package entity

import (
	"github.com/google/uuid"
	"user/model/entity"
)

type Post interface {
	GetUUID() uuid.UUID
	GetTitle() string
	GetBody() string
	GetAuthor() entity.User
	Update(title, body string)
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

func (p *post) GetUUID() uuid.UUID {
	return p.uuid
}

func (p *post) GetBody() string {
	return p.body
}

func (p *post) GetTitle() string {
	return p.title
}

func (p *post) GetAuthor() entity.User {
	return p.author
}

func (p *post) Update(title, body string) {
	p.title = title
	p.body = body
}
