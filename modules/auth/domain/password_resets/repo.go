package password_resets

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 PasswordReset 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, pr *PasswordReset) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*PasswordReset, error)
	ByResetID(ctx context.Context, resetID string) (*PasswordReset, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*PasswordReset, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByResetID(ctx context.Context, resetID string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, pr *PasswordReset) error
	MarkAsUsed(ctx context.Context, resetID string) error
	IncrementAttemptCount(ctx context.Context, resetID string) error
	UpdateStatus(ctx context.Context, resetID string, status ResetStatus) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByResetID(ctx context.Context, resetID string) error
}
