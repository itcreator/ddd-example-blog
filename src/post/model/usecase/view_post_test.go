package usecase

import (
	"github.com/google/uuid"
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

	id := uuid.New()
	post, err := uc.Execute(entity.NewPost(id, user, "test title", "test body"), user)
	s.NoError(err)
	s.NotNil(post)
	s.Equal(id, post.GetUUID())
	s.Equal("test title", post.GetTitle())
	s.Equal("test body", post.GetBody())
}

func TestViewPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(viewPostSuite))
}
