package badge

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 Badge 的仓库结构体，包含增删改查四个子接口
type Repo struct {
	Create Create
	Get    Get
	Update Update
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, b *Badge) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Badge, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, b *Badge) error
}
