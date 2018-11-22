package repository

import (
	"post/model/actor"
	"post/model/entity"
	userEntity "user/model/entity"
)

type PostEditor interface {
	LoadForPost(user userEntity.User, post entity.Post) (actor.PostEditor, error)
	//FindByUser(user entity.User) (actor.PostEditor, error)
	//Save(creator actor.PostEditor) error
}