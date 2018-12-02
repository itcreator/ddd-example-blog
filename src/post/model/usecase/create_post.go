package usecase

import (
	"github.com/google/uuid"
	"post/model/entity"
	modelError "post/model/error"
	"post/model/repository"
	"post/model/specification/actor"
	userEntity "user/model/entity"
)

type CreatePostUC interface {
	Execute(uuid uuid.UUID, title, body string, user userEntity.User) error
}

type createPostUC struct {
	postRepository              repository.Post
	creatorSpecificationFactory actor.CreatorSpecificationFactory
}

func NewCreatePostUc(postRepository repository.Post, factory actor.CreatorSpecificationFactory) CreatePostUC {
	return &createPostUC{
		postRepository:              postRepository,
		creatorSpecificationFactory: factory,
	}
}

func (uc *createPostUC) Execute(uuid uuid.UUID, title, body string, user userEntity.User) error {
	spec := uc.creatorSpecificationFactory.Create()

	//if user can't be actor for this UC
	if !spec.IsSatisfiedBy(user) {
		return modelError.NewAccessDeniedError("create post", user)
	}

	//creator.CreatePost
	post := entity.NewPost(uuid, user, title, body)

	err := uc.postRepository.Save(post)

	return err
}
