package commands

import (
	"nfxid/modules/access/domain/roles"

	"github.com/google/uuid"
)

// CreateRoleCmd 创建角色命令
type CreateRoleCmd struct {
	Key         string
	Name        string
	Description *string
	ScopeType   roles.ScopeType
	IsSystem    bool
	CreatedBy   *uuid.UUID
}

// UpdateRoleCmd 更新角色命令
type UpdateRoleCmd struct {
	RoleID      uuid.UUID
	Name        string
	Description *string
	ScopeType   roles.ScopeType
}

// DeleteRoleCmd 删除角色命令
type DeleteRoleCmd struct {
	RoleID uuid.UUID
}
