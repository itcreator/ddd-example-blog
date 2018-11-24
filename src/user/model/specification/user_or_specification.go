package specification

import "user/model/entity"

type orSpecification struct {
	one UserSpecification
	two UserSpecification
}

func (a *orSpecification) IsSatisfiedBy(user entity.User) bool {
	return a.one.IsSatisfiedBy(user) || a.two.IsSatisfiedBy(user)
}

func NewOrSpecification(one UserSpecification, two UserSpecification) UserSpecification {
	return &orSpecification{one, two}
}
