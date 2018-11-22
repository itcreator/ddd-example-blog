package mock

import (
	"post/model/actor"
	"post/model/entity"
	"post/model/repository"
	userEntity "user/model/entity"
)

type postViewerRepository struct {
}

func NewPostViewerRepository() repository.PostViewer {
	return &postViewerRepository{}
}

func (r *postViewerRepository) LoadForPost(user userEntity.User, post entity.Post) (actor.PostViewer, error) {
	//TODO: add possibility to create private posts and to grant users
	return actor.NewPostViewer(user), nil
}
