package usecase

import (
	"comment/model/entity"
	"comment/model/repository"
	"comment/model/specification/actor"
	postEntity "post/model/entity"
	userEntity "user/model/entity"
)

type ListCommentUC interface {
	Execute(post postEntity.Post, start, limit int, user userEntity.User) ([]entity.Comment, int, error)
}

type listCommentUC struct {
	commentRepository          repository.Comment
	viewerSpecificationFactory actor.ViewerSpecificationFactory
}

func NewListCommentUc(commentRepository repository.Comment, factory actor.ViewerSpecificationFactory) ListCommentUC {
	return &listCommentUC{
		commentRepository:          commentRepository,
		viewerSpecificationFactory: factory,
	}
}

func (uc *listCommentUC) Execute(post postEntity.Post, start, limit int, user userEntity.User) ([]entity.Comment, int, error) {
	spec := uc.viewerSpecificationFactory.Create()

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return nil, 0, uc.viewerSpecificationFactory.CreateAccessDeniedError(user)
	}

	comments, err := uc.commentRepository.List(post, start, limit)

	if err != nil {
		return nil, 0, err
	}

	total, err := uc.commentRepository.GetTotal(post)

	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}
