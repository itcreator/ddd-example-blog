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
	Delete()
	IsEnabled() bool
}

type comment struct {
	post    postEntity.Post
	uuid    uuid.UUID
	author  entity.User
	body    string
	enabled bool
}

func NewComment(id uuid.UUID, post postEntity.Post, author entity.User, body string) Comment {
	return &comment{
		post:    post,
		uuid:    id,
		author:  author,
		body:    body,
		enabled: true,
	}
}

func (c *comment) GetPost() postEntity.Post {
	return c.post
}

func (c *comment) GetUUID() uuid.UUID {
	return c.uuid
}

func (c *comment) GetBody() string {
	return c.body
}

func (c *comment) GetAuthor() entity.User {
	return c.author
}

func (c *comment) Update(body string) {
	c.body = body
}

func (c *comment) Delete() {
	c.body = ""
	c.enabled = false
}

func (c *comment) IsEnabled() bool {
	return c.enabled
}
