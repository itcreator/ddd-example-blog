package model

import "user/model/entity"

type Authenticator interface {
	IsAuthenticated(user entity.User) (bool, error)
	Authenticate(user entity.User) error
}
