package mock

import (
	"github.com/google/uuid"
	"post/model/actor"
	"post/model/entity"
	"post/model/repository"
	userEntity "user/model/entity"
)

type postEditorRepository struct {
	storage map[uuid.UUID]actor.PostEditor
	//adminChecker
	//savedRoleLoader
}

func NewPostEditorRepository() repository.PostEditor {
	return &postEditorRepository{make(map[uuid.UUID]actor.PostEditor)}
}

func (r *postEditorRepository) LoadForPost(user userEntity.User, post entity.Post) (actor.PostEditor, error) {
	//check if user is author fn(user, post)
	//check if user has permanent role postViewer fn(user)
	//check if user is admin fn(user)
	//TODO: add possibility to create private posts and to grant users
	return actor.NewPostEditor(user), nil
}

func NewAdmSpec() {

}

type admFactory struct {
}

func (f *admFactory) getSpecification(user userEntity.User) {
	//return NewAdmSpec(admRep)
	//return  f.adminRepository.isAdmin(user)
}
