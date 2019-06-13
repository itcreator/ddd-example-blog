package actor

import (
	"comment/model/entity"
	"comment/model/specification"
	modelError "post/model/error"
	userEntity "user/model/entity"
	"user/model/permission"
	userSpec "user/model/specification"
)

type EditorSpecificationFactory interface {
	Create(comment entity.Comment) userSpec.UserSpecification
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

func (f *editorSpecificationFactory) Create(comment entity.Comment) userSpec.UserSpecification {
	return userSpec.NewOrSpecification(
		specification.NewAuthorSpecification(comment),
		userSpec.NewGrantedUserSpecification(f.adminChecker),
	)
}

func (f *editorSpecificationFactory) CreateAccessDeniedError(user userEntity.User) error {
	return modelError.NewAccessDeniedError("edit comment", user)
}
