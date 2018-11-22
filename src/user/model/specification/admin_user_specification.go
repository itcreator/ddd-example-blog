package specification

import "user/model/entity"

type authenticatedUserSpecification struct {
	//todo: add authentication service dependency
}

func (s *authenticatedUserSpecification) IsSatisfiedBy(user entity.User) bool {
	//todo: request auth service for status
	return true
}

func NewAuthenticatedUserSpecification() UserSpecification {
	return &authenticatedUserSpecification{}
}
