package actor

import (
	"post/model/entity"
	userEntity "user/model/entity"
)

type PostCreator interface {
	CreatePost(title, body string) entity.Post
	GetUser() userEntity.User
}

type postCreator struct {
	user userEntity.User
}

func NewPostCreator(user userEntity.User) PostCreator {
	return &postCreator{user}
}

func (pc *postCreator) CreatePost(title, body string) entity.Post {
	return entity.NewPost(pc.user, title, body)
}

func (pc *postCreator) GetUser() userEntity.User {
	return pc.user
}
