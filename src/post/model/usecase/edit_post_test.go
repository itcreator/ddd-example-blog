package usecase

import (
	"github.com/stretchr/testify/suite"
	"post/mock"
	"post/model/entity"
	"testing"
	userEntity "user/model/entity"
)

type editPostSuite struct {
	suite.Suite
}

func (s *editPostSuite) TestExecute() {
	userRepository := mock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	post := entity.NewPost(user, "test title", "test body")

	uc := NewEditPostUc(postRepository)

	err = uc.Execute("test title2", "test body2", post, user)
	s.NoError(err)

	s.Equal("test title2", post.GetTitle())
	s.Equal("test body2", post.GetBody())
}

func TestEditPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(editPostSuite))
}
