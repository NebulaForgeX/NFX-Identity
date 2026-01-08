package image_tags

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 ImageTag 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, it *ImageTag) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ImageTag, error)
	ByImageID(ctx context.Context, imageID uuid.UUID) ([]*ImageTag, error)
	ByTag(ctx context.Context, tag string) ([]*ImageTag, error)
	ByImageIDAndTag(ctx context.Context, imageID uuid.UUID, tag string) (*ImageTag, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByImageIDAndTag(ctx context.Context, imageID uuid.UUID, tag string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, it *ImageTag) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByImageIDAndTag(ctx context.Context, imageID uuid.UUID, tag string) error
}
