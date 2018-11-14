package post

import (
	"model/entity"
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

func (c *createPostUC) Execute(title, body string, user entity.User) error {
	//load actor by user
	actor, err := c.actorRepository.FindByUser(user)
	if err != nil {
		return err
	}

	//creator.CreatePost
	post := actor.CreatePost(title, body)
	_ = post
	//savePost
	//return error or nil

	return nil
}
