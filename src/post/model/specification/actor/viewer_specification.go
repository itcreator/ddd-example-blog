package actor

import (
	"user/model/entity"
	"user/model/specification"
)

type viewerSpecification struct {
	//todo: add authentication service dependency
}

// IsSatisfiedBy allows everybody view posts
func (s *viewerSpecification) IsSatisfiedBy(user entity.User) bool {
	return true
}

func NewViewerSpecification() specification.UserSpecification {
	return &viewerSpecification{}
}
