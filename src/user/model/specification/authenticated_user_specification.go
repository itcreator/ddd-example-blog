package specification

import "user/model/entity"

type adminSpecification struct {
	//todo: add authentication service dependency
}

func (s *adminSpecification) IsSatisfiedBy(user entity.User) bool {
	//todo: request auth service for status
	return true
}

func NewAdminSpecification() UserSpecification {
	return &adminSpecification{}
}
