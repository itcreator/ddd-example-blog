package error

import (
	"fmt"
	userEntity "user/model/entity"
)

type accessDeniedError struct {
	actorName string
	user      userEntity.User
}

func NewAccessDeniedError(actorName string, user userEntity.User) error {
	return &accessDeniedError{actorName, user}
}

func (e *accessDeniedError) Error() string {
	return fmt.Sprintf("User %s doesn't have role %s", e.user.GetUserName(), e.actorName)
}
