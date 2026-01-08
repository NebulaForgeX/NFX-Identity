package user_badges

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 UserBadge 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ub *UserBadge) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*UserBadge, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*UserBadge, error)
	ByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]*UserBadge, error)
	ByUserIDAndBadgeID(ctx context.Context, userID, badgeID uuid.UUID) (*UserBadge, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByUserIDAndBadgeID(ctx context.Context, userID, badgeID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ub *UserBadge) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByUserIDAndBadgeID(ctx context.Context, userID, badgeID uuid.UUID) error
}
