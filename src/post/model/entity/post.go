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
	IsEnabled() bool
	Enable()
	Disable()
}

type post struct {
	uuid        uuid.UUID
	author      entity.User
	title, body string
	enabled     bool
}

func NewPost(id uuid.UUID, author entity.User, title, body string) Post {
	return &post{
		uuid:    id,
		author:  author,
		title:   title,
		body:    body,
		enabled: true,
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

func (p *post) Enable() {
	p.enabled = true
}

func (p *post) Disable() {
	p.enabled = false
}

func (p *post) IsEnabled() bool {
	return p.enabled
}
