package mock

import (
	"github.com/google/uuid"
	"post/model/actor"
	"post/model/repository"
	userEntity "user/model/entity"
)

type postCreatorRepository struct {
	storage map[uuid.UUID]actor.PostCreator
}

func NewPostCreatorRepository() repository.PostCreator {
	return &postCreatorRepository{make(map[uuid.UUID]actor.PostCreator)}
}

func (r *postCreatorRepository) Save(postCreator actor.PostCreator) error {
	r.storage[postCreator.GetUser().GetUUID()] = postCreator

	return nil
}

func (r *postCreatorRepository) FindByUser(user userEntity.User) (actor.PostCreator, error) {
	return r.storage[user.GetUUID()], nil
}

func (r *postCreatorRepository) Delete(postCreator actor.PostCreator) error {
	id := postCreator.GetUser().GetUUID()
	if nil != r.storage[id] {
		delete(r.storage, id)
	}

	return nil
}
