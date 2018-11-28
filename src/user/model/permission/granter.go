package permission

import "user/model/entity"

type Granter interface {
	GrantPermission(user entity.User) error
	RevokePermission(user entity.User) error
}
