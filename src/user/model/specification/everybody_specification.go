package specification

import "user/model/entity"

type everybodySpecification struct {
	//todo: add authentication service dependency
}

func (s *everybodySpecification) IsSatisfiedBy(user entity.User) bool {
	return true
}

func NewEverybodySpecification() UserSpecification {
	return &everybodySpecification{}
}
