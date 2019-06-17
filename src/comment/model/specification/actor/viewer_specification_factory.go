package actor

import (
	modelError "comment/model/error"
	"user/model/entity"
	"user/model/specification"
)

type ViewerSpecificationFactory interface {
	Create() specification.UserSpecification
	CreateAccessDeniedError(user entity.User) error
}

type viewerSpecificationFactory struct {
}

//everybody can view post
func NewViewerSpecificationFactory() ViewerSpecificationFactory {
	return &viewerSpecificationFactory{}
}

func (f *viewerSpecificationFactory) Create() specification.UserSpecification {
	return specification.NewEverybodySpecification()
}

func (f *viewerSpecificationFactory) CreateAccessDeniedError(user entity.User) error {
	return modelError.NewAccessDeniedError("view comment", user)
}
