package image

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 Image 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, img *Image) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Image, error)
	ByFilename(ctx context.Context, filename string) (*Image, error)
	ByStoragePath(ctx context.Context, storagePath string) (*Image, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, img *Image) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByStoragePath(ctx context.Context, storagePath string) error
}
