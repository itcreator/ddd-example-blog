package usecase

import (
	"post/model/entity"
	modelError "post/model/error"
	"post/model/repository"
	"post/model/specification/actor"
	userEntity "user/model/entity"
)

type EditPostUC interface {
	Execute(title, body string, post entity.Post, user userEntity.User) error
}

type editPostUC struct {
	postRepository             repository.Post
	editorSpecificationFactory actor.EditorSpecificationFactory
}

func NewEditPostUc(postRepository repository.Post, factory actor.EditorSpecificationFactory) EditPostUC {
	return &editPostUC{
		postRepository:             postRepository,
		editorSpecificationFactory: factory,
	}
}

func (uc *editPostUC) Execute(title, body string, post entity.Post, user userEntity.User) error {
	spec := uc.editorSpecificationFactory.Create(post)

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return modelError.NewAccessDeniedError("edit post", user)
	}

	post.Update(title, body)

	err := uc.postRepository.Save(post)

	return err
}
