package usecase

import (
	"post/model/entity"
	"post/model/repository"
	"post/model/specification/actor"
	userEntity "user/model/entity"
)

type ListPostUC interface {
	Execute(start, limit int, user userEntity.User) ([]entity.Post, int, error)
}

type listPostUC struct {
	postRepository             repository.Post
	viewerSpecificationFactory actor.ViewerSpecificationFactory
}

func NewListPostUc(postRepository repository.Post, factory actor.ViewerSpecificationFactory) ListPostUC {
	return &listPostUC{
		postRepository:             postRepository,
		viewerSpecificationFactory: factory,
	}
}

func (uc *listPostUC) Execute(start, limit int, user userEntity.User) ([]entity.Post, int, error) {
	spec := uc.viewerSpecificationFactory.Create()

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return nil, 0, uc.viewerSpecificationFactory.CreateAccessDeniedError(user)
	}

	posts, err := uc.postRepository.List(start, limit)

	if err != nil {
		return nil, 0, err
	}

	total, err := uc.postRepository.GetTotal()

	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}
