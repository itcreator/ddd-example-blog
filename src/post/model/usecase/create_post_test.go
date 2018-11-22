package usecase

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"post/mock"
	"post/model/actor"
	error2 "post/model/error"
	"post/model/repository"
	"testing"
	userEntity "user/model/entity"
)

type createPostSuite struct {
	suite.Suite
}

func (s *createPostSuite) TestExecute() {
	actorRepository := mock.NewPostCreatorRepository()
	userRepository := mock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	postCreator := actor.NewPostCreator(user)
	err = actorRepository.Save(postCreator)
	s.NoError(err)

	uc := NewCreatePostUc(actorRepository, postRepository)

	err = uc.Execute("test title", "testBody", user)
	s.NoError(err)
}

func (s *createPostSuite) TestExecuteWithOutPermissions() {
	actorRepository := mock.NewPostCreatorRepository()
	userRepository := mock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	uc := NewCreatePostUc(actorRepository, postRepository)

	err = uc.Execute("test title", "testBody", user)
	s.EqualError(err, error2.NewAccessDeniedError("create post", user).Error())
}

func (s *createPostSuite) TestExecuteHandlesInfrastructureError() {
	brokenActorRepository := NewMockBrokenRepository()
	userRepository := mock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	uc := NewCreatePostUc(brokenActorRepository, postRepository)

	err = uc.Execute("test title", "testBody", user)
	s.EqualError(err, NewMockInfrastructureError().Error())
}

func TestCreatePostHandlerSuite(t *testing.T) {
	suite.Run(t, new(createPostSuite))
}

type mockInfrastructureError struct {
	actorName string
	user      userEntity.User
}

func NewMockInfrastructureError() error {
	return &mockInfrastructureError{}
}

func (e *mockInfrastructureError) Error() string {
	return fmt.Sprintf("test error")
}

type mockBrokenRepository struct {
}

func NewMockBrokenRepository() repository.PostCreator {
	return &mockBrokenRepository{}
}

func (r *mockBrokenRepository) Save(postCreator actor.PostCreator) error {
	return NewMockInfrastructureError()
}

func (r *mockBrokenRepository) FindByUser(user userEntity.User) (actor.PostCreator, error) {
	return nil, NewMockInfrastructureError()
}

func (r *mockBrokenRepository) Delete(postCreator actor.PostCreator) error {
	return NewMockInfrastructureError()
}
