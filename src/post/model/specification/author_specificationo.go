package specification

import (
	"post/model/entity"
	userEntity "user/model/entity"
	"user/model/specification"
)

type authorSpecification struct {
	post entity.Post
}

//
func (s *authorSpecification) IsSatisfiedBy(user userEntity.User) bool {
	return user.GetUUID() == s.post.GetAuthor().GetUUID()
}

func NewdAuthorSpecification(post entity.Post) specification.UserSpecification {
	return &authorSpecification{post}
}
