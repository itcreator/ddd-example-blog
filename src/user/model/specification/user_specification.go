package specification

import "user/model/entity"

type UserSpecification interface {
	IsSatisfiedBy(user entity.User) bool
}
