package rate_limits

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 RateLimit 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, rl *RateLimit) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*RateLimit, error)
	ByAppID(ctx context.Context, appID uuid.UUID) ([]*RateLimit, error)
	ByAppIDAndLimitType(ctx context.Context, appID uuid.UUID, limitType RateLimitType) (*RateLimit, error)
	ActiveByAppID(ctx context.Context, appID uuid.UUID) ([]*RateLimit, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByAppIDAndLimitType(ctx context.Context, appID uuid.UUID, limitType RateLimitType) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, rl *RateLimit) error
	UpdateLimit(ctx context.Context, id uuid.UUID, limitValue, windowSeconds int, description *string) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByAppIDAndLimitType(ctx context.Context, appID uuid.UUID, limitType RateLimitType) error
}
