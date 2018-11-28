package permission

import "user/model/entity"

type Checker interface {
	CheckPermission(user entity.User) (bool, error)
}
