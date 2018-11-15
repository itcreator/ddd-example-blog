package actor

import "model/entity"

type PostCreator interface {
	CreatePost(title, body string) entity.Post
	GetUser() entity.User
}

type postCreator struct {
	user entity.User
}

func NewPostCreator(user entity.User) PostCreator {
	return &postCreator{user}
}

func (pc *postCreator) CreatePost(title, body string) entity.Post {
	return entity.NewPost(pc.user, title, body)
}

func (pc *postCreator) GetUser() entity.User {
	return pc.user
}
