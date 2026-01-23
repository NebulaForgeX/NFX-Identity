package login_attempts

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 LoginAttempt 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, la *LoginAttempt) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*LoginAttempt, error)
	ByIdentifier(ctx context.Context, identifier string) ([]*LoginAttempt, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*LoginAttempt, error)
	ByIP(ctx context.Context, ip string) ([]*LoginAttempt, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, la *LoginAttempt) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
