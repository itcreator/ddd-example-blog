package specification

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"user/model/entity"
)

type everybodySpecificationSuite struct {
	suite.Suite
}

func (s *everybodySpecificationSuite) TestExecute() {
	spec := NewEverybodySpecification()
	user := entity.NewUser("Test user")
	s.True(spec.IsSatisfiedBy(user))
}

func TestEverybodySpecificationSuite(t *testing.T) {
	suite.Run(t, new(everybodySpecificationSuite))
}
