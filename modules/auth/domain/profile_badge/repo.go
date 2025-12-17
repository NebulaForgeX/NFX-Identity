package profile_badge

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 ProfileBadge 的仓库结构体，包含增删改查四个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, pb *ProfileBadge) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ProfileBadge, error)
	ByProfileID(ctx context.Context, profileID uuid.UUID) ([]*ProfileBadge, error)
	ByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]*ProfileBadge, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByProfileAndBadge(ctx context.Context, profileID, badgeID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, pb *ProfileBadge) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByProfileAndBadge(ctx context.Context, profileID, badgeID uuid.UUID) error
}
