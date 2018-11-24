package specification

import "user/model/entity"

type andSpecification struct {
	one UserSpecification
	two UserSpecification
}

func (a *andSpecification) IsSatisfiedBy(user entity.User) bool {
	return a.one.IsSatisfiedBy(user) && a.two.IsSatisfiedBy(user)
}

func NewAndSpecification(one UserSpecification, two UserSpecification) UserSpecification {
	return &andSpecification{one, two}
}
