package user_role

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 UserRole 的仓库结构体，包含增删改查四个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ur *UserRole) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*UserRole, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*UserRole, error)
	ByRoleID(ctx context.Context, roleID uuid.UUID) ([]*UserRole, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByUserAndRole(ctx context.Context, userID, roleID uuid.UUID) (bool, error)
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByUserAndRole(ctx context.Context, userID, roleID uuid.UUID) error
}
