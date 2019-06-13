package usecase

import (
	"comment/model/entity"
	"comment/model/repository"
	"comment/model/specification/actor"
	"github.com/google/uuid"
	postEntity "post/model/entity"
	userEntity "user/model/entity"
)

type CreateCommentUC interface {
	Execute(uuid uuid.UUID, post postEntity.Post, body string, user userEntity.User) error
}

type createCommentUC struct {
	commentRepository           repository.Comment
	creatorSpecificationFactory actor.CreatorSpecificationFactory
}

func NewCreateCommentUc(commentRepository repository.Comment, factory actor.CreatorSpecificationFactory) CreateCommentUC {
	return &createCommentUC{
		commentRepository:           commentRepository,
		creatorSpecificationFactory: factory,
	}
}

func (uc *createCommentUC) Execute(uuid uuid.UUID, post postEntity.Post, body string, user userEntity.User) error {
	spec := uc.creatorSpecificationFactory.Create()

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return uc.creatorSpecificationFactory.CreateAccessDeniedError(user)
	}

	//creator.CreateComment
	comment := entity.NewComment(uuid, post, user, body)

	err := uc.commentRepository.Save(comment)

	return err
}
