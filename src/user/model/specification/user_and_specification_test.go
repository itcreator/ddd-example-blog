package specification

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"user/model/entity"
)

type andSpecificationSuite struct {
	suite.Suite
}

func (s *andSpecificationSuite) TestExecute() {
	nb := NewNobodySpecification()
	eb := NewEverybodySpecification()
	user := entity.NewUser("Test user")

	spec := NewAndSpecification(nb, eb)
	s.False(spec.IsSatisfiedBy(user))

	spec = NewAndSpecification(nb, nb)
	s.False(spec.IsSatisfiedBy(user))

	spec = NewAndSpecification(eb, eb)
	s.True(spec.IsSatisfiedBy(user))
}

func TestAndSpecificationSuite(t *testing.T) {
	suite.Run(t, new(andSpecificationSuite))
}
