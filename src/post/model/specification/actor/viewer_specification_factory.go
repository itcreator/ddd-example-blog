package actor

import (
	"user/model/specification"
)

type ViewerSpecificationFactory interface {
	Create() specification.UserSpecification
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
