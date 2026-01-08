package account_lockouts

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 AccountLockout 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, al *AccountLockout) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByUserID(ctx context.Context, userID uuid.UUID) (*AccountLockout, error)
	ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) (*AccountLockout, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*AccountLockout, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
	ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) (bool, error)
	IsLocked(ctx context.Context, userID, tenantID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, al *AccountLockout) error
	Unlock(ctx context.Context, userID, tenantID uuid.UUID, unlockedBy string, unlockActorID *uuid.UUID) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByUserID(ctx context.Context, userID uuid.UUID) error
	ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) error
}
