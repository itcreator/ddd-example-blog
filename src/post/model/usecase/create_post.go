package usecase

import (
	modelError "post/model/error"
	"post/model/repository"
	userEntity "user/model/entity"
)

type CreatePostUC interface {
	Execute(title, body string, user userEntity.User) error
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

func (uc *createPostUC) Execute(title, body string, user userEntity.User) error {
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
