package user_avatars

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 UserAvatar 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ua *UserAvatar) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByUserID(ctx context.Context, userID uuid.UUID) (*UserAvatar, error)
	ByImageID(ctx context.Context, imageID uuid.UUID) (*UserAvatar, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
	ByImageID(ctx context.Context, imageID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ua *UserAvatar) error
	ImageID(ctx context.Context, userID uuid.UUID, imageID uuid.UUID) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByUserID(ctx context.Context, userID uuid.UUID) error
}
