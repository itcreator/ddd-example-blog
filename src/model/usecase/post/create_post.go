package post

import (
	"model/entity"
	modelError "model/error"
	"model/repository"
)

type CreatePostUC interface {
	Execute(title, body string, user entity.User) error
}

type createPostUC struct {
	actorRepository repository.PostCreator
	postRepository  repository.Post
}

func NewCreatePostUc(actorRepository repository.PostCreator, postRepository repository.Post) CreatePostUC {
	return &createPostUC{
		actorRepository: actorRepository,
		postRepository:  postRepository,
	}
}

func (uc *createPostUC) Execute(title, body string, user entity.User) error {
	//load actor by user
	actor, err := uc.actorRepository.FindByUser(user)
	if err != nil {
		//some infrastructure error
		return err
	}

	if actor == nil {
		return modelError.NewAccessDeniedError("create post", user)
	}

	//creator.CreatePost
	post := actor.CreatePost(title, body)

	err = uc.postRepository.Save(post)

	return err
}
