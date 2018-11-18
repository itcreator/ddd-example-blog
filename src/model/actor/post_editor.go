package actor

import "model/entity"

type PostEditor interface {
	CreatePost(title, body string) entity.Post
	GetUser() entity.User
}

type postEditor struct {
	user entity.User
}

func NewPostEditor(user entity.User) PostEditor {
	return &postEditor{user}
}

func (pc *postEditor) CreatePost(title, body string) entity.Post {
	return entity.NewPost(pc.user, title, body)
}

func (pc *postEditor) GetUser() entity.User {
	return pc.user
}
