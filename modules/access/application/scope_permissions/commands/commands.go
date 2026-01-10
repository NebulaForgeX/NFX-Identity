package commands

import (
	"github.com/google/uuid"
)

// CreateScopePermissionCmd 创建作用域权限命令
type CreateScopePermissionCmd struct {
	Scope        string
	PermissionID uuid.UUID
}

// DeleteScopePermissionCmd 删除作用域权限命令
type DeleteScopePermissionCmd struct {
	ScopePermissionID uuid.UUID
}

// DeleteScopePermissionByScopeAndPermissionCmd 根据作用域和权限删除作用域权限命令
type DeleteScopePermissionByScopeAndPermissionCmd struct {
	Scope        string
	PermissionID uuid.UUID
}

// DeleteScopePermissionsByScopeCmd 根据作用域删除所有作用域权限命令
type DeleteScopePermissionsByScopeCmd struct {
	Scope string
}
