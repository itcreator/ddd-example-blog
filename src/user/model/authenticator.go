package model

import "user/model/entity"

//TODO: probably will be better separate it to two interfaces (Authenticator, Checker)
type Authenticator interface {
	IsAuthenticated(user entity.User) (bool, error)
	Authenticate(user entity.User) error
}
