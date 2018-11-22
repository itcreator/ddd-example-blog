package usecase

import (
	"post/model/entity"
	modelError "post/model/error"
	"post/model/repository"
	userEntity "user/model/entity"
)

type EditPostUC interface {
	Execute(post entity.Post, user userEntity.User) (entity.Post, error)
}

type editPostUC struct {
	actorRepository repository.PostEditor
	postRepository  repository.Post
}

func NewEditPostUc(actorRepository repository.PostEditor, postRepository repository.Post) EditPostUC {
	return &editPostUC{
		actorRepository: actorRepository,
		postRepository:  postRepository,
	}
}

func (uc *editPostUC) Execute(post entity.Post, user userEntity.User) (entity.Post, error) {
	//load actor by user
	actor, err := uc.actorRepository.LoadForPost(user, post)
	if err != nil {
		//some infrastructure error
		return nil, err
	}

	if actor == nil {
		return nil, modelError.NewAccessDeniedError("edit post", user)
	}

	err = uc.postRepository.Save(post)

	return post, err
}
