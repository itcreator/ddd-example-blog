package actor

import (
	userEntity "user/model/entity"
)

type PostViewer interface {
	//TODO: is this method needed?
	//ViewPost(title, body string) entity.Post
	GetUser() userEntity.User
}

type postViewer struct {
	user userEntity.User
}

func NewPostViewer(user userEntity.User) PostViewer {
	return &postViewer{user}
}

//func (pc *postViewer) ViewPost(title, body string) entity.Post {
//	return entity.NewPost(pc.user, title, body)
//}

func (pc *postViewer) GetUser() userEntity.User {
	return pc.user
}
