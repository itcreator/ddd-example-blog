package repository

import (
	"model/actor"
	"model/entity"
)

type PostCreator interface {
	FindByUser(user entity.User) (actor.PostCreator, error)
	Save(creator actor.PostCreator) error
}
