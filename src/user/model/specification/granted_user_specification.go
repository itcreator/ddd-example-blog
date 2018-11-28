package specification

import (
	"user/model/entity"
	"user/model/permission"
)

type grantedUserSpecification struct {
	permissionChecker permission.Checker
}

func (s *grantedUserSpecification) IsSatisfiedBy(user entity.User) bool {
	result, _ := s.permissionChecker.CheckPermission(user)

	return result
}

func NewGrantedUserSpecification(checker permission.Checker) UserSpecification {
	return &grantedUserSpecification{checker}
}
