package mock

import (
	"github.com/google/uuid"
	"user/model"
	"user/model/entity"
)

type authenticator struct {
	storage map[uuid.UUID]entity.User
}

func NewAuthenticator() model.Authenticator {
	return &authenticator{make(map[uuid.UUID]entity.User)}
}

func (a *authenticator) IsAuthenticated(user entity.User) (bool, error) {
	return nil != a.storage[user.GetUUID()], nil
}

func (a *authenticator) Authenticate(user entity.User) error {
	a.storage[user.GetUUID()] = user

	return nil
}
