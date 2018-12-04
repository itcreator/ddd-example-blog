package usecase

import (
	"post/model/entity"
	"post/model/repository"
	"post/model/specification/actor"
	userEntity "user/model/entity"
)

type ViewPostUC interface {
	Execute(post entity.Post, user userEntity.User) (entity.Post, error)
}

type viewPostUC struct {
	postRepository             repository.Post
	viewerSpecificationFactory actor.ViewerSpecificationFactory
}

func NewViewPostUc(postRepository repository.Post, factory actor.ViewerSpecificationFactory) ViewPostUC {
	return &viewPostUC{
		postRepository:             postRepository,
		viewerSpecificationFactory: factory,
	}
}

func (uc *viewPostUC) Execute(post entity.Post, user userEntity.User) (entity.Post, error) {
	spec := uc.viewerSpecificationFactory.Create()

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return nil, uc.viewerSpecificationFactory.CreateAccessDeniedError(user)
	}

	//there you can run some addition activity. E.g. logging, decorating

	return post, nil
}
