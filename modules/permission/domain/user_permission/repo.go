package user_permission

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 UserPermission 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, up *UserPermission) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*UserPermission, error)
	ByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) (*UserPermission, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) (bool, error)
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) error
	ByUserID(ctx context.Context, userID uuid.UUID) error
}
