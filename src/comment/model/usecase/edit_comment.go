package usecase

import (
	"comment/model/entity"
	"comment/model/repository"
	"comment/model/specification/actor"
	userEntity "user/model/entity"
)

type EditCommentUC interface {
	Execute(comment entity.Comment, body string, user userEntity.User) error
}

type editCommentUC struct {
	commentRepository          repository.Comment
	editorSpecificationFactory actor.EditorSpecificationFactory
}

func NewEditCommentUc(commentRepository repository.Comment, factory actor.EditorSpecificationFactory) EditCommentUC {
	return &editCommentUC{
		commentRepository:          commentRepository,
		editorSpecificationFactory: factory,
	}
}

func (uc *editCommentUC) Execute(comment entity.Comment, body string, user userEntity.User) error {
	spec := uc.editorSpecificationFactory.Create(comment)

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return uc.editorSpecificationFactory.CreateAccessDeniedError(user)
	}

	comment.Update(body)
	err := uc.commentRepository.Save(comment)

	return err
}
