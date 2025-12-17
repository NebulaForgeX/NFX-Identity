package profile

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 Profile 的仓库结构体，包含增删改查四个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, p *Profile) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Profile, error)
	ByUserID(ctx context.Context, userID uuid.UUID) (*Profile, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, p *Profile) error
	AndInsert(ctx context.Context, p *Profile) error // Upsert: Update if exists, Insert if not
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
