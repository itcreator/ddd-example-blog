package repository

import (
	"post/model/actor"
	userEntity "user/model/entity"
)

type PostCreator interface {
	FindByUser(user userEntity.User) (actor.PostCreator, error)
	Save(creator actor.PostCreator) error
}
