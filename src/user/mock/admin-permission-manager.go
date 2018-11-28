package mock

import (
	"github.com/google/uuid"
	"user/model/entity"
)

//adminPermissionManager implements permission.Checker and permission.Granter interfaces
type adminPermissionManager struct {
	storage map[uuid.UUID]entity.User
}

func NewAdminPermissionManager() *adminPermissionManager {
	return &adminPermissionManager{make(map[uuid.UUID]entity.User)}
}

func (m *adminPermissionManager) GrantPermission(user entity.User) error {
	m.storage[user.GetUUID()] = user

	return nil
}

func (m *adminPermissionManager) RevokePermission(user entity.User) error {
	if nil != m.storage[user.GetUUID()] {
		delete(m.storage, user.GetUUID())
	}

	return nil
}

func (m *adminPermissionManager) CheckPermission(user entity.User) (bool, error) {
	return nil != m.storage[user.GetUUID()], nil
}
