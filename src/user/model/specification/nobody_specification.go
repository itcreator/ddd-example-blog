package specification

import "user/model/entity"

type nobodySpecification struct {
}

func (s *nobodySpecification) IsSatisfiedBy(user entity.User) bool {
	return false
}

func NewNobodySpecification() UserSpecification {
	return &nobodySpecification{}
}
