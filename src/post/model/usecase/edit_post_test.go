package usecase

import (
	"github.com/stretchr/testify/suite"
	"post/mock"
	"post/model/entity"
	error2 "post/model/error"
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

func (s *editPostSuite) TestExecuteWithOutPermissions() {
	userRepository := mock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	author := userEntity.NewUser("test author")
	err := userRepository.Save(author)
	s.NoError(err)

	post := entity.NewPost(author, "test", "test")
	err = postRepository.Save(post)
	s.NoError(err)

	notAuthor := userEntity.NewUser("test not author")

	uc := NewEditPostUc(postRepository)

	err = uc.Execute("new title", "new body", post, notAuthor)
	s.EqualError(err, error2.NewAccessDeniedError("edit post", notAuthor).Error())
}

func TestEditPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(editPostSuite))
}
