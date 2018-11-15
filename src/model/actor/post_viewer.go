package actor

import "model/entity"

type PostViewer interface {
	//TODO: is this method needed?
	//ViewPost(title, body string) entity.Post
	GetUser() entity.User
}

type postViewer struct {
	user entity.User
}

func NewPostViewer(user entity.User) PostViewer {
	return &postViewer{user}
}

//func (pc *postViewer) ViewPost(title, body string) entity.Post {
//	return entity.NewPost(pc.user, title, body)
//}

func (pc *postViewer) GetUser() entity.User {
	return pc.user
}
