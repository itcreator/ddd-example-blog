package usecase

import (
	"github.com/stretchr/testify/suite"
	"post/mock"
	"post/model/entity"
	"post/model/specification/actor"
	"testing"
	userMock "user/mock"
	userEntity "user/model/entity"
)

type viewPostSuite struct {
	suite.Suite
}

func (s *viewPostSuite) TestExecute() {
	userRepository := userMock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	uc := NewViewPostUc(postRepository, actor.NewViewerSpecificationFactory())

	post, err := uc.Execute(entity.NewPost(user, "test", "test"), user)
	s.NoError(err)
	s.NotNil(post)
}

func TestViewPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(viewPostSuite))
}
