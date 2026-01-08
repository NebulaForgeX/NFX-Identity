package scope_permissions

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 ScopePermission 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, sp *ScopePermission) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ScopePermission, error)
	ByScopeAndPermissionID(ctx context.Context, scope string, permissionID uuid.UUID) (*ScopePermission, error)
	ByScope(ctx context.Context, scope string) ([]*ScopePermission, error)
	ByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]*ScopePermission, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByScopeAndPermissionID(ctx context.Context, scope string, permissionID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, sp *ScopePermission) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByScopeAndPermissionID(ctx context.Context, scope string, permissionID uuid.UUID) error
	ByScope(ctx context.Context, scope string) error
}
