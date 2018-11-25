package specification

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"user/mock"
	"user/model/entity"
)

type authenticatedUserSpecificationSuite struct {
	suite.Suite
}

func (s *authenticatedUserSpecificationSuite) TestExecute() {
	auth := mock.NewAuthenticator()
	spec := NewAuthenticatedUserSpecification(auth)
	user := entity.NewUser("Test user")
	s.False(spec.IsSatisfiedBy(user))

	_ = auth.Authenticate(user)
	s.True(spec.IsSatisfiedBy(user))
}

func TestAuthenticatedUserSpecificationSuite(t *testing.T) {
	suite.Run(t, new(authenticatedUserSpecificationSuite))
}
