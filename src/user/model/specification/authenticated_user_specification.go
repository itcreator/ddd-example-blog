package specification

import (
	"user/model"
	"user/model/entity"
)

type authenticatedUserSpecification struct {
	authenticator model.Authenticator
}

func (s *authenticatedUserSpecification) IsSatisfiedBy(user entity.User) bool {
	ok, err := s.authenticator.IsAuthenticated(user)

	if err == nil && ok {
		return true
	} else {
		return false
	}
}

func NewAuthenticatedUserSpecification(auth model.Authenticator) UserSpecification {
	return &authenticatedUserSpecification{auth}
}
