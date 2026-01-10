package commands

import (
	"github.com/google/uuid"
)

// CreatePermissionCmd 创建权限命令
type CreatePermissionCmd struct {
	Key         string
	Name        string
	Description *string
	IsSystem    bool
	CreatedBy   *uuid.UUID
}

// UpdatePermissionCmd 更新权限命令
type UpdatePermissionCmd struct {
	PermissionID uuid.UUID
	Name         string
	Description  *string
}

// DeletePermissionCmd 删除权限命令
type DeletePermissionCmd struct {
	PermissionID uuid.UUID
}
