package mock

import (
	"model/actor"
	"model/entity"
	"model/repository"
)

type postViewerRepository struct {
}

func NewPostViewerRepository() repository.PostViewer {
	return &postViewerRepository{}
}

func (r *postViewerRepository) LoadForPost(user entity.User, post entity.Post) (actor.PostViewer, error) {
	//TODO: add possibility to create private posts and to grant users
	return actor.NewPostViewer(user), nil
}
