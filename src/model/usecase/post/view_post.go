package post

import (
	"model/entity"
	modelError "model/error"
	"model/repository"
)

type ViewPostUC interface {
	Execute(post entity.Post, user entity.User) (entity.Post, error)
}

type viewPostUC struct {
	actorRepository repository.PostViewer
	postRepository  repository.Post
}

func NewViewPostUc(actorRepository repository.PostViewer, postRepository repository.Post) ViewPostUC {
	return &viewPostUC{
		actorRepository: actorRepository,
		postRepository:  postRepository,
	}
}

func (uc *viewPostUC) Execute(post entity.Post, user entity.User) (entity.Post, error) {
	//load actor by user
	actor, err := uc.actorRepository.LoadForPost(user, post)
	if err != nil {
		//some infrastructure error
		return nil, err
	}

	if actor == nil {
		return nil, modelError.NewAccessDeniedError("view post", user)
	}

	err = uc.postRepository.Save(post)

	return post, err
}
