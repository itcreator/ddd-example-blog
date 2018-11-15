package repository

import (
	"model/actor"
	"model/entity"
)

type PostViewer interface {
	LoadForPost(user entity.User, post entity.Post) (actor.PostViewer, error)
	//FindByUser(user entity.User) (actor.PostViewer, error)
	//Save(creator actor.PostViewer) error
}
