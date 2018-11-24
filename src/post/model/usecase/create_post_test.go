package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"post/mock"
	"post/model/entity"
	"post/model/repository"
	"testing"
	userEntity "user/model/entity"
)

type createPostSuite struct {
	suite.Suite
}

func (s *createPostSuite) TestExecute() {
	userRepository := mock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	uc := NewCreatePostUc(postRepository)

	err = uc.Execute("test title", "testBody", user)
	s.NoError(err)
}

func (s *createPostSuite) TestExecuteHandlesInfrastructureError() {
	userRepository := mock.NewUserRepository()
	postRepository := NewMockBrokenPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	uc := NewCreatePostUc(postRepository)

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

type mockBrokenPostRepository struct {
}

func NewMockBrokenPostRepository() repository.Post {
	return &mockBrokenPostRepository{}
}

func (r *mockBrokenPostRepository) Save(post entity.Post) error {
	return NewMockInfrastructureError()
}

func (r *mockBrokenPostRepository) Find(uuid uuid.UUID) (entity.Post, error) {
	return nil, NewMockInfrastructureError()
}
