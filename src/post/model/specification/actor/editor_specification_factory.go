package actor

import (
	"post/model/entity"
	modelError "post/model/error"
	"post/model/specification"
	userEntity "user/model/entity"
	"user/model/permission"
	userSpec "user/model/specification"
)

type EditorSpecificationFactory interface {
	Create(post entity.Post) userSpec.UserSpecification
	CreateAccessDeniedError(user userEntity.User) error
}

type editorSpecificationFactory struct {
	adminChecker permission.Checker
}

//NewEditorSpecificationFactory is a constructor
//Be sure that adminChecker which implement permission.Checker really check admin permission
func NewEditorSpecificationFactory(adminChecker permission.Checker) EditorSpecificationFactory {
	return &editorSpecificationFactory{adminChecker}
}

func (f *editorSpecificationFactory) Create(post entity.Post) userSpec.UserSpecification {
	return userSpec.NewOrSpecification(
		specification.NewAuthorSpecification(post),
		userSpec.NewGrantedUserSpecification(f.adminChecker),
	)
}

func (f *editorSpecificationFactory) CreateAccessDeniedError(user userEntity.User) error {
	return modelError.NewAccessDeniedError("edit post", user)
}
