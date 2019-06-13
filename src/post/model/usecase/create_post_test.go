package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"post/mock"
	"post/model/entity"
	error2 "post/model/error"
	"post/model/repository"
	"post/model/specification/actor"
	"testing"
	userMock "user/mock"
	userEntity "user/model/entity"
)

type createPostSuite struct {
	suite.Suite
}

func (s *createPostSuite) TestExecute() {
	userRepository := userMock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	authenticator := userMock.NewAuthenticator()
	_ = authenticator.Authenticate(user)
	creatorFactory := actor.NewCreatorSpecificationFactory(authenticator)
	uc := NewCreatePostUc(postRepository, creatorFactory)

	id := uuid.New()
	err = uc.Execute(id, "test title", "test body", user)
	s.NoError(err)

	post, err := postRepository.Find(id)
	s.NoError(err)

	s.Equal(id, post.GetUUID())
	s.Equal("test title", post.GetTitle())
	s.Equal("test body", post.GetBody())
}

func (s *createPostSuite) TestExecuteWithOutPermissions() {
	userRepository := userMock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	authenticator := userMock.NewAuthenticator()
	creatorFactory := actor.NewCreatorSpecificationFactory(authenticator)
	uc := NewCreatePostUc(postRepository, creatorFactory)

	id := uuid.New()
	err = uc.Execute(id, "test title", "testBody", user)
	s.EqualError(err, error2.NewAccessDeniedError("create post", user).Error())
}

func (s *createPostSuite) TestExecuteHandlesInfrastructureError() {
	userRepository := userMock.NewUserRepository()
	postRepository := NewMockBrokenPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	authenticator := userMock.NewAuthenticator()
	_ = authenticator.Authenticate(user)
	creatorFactory := actor.NewCreatorSpecificationFactory(authenticator)
	uc := NewCreatePostUc(postRepository, creatorFactory)

	id := uuid.New()
	err = uc.Execute(id, "test title", "testBody", user)
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

func (r *mockBrokenPostRepository) List(start, limit int) ([]entity.Post, error) {
	return nil, nil
}

func (r *mockBrokenPostRepository) GetTotal() (int, error) {
	return 0, nil
}
