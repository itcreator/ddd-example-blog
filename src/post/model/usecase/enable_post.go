package usecase

import (
	"post/model/entity"
	modelError "post/model/error"
	"post/model/repository"
	"post/model/specification/actor"
	userEntity "user/model/entity"
)

type EnablePostUC interface {
	Execute(post entity.Post, user userEntity.User) error
}

type enablePostUC struct {
	postRepository             repository.Post
	editorSpecificationFactory actor.EditorSpecificationFactory
}

func NewEnablePostUc(postRepository repository.Post, factory actor.EditorSpecificationFactory) EnablePostUC {
	return &enablePostUC{
		postRepository:             postRepository,
		editorSpecificationFactory: factory,
	}
}

func (uc *enablePostUC) Execute(post entity.Post, user userEntity.User) error {
	spec := uc.editorSpecificationFactory.Create(post)

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		//todo: create error in specification factory
		return modelError.NewAccessDeniedError("edit post", user)
	}

	post.Enable()

	return uc.postRepository.Save(post)
}
