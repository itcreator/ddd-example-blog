package usecase

import (
	"comment/mock"
	error2 "comment/model/error"
	"comment/model/specification/actor"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	postMock "post/mock"
	postEntity "post/model/entity"
	"testing"
	userMock "user/mock"
	userEntity "user/model/entity"
)

type createCommentSuite struct {
	suite.Suite
}

func (s *createCommentSuite) authorExists() userEntity.User {
	author := userEntity.NewUser("test user")
	_ = userMock.NewUserRepository().Save(author)

	return author
}

func (s *createCommentSuite) postExists() postEntity.Post {
	post := postEntity.NewPost(uuid.New(), s.authorExists(), "test title", "test body")
	_ = postMock.NewPostRepository().Save(post)

	return post
}

func (s *createCommentSuite) TestExecute() {
	commentRepository := mock.NewCommentRepository()

	user := s.authorExists()
	post := s.postExists()

	//todo: mock creatorFactory
	authenticator := userMock.NewAuthenticator()
	_ = authenticator.Authenticate(user)
	creatorFactory := actor.NewCreatorSpecificationFactory(authenticator)

	uc := NewCreateCommentUc(commentRepository, creatorFactory)

	id := uuid.New()

	err := uc.Execute(id, post, "test comment", user)

	comment, err := commentRepository.Find(id)
	s.NoError(err)

	s.Equal(id, comment.GetUUID())
	s.Equal("test comment", comment.GetBody())
	s.Equal(post, comment.GetPost())
	s.Equal(user, comment.GetAuthor())
}

func (s *createCommentSuite) TestExecuteWithOutPermissions() {
	user := s.authorExists()
	post := s.postExists()

	commentRepository := mock.NewCommentRepository()

	//todo: mock creatorFactory
	authenticator := userMock.NewAuthenticator()
	creatorFactory := actor.NewCreatorSpecificationFactory(authenticator)

	id := uuid.New()
	uc := NewCreateCommentUc(commentRepository, creatorFactory)
	err := uc.Execute(id, post, "test comment", user)

	s.EqualError(err, error2.NewAccessDeniedError("create comment", user).Error())
}

func TestCreateCommentHandlerSuite(t *testing.T) {
	suite.Run(t, new(createCommentSuite))
}
