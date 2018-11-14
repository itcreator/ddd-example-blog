package post

import (
	"github.com/stretchr/testify/suite"
	"mock"
	"model/actor"
	"model/entity"
	"testing"
)

type createPostSuite struct {
	suite.Suite
}

func (s *createPostSuite) TestExecute() {
	actorRepository := mock.NewPostCreatorRepository()
	userRepository := mock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := entity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	postCreator := actor.NewPostCreator(user)
	err = actorRepository.Save(postCreator)
	s.NoError(err)

	uc := NewCreatePostUc(actorRepository, postRepository)

	err = uc.Execute("test title", "testBody", user)
	s.NoError(err)
}

func TestCreatePostHandlerSuite(t *testing.T) {
	suite.Run(t, new(createPostSuite))
}
