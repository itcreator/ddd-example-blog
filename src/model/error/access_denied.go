package error

import (
	"fmt"
	"model/entity"
)

type accessDeniedError struct {
	actorName string
	user      entity.User
}

func NewAccessDeniedError(actorName string, user entity.User) error {
	return &accessDeniedError{actorName, user}
}

func (e *accessDeniedError) Error() string {
	return fmt.Sprintf("User %s doesn't have role %s", e.user.GetUserName(), e.actorName)
}
