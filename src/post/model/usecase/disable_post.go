package usecase

import (
	"post/model/entity"
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
		return uc.editorSpecificationFactory.CreateAccessDeniedError(user)
	}

	post.Disable()

	return uc.postRepository.Save(post)
}
