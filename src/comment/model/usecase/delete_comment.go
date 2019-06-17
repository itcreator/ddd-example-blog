package usecase

import (
	"comment/model/entity"
	"comment/model/repository"
	"comment/model/specification/actor"
	userEntity "user/model/entity"
)

type DeleteCommentUC interface {
	Execute(comment entity.Comment, user userEntity.User) error
}

type deleteCommentUC struct {
	commentRepository          repository.Comment
	editorSpecificationFactory actor.EditorSpecificationFactory
}

func NewDeleteCommentUc(commentRepository repository.Comment, factory actor.EditorSpecificationFactory) DeleteCommentUC {
	return &deleteCommentUC{
		commentRepository:          commentRepository,
		editorSpecificationFactory: factory,
	}
}

func (uc *deleteCommentUC) Execute(comment entity.Comment, user userEntity.User) error {
	spec := uc.editorSpecificationFactory.Create(comment)

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return uc.editorSpecificationFactory.CreateAccessDeniedError(user)
	}

	comment.Delete()
	err := uc.commentRepository.Save(comment)

	return err
}
