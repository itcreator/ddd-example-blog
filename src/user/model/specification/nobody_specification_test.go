package specification

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"user/model/entity"
)

type nobodySpecificationSuite struct {
	suite.Suite
}

func (s *nobodySpecificationSuite) TestExecute() {
	spec := NewNobodySpecification()
	user := entity.NewUser("Test user")
	s.False(spec.IsSatisfiedBy(user))
}

func TestNobodySpecificationSuite(t *testing.T) {
	suite.Run(t, new(nobodySpecificationSuite))
}
