package actor

import (
	"user/model/specification"
)

func NewViewerSpecification() specification.UserSpecification {
	return specification.NewEverybodySpecification()
}
