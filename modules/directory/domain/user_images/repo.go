package user_images

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 UserImage 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ui *UserImage) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*UserImage, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*UserImage, error)
	ByImageID(ctx context.Context, imageID uuid.UUID) ([]*UserImage, error)
	ByUserIDAndImageID(ctx context.Context, userID, imageID uuid.UUID) (*UserImage, error)
	CurrentByUserID(ctx context.Context, userID uuid.UUID) (*UserImage, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByUserIDAndImageID(ctx context.Context, userID, imageID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ui *UserImage) error
	DisplayOrder(ctx context.Context, id uuid.UUID, displayOrder int) error
	ImageID(ctx context.Context, id uuid.UUID, imageID uuid.UUID) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByUserID(ctx context.Context, userID uuid.UUID) error
}
