package usecase

import (
	"post/model/entity"
	modelError "post/model/error"
	"post/model/repository"
	"post/model/specification/actor"
	userEntity "user/model/entity"
)

type ViewPostUC interface {
	Execute(post entity.Post, user userEntity.User) (entity.Post, error)
}

type viewPostUC struct {
	postRepository              repository.Post
	creatorSpecificationFactory actor.ViewerSpecificationFactory
}

func NewViewPostUc(postRepository repository.Post, factory actor.ViewerSpecificationFactory) ViewPostUC {
	return &viewPostUC{
		postRepository:              postRepository,
		creatorSpecificationFactory: factory,
	}
}

func (uc *viewPostUC) Execute(post entity.Post, user userEntity.User) (entity.Post, error) {
	spec := uc.creatorSpecificationFactory.Create()

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return nil, modelError.NewAccessDeniedError("view post", user)
	}

	//there you can run some addition activity. E.g. logging, decorating

	return post, nil
}
