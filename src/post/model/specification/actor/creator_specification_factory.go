package actor

import (
	"user/model"
	"user/model/specification"
)

type CreatorSpecificationFactory interface {
	Create() specification.UserSpecification
}

type creatorSpecificationFactory struct {
	authenticator model.Authenticator
}

//every authenticated user can create new post
func NewCreatorSpecificationFactory(authenticator model.Authenticator) CreatorSpecificationFactory {
	return &creatorSpecificationFactory{authenticator}
}

func (f *creatorSpecificationFactory) Create() specification.UserSpecification {
	return specification.NewAuthenticatedUserSpecification(f.authenticator)
}
