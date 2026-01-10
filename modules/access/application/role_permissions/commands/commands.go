package commands

import (
	"github.com/google/uuid"
)

// CreateRolePermissionCmd 创建角色权限命令
type CreateRolePermissionCmd struct {
	RoleID       uuid.UUID
	PermissionID uuid.UUID
	CreatedBy    *uuid.UUID
}

// DeleteRolePermissionCmd 删除角色权限命令
type DeleteRolePermissionCmd struct {
	RolePermissionID uuid.UUID
}

// DeleteRolePermissionByRoleAndPermissionCmd 根据角色和权限删除角色权限命令
type DeleteRolePermissionByRoleAndPermissionCmd struct {
	RoleID       uuid.UUID
	PermissionID uuid.UUID
}

// DeleteRolePermissionsByRoleCmd 根据角色删除所有角色权限命令
type DeleteRolePermissionsByRoleCmd struct {
	RoleID uuid.UUID
}
