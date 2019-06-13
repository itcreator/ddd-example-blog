package specification

import (
	"comment/model/entity"
	userEntity "user/model/entity"
	"user/model/specification"
)

type authorSpecification struct {
	comment entity.Comment
}

//
func (s *authorSpecification) IsSatisfiedBy(user userEntity.User) bool {
	return user.GetUUID() == s.comment.GetAuthor().GetUUID()
}

func NewAuthorSpecification(comment entity.Comment) specification.UserSpecification {
	return &authorSpecification{comment}
}
