package repository

import (
	"model/actor"
	"model/entity"
)

type PostEditor interface {
	LoadForPost(user entity.User, post entity.Post) (actor.PostViewer, error)
	//FindByUser(user entity.User) (actor.PostEditor, error)
	//Save(creator actor.PostEditor) error
}
