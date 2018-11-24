package mock

import (
	"github.com/google/uuid"
	"user/model/entity"
	"user/model/repository"
)

type userRepository struct {
	storage map[uuid.UUID]entity.User
}

func NewUserRepository() repository.User {
	return &userRepository{make(map[uuid.UUID]entity.User)}
}

func (r *userRepository) Save(user entity.User) error {
	r.storage[user.GetUUID()] = user

	return nil
}

func (r *userRepository) Load(uuid uuid.UUID) (entity.User, error) {
	return r.storage[uuid], nil
}

func (r *userRepository) Delete(user entity.User) error {
	if nil != r.storage[user.GetUUID()] {
		delete(r.storage, user.GetUUID())
	}

	return nil
}
