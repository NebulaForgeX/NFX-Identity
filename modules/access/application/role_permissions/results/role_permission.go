package results

import (
	"time"

	"nfxid/modules/access/domain/role_permissions"

	"github.com/google/uuid"
)

type RolePermissionRO struct {
	ID           uuid.UUID
	RoleID       uuid.UUID
	PermissionID uuid.UUID
	CreatedAt    time.Time
	CreatedBy    *uuid.UUID
}

// RolePermissionMapper 将 Domain RolePermission 转换为 Application RolePermissionRO
func RolePermissionMapper(rp *role_permissions.RolePermission) RolePermissionRO {
	if rp == nil {
		return RolePermissionRO{}
	}

	return RolePermissionRO{
		ID:           rp.ID(),
		RoleID:       rp.RoleID(),
		PermissionID: rp.PermissionID(),
		CreatedAt:    rp.CreatedAt(),
		CreatedBy:    rp.CreatedBy(),
	}
}
