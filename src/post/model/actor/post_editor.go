package actor

import (
	"post/model/entity"
	userEntity "user/model/entity"
)

type PostEditor interface {
	CreatePost(title, body string) entity.Post
	GetUser() userEntity.User
}

type postEditor struct {
	user userEntity.User
}

func NewPostEditor(user userEntity.User) PostEditor {
	return &postEditor{user}
}

func (pc *postEditor) CreatePost(title, body string) entity.Post {
	return entity.NewPost(pc.user, title, body)
}

func (pc *postEditor) GetUser() userEntity.User {
	return pc.user
}
