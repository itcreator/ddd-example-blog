package usecase

import (
	"comment/mock"
	"comment/model/entity"
	commentError "comment/model/error"
	"comment/model/specification/actor"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	postMock "post/mock"
	postEntity "post/model/entity"
	"testing"
	userMock "user/mock"
	userEntity "user/model/entity"
)

type editCommentSuite struct {
	suite.Suite
}

func (s *editCommentSuite) authorExists() userEntity.User {
	author := userEntity.NewUser("test user")
	_ = userMock.NewUserRepository().Save(author)

	return author
}

func (s *editCommentSuite) postExists() postEntity.Post {
	post := postEntity.NewPost(uuid.New(), s.authorExists(), "test title", "test body")
	_ = postMock.NewPostRepository().Save(post)

	return post
}

func (s *editCommentSuite) TestExecute() {
	commentRepository := mock.NewCommentRepository()

	post := s.postExists()
	author := s.authorExists()

	comment := entity.NewComment(uuid.New(), post, author, "test body")
	err := commentRepository.Save(comment)
	s.NoError(err)

	//todo: mock editorFactory
	pm := userMock.NewAdminPermissionManager()
	editorFactory := actor.NewEditorSpecificationFactory(pm)

	uc := NewEditCommentUc(commentRepository, editorFactory)

	err = uc.Execute(comment, "test comment", author)

	comment, err = commentRepository.Find(comment.GetUUID())

	s.NoError(err)

	s.Equal("test comment", comment.GetBody())
	s.Equal(post, comment.GetPost())
	s.Equal(author, comment.GetAuthor())
}

func (s *editCommentSuite) TestExecuteWithOutPermissions() {
	commentRepository := mock.NewCommentRepository()

	author := s.authorExists()
	user := s.authorExists()
	post := s.postExists()

	comment := entity.NewComment(uuid.New(), post, author, "test body")

	editorFactory := actor.NewEditorSpecificationFactory(userMock.NewAdminPermissionManager())

	uc := NewEditCommentUc(commentRepository, editorFactory)
	err := uc.Execute(comment, "test comment", user)

	s.EqualError(err, commentError.NewAccessDeniedError("edit comment", author).Error())
}

func TestEditCommentHandlerSuite(t *testing.T) {
	suite.Run(t, new(editCommentSuite))
}
