package repository

import (
	"post/model/actor"
	"post/model/entity"
	userEntity "user/model/entity"
)

type PostViewer interface {
	LoadForPost(user userEntity.User, post entity.Post) (actor.PostViewer, error)
	//FindByUser(user entity.User) (actor.PostViewer, error)
	//Save(creator actor.PostViewer) error
}
