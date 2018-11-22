package actor

import (
	"user/model/specification"
)

//every authenticated user can create new post
func NewCreatorSpecification() specification.UserSpecification {
	return specification.NewAuthenticatedUserSpecification()
}
