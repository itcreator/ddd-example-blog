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
	postRepository repository.Post
}

func NewViewPostUc(postRepository repository.Post) ViewPostUC {
	return &viewPostUC{
		postRepository: postRepository,
	}
}

func (uc *viewPostUC) Execute(post entity.Post, user userEntity.User) (entity.Post, error) {
	spec := actor.NewViewerSpecification()

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return nil, modelError.NewAccessDeniedError("view post", user)
	}

	//there you can run some addition activity. E.g. logging, decorating

	return post, nil
}
