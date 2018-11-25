package specification

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"user/model/entity"
)

type orSpecificationSuite struct {
	suite.Suite
}

func (s *orSpecificationSuite) TestExecute() {
	nb := NewNobodySpecification()
	eb := NewEverybodySpecification()
	user := entity.NewUser("Test user")

	spec := NewOrSpecification(nb, eb)
	s.True(spec.IsSatisfiedBy(user))

	spec = NewOrSpecification(nb, nb)
	s.False(spec.IsSatisfiedBy(user))

	spec = NewOrSpecification(eb, eb)
	s.True(spec.IsSatisfiedBy(user))
}

func TestOrSpecificationSuite(t *testing.T) {
	suite.Run(t, new(orSpecificationSuite))
}
