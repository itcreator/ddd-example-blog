package entity

import (
	"github.com/google/uuid"
	postEntity "post/model/entity"
	"user/model/entity"
)

type Comment interface {
	GetUUID() uuid.UUID
	GetBody() string
	GetAuthor() entity.User
	GetPost() postEntity.Post
	Update(body string)
}

type comment struct {
	post   postEntity.Post
	uuid   uuid.UUID
	author entity.User
	body   string
}

func NewComment(id uuid.UUID, post postEntity.Post, author entity.User, body string) Comment {
	return &comment{
		post:   post,
		uuid:   id,
		author: author,
		body:   body,
	}
}

func (p *comment) GetPost() postEntity.Post {
	return p.post
}

func (p *comment) GetUUID() uuid.UUID {
	return p.uuid
}

func (p *comment) GetBody() string {
	return p.body
}

func (p *comment) GetAuthor() entity.User {
	return p.author
}

func (p *comment) Update(body string) {
	p.body = body
}
