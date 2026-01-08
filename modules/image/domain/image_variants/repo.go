package image_variants

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 ImageVariant 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, iv *ImageVariant) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ImageVariant, error)
	ByImageID(ctx context.Context, imageID uuid.UUID) ([]*ImageVariant, error)
	ByImageIDAndVariantKey(ctx context.Context, imageID uuid.UUID, variantKey string) (*ImageVariant, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByImageIDAndVariantKey(ctx context.Context, imageID uuid.UUID, variantKey string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, iv *ImageVariant) error
	UpdateURL(ctx context.Context, id uuid.UUID, url string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByImageIDAndVariantKey(ctx context.Context, imageID uuid.UUID, variantKey string) error
}
