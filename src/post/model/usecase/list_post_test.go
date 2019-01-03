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

type listPostSuite struct {
	suite.Suite
}

func (s *listPostSuite) TestExecute() {
	userRepository := userMock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	uc := NewListPostUc(postRepository, actor.NewViewerSpecificationFactory())

	id := uuid.New()
	author := userEntity.NewUser("test author")
	post := entity.NewPost(id, author, "test", "test")
	err = postRepository.Save(post)

	posts, total, err := uc.Execute(0, 1, user)
	s.NoError(err)
	s.NotNil(posts)
	s.Equal(1, total)

	err = postRepository.Save(entity.NewPost(uuid.New(), author, "test 2", "test 2"))
	err = postRepository.Save(entity.NewPost(uuid.New(), author, "test 3", "test 3"))
	err = postRepository.Save(entity.NewPost(uuid.New(), author, "test 4", "test 4"))
	err = postRepository.Save(entity.NewPost(uuid.New(), author, "test 5", "test 5"))

	posts, total, err = uc.Execute(1, 3, user)
	s.NoError(err)
	s.Equal(3, len(posts))
	s.Equal(5, total)
}

func (s *listPostSuite) TestExecuteEmpty() {
	userRepository := userMock.NewUserRepository()
	postRepository := mock.NewPostRepository()

	user := userEntity.NewUser("test user")
	err := userRepository.Save(user)
	s.NoError(err)

	uc := NewListPostUc(postRepository, actor.NewViewerSpecificationFactory())

	posts, total, err := uc.Execute(0, 1, user)
	s.NoError(err)
	s.Nil(posts)
	s.Equal(0, total)
}

func TestListPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(listPostSuite))
}
