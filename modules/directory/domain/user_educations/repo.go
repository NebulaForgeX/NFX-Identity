package user_educations

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 UserEducation 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ue *UserEducation) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*UserEducation, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*UserEducation, error)
	CurrentByUserID(ctx context.Context, userID uuid.UUID) ([]*UserEducation, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ue *UserEducation) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
