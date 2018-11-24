package actor

import (
	"post/model/entity"
	"post/model/specification"
	userSpecification "user/model/specification"
)

func NewEditorSpecification(post entity.Post) userSpecification.UserSpecification {
	return userSpecification.NewOrSpecification(
		specification.NewAuthorSpecification(post),
		userSpecification.NewAdminSpecification(),
	)
}
