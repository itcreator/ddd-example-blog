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

type enablePostSuite struct {
	suite.Suite
}

func (s *enablePostSuite) TestExecute() {
	userRepository := userMock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	pm := userMock.NewAdminPermissionManager()
	uc := NewEnablePostUc(postRepository, actor.NewEditorSpecificationFactory(pm))

	id := uuid.New()
	post := entity.NewPost(id, user, "test title", "test body")
	post.Disable()
	err = postRepository.Save(post)
	s.NoError(err)

	err = uc.Execute(post, user)
	s.NoError(err)

	storedPost, err := postRepository.Find(id)
	s.NoError(err)
	s.NotNil(storedPost)
	s.True(storedPost.IsEnabled())
}

func (s *enablePostSuite) TestExecuteWithOutPermissions() {
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
	uc := NewEnablePostUc(postRepository, specFactory)

	err = uc.Execute(post, notAuthor)
	s.EqualError(err, error2.NewAccessDeniedError("edit post", notAuthor).Error())
}

func TestEnablePostHandlerSuite(t *testing.T) {
	suite.Run(t, new(enablePostSuite))
}
