package usecase

import (
	"comment/mock"
	"comment/model/entity"
	commentError "comment/model/error"
	"comment/model/specification/actor"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	postEntity "post/model/entity"
	"testing"
	userEntity "user/model/entity"
	"user/model/specification"
)

type listPostSuite struct {
	suite.Suite
}

type notViewerSpecificationFactory struct {
}

func (f *notViewerSpecificationFactory) Create() specification.UserSpecification {
	return specification.NewNobodySpecification()
}

func (f *notViewerSpecificationFactory) CreateAccessDeniedError(user userEntity.User) error {
	return commentError.NewAccessDeniedError("view comment", user)
}

func (s *listPostSuite) TestExecute() {
	commentRepository := mock.NewCommentRepository()

	user := userEntity.NewUser("test user")

	uc := NewListCommentUc(commentRepository, actor.NewViewerSpecificationFactory())

	id := uuid.New()
	author := userEntity.NewUser("test author")
	post := postEntity.NewPost(uuid.New(), author, "test", "test")
	comment := entity.NewComment(id, post, author, "test")
	err := commentRepository.Save(comment)

	comments, total, err := uc.Execute(post, 0, 1, user)
	s.NoError(err)
	s.NotNil(comments)
	s.Equal(1, total)

	err = commentRepository.Save(entity.NewComment(uuid.New(), post, author, "test 2"))
	err = commentRepository.Save(entity.NewComment(uuid.New(), post, author, "test 3"))
	err = commentRepository.Save(entity.NewComment(uuid.New(), post, author, "test 4"))
	err = commentRepository.Save(entity.NewComment(uuid.New(), post, author, "test 5"))

	comments, total, err = uc.Execute(post, 1, 3, user)
	s.NoError(err)
	s.Equal(3, len(comments))
	s.Equal(5, total)
	s.Equal(comments[0].GetPost().GetUUID(), post.GetUUID())
}

func (s *listPostSuite) TestExecuteEmpty() {
	commentRepository := mock.NewCommentRepository()

	author := userEntity.NewUser("test author")
	post := postEntity.NewPost(uuid.New(), author, "test", "test")

	uc := NewListCommentUc(commentRepository, actor.NewViewerSpecificationFactory())

	comments, total, err := uc.Execute(post, 0, 1, author)
	s.NoError(err)
	s.Nil(comments)
	s.Equal(0, total)
}

func (s *listPostSuite) TestExecuteWithoutPermissions() {
	commentRepository := mock.NewCommentRepository()

	author := userEntity.NewUser("test author")
	post := postEntity.NewPost(uuid.New(), author, "test", "test")

	uc := NewListCommentUc(commentRepository, new(notViewerSpecificationFactory))

	comments, total, err := uc.Execute(post, 0, 1, author)
	s.EqualError(err, commentError.NewAccessDeniedError("view comment", author).Error())
	s.Nil(comments)
	s.Equal(0, total)
}

func TestListPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(listPostSuite))
}
