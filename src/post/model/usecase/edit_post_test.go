package usecase

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"post/mock"
	"post/model/entity"
	error2 "post/model/error"
	"post/model/specification/actor"
	"testing"
	userMock "user/mock"
	userEntity "user/model/entity"
)

type editPostSuite struct {
	suite.Suite
}

func (s *editPostSuite) TestExecute() {
	userRepository := userMock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	id := uuid.New()
	post := entity.NewPost(id, user, "test title", "test body")

	pm := userMock.NewAdminPermissionManager()
	specFactory := actor.NewEditorSpecificationFactory(pm)
	uc := NewEditPostUc(postRepository, specFactory)

	err = uc.Execute("test title2", "test body2", post, user)
	s.NoError(err)

	s.Equal("test title2", post.GetTitle())
	s.Equal("test body2", post.GetBody())
}

func (s *editPostSuite) TestExecuteWithOutPermissions() {
	userRepository := userMock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	author := userEntity.NewUser("test author")
	err := userRepository.Save(author)
	s.NoError(err)

	id := uuid.New()
	post := entity.NewPost(id, author, "test", "test")
	err = postRepository.Save(post)
	s.NoError(err)

	notAuthor := userEntity.NewUser("test not author")

	pm := userMock.NewAdminPermissionManager()
	specFactory := actor.NewEditorSpecificationFactory(pm)
	uc := NewEditPostUc(postRepository, specFactory)

	err = uc.Execute("new title", "new body", post, notAuthor)
	s.EqualError(err, error2.NewAccessDeniedError("edit post", notAuthor).Error())
}

func TestEditPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(editPostSuite))
}
