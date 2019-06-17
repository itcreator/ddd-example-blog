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
	userMock "user/mock"
	userEntity "user/model/entity"
)

type deleteCommentSuite struct {
	suite.Suite
}

func (s *deleteCommentSuite) TestExecute() {
	commentRepository := mock.NewCommentRepository()

	user := userEntity.NewUser("test user")
	post := postEntity.NewPost(uuid.New(), user, "test", "test")

	pm := userMock.NewAdminPermissionManager()
	uc := NewDeleteCommentUc(commentRepository, actor.NewEditorSpecificationFactory(pm))

	id := uuid.New()
	comment := entity.NewComment(id, post, user, "test body")
	err := commentRepository.Save(comment)
	s.NoError(err)

	err = uc.Execute(comment, user)
	s.NoError(err)

	storedComment, err := commentRepository.Find(id)
	s.NoError(err)
	s.NotNil(storedComment)
	s.False(storedComment.IsEnabled())
}

func (s *deleteCommentSuite) TestExecuteWithOutPermissions() {
	commentRepository := mock.NewCommentRepository()

	author := userEntity.NewUser("test author")

	id := uuid.New()
	post := postEntity.NewPost(id, author, "test", "test")
	comment := entity.NewComment(uuid.New(), post, author, "body")
	err := commentRepository.Save(comment)
	s.NoError(err)

	notAuthor := userEntity.NewUser("test not author")

	pm := userMock.NewAdminPermissionManager()
	specFactory := actor.NewEditorSpecificationFactory(pm)
	uc := NewDeleteCommentUc(commentRepository, specFactory)

	err = uc.Execute(comment, notAuthor)
	s.EqualError(err, commentError.NewAccessDeniedError("edit comment", notAuthor).Error())
}

func TestDeleteCommentHandlerSuite(t *testing.T) {
	suite.Run(t, new(deleteCommentSuite))
}
