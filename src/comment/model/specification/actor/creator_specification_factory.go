package actor

import (
	modelError "comment/model/error"
	"user/model"
	"user/model/entity"
	"user/model/specification"
)

type CreatorSpecificationFactory interface {
	Create() specification.UserSpecification
	CreateAccessDeniedError(user entity.User) error
}

type creatorSpecificationFactory struct {
	authenticator model.Authenticator
}

//every authenticated user can create new comment
func NewCreatorSpecificationFactory(authenticator model.Authenticator) CreatorSpecificationFactory {
	return &creatorSpecificationFactory{authenticator}
}

func (f *creatorSpecificationFactory) Create() specification.UserSpecification {
	return specification.NewAuthenticatedUserSpecification(f.authenticator)
}

func (f *creatorSpecificationFactory) CreateAccessDeniedError(user entity.User) error {
	return modelError.NewAccessDeniedError("create comment", user)
}
