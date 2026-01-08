package role_permissions

import "github.com/google/uuid"

func (rp *RolePermission) Validate() error {
	if rp.RoleID() == uuid.Nil {
		return ErrRoleIDRequired
	}
	if rp.PermissionID() == uuid.Nil {
		return ErrPermissionIDRequired
	}
	return nil
}
