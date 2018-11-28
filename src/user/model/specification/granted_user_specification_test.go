package specification

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"user/mock"
	"user/model/entity"
)

type adminSpecificationSuite struct {
	suite.Suite
}

func (s *adminSpecificationSuite) TestExecute() {
	pm := mock.NewAdminPermissionManager()
	spec := NewGrantedUserSpecification(pm)

	user := entity.NewUser("Test user")
	s.False(spec.IsSatisfiedBy(user))

	admin := entity.NewUser("Test admin")
	_ = pm.GrantPermission(admin)
	s.True(spec.IsSatisfiedBy(admin))

	_ = pm.RevokePermission(admin)
	s.False(spec.IsSatisfiedBy(admin))
}

func TestAdminSpecificationSuite(t *testing.T) {
	suite.Run(t, new(adminSpecificationSuite))
}
