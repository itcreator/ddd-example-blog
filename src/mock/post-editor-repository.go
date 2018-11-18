package mock

import (
	"github.com/google/uuid"
	"model/actor"
	"model/entity"
	"model/repository"
)

type postEditorRepository struct {
	storage map[uuid.UUID]actor.PostEditor
}

func NewPostEditorRepository() repository.PostEditor {
	return &postEditorRepository{make(map[uuid.UUID]actor.PostEditor)}
}

func (r *postEditorRepository) Save(postEditor actor.PostEditor) error {
	r.storage[postEditor.GetUser().GetUUID()] = postEditor

	return nil
}

func (r *postEditorRepository) FindByUser(user entity.User) (actor.PostEditor, error) {
	return r.storage[user.GetUUID()], nil
}

func (r *postEditorRepository) Delete(postEditor actor.PostEditor) error {
	id := postEditor.GetUser().GetUUID()
	if nil != r.storage[id] {
		delete(r.storage, id)
	}

	return nil
}
