package user_emails

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 UserEmail 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ue *UserEmail) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*UserEmail, error)
	ByEmail(ctx context.Context, email string) (*UserEmail, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*UserEmail, error)
	PrimaryByUserID(ctx context.Context, userID uuid.UUID) (*UserEmail, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByEmail(ctx context.Context, email string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ue *UserEmail) error
	SetPrimary(ctx context.Context, id uuid.UUID) error
	Verify(ctx context.Context, id uuid.UUID) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByEmail(ctx context.Context, email string) error
}
