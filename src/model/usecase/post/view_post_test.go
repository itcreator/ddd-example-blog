package post

import (
	"github.com/stretchr/testify/suite"
	"mock"
	"model/entity"
	"testing"
)

type viewPostSuite struct {
	suite.Suite
}

func (s *viewPostSuite) TestExecute() {
	actorRepository := mock.NewPostViewerRepository()
	userRepository := mock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := entity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	uc := NewViewPostUc(actorRepository, postRepository)

	post, err := uc.Execute(entity.NewPost(user, "test", "test"), user)
	s.NoError(err)
	s.NotNil(post)
}

func TestViewPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(viewPostSuite))
}
