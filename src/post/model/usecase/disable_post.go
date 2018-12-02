package usecase

import (
	"post/model/entity"
	modelError "post/model/error"
	"post/model/repository"
	"post/model/specification/actor"
	userEntity "user/model/entity"
)

type DisablePostUC interface {
	Execute(post entity.Post, user userEntity.User) error
}

type disablePostUC struct {
	postRepository             repository.Post
	editorSpecificationFactory actor.EditorSpecificationFactory
}

func NewDisablePostUc(postRepository repository.Post, factory actor.EditorSpecificationFactory) DisablePostUC {
	return &disablePostUC{
		postRepository:             postRepository,
		editorSpecificationFactory: factory,
	}
}

func (uc *disablePostUC) Execute(post entity.Post, user userEntity.User) error {
	spec := uc.editorSpecificationFactory.Create(post)

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		//todo: create error in specification factory
		return modelError.NewAccessDeniedError("edit post", user)
	}

	post.Disable()

	return uc.postRepository.Save(post)
}
