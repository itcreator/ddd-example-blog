package repository

import (
	"github.com/google/uuid"
	"user/model/entity"
)

type User interface {
	Save(user entity.User) error
	Load(uuid uuid.UUID) (entity.User, error)
	Delete(user entity.User) error
}
