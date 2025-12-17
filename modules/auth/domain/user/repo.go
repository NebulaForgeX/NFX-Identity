package user

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 User 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, u *User) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*User, error)
	ByUsername(ctx context.Context, username string) (*User, error)
	ByEmail(ctx context.Context, email string) (*User, error)
	ByPhone(ctx context.Context, phone string) (*User, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByUsername(ctx context.Context, username string) (bool, error)
	ByEmail(ctx context.Context, email string) (bool, error)
	ByPhone(ctx context.Context, phone string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, u *User) error
	Password(ctx context.Context, userID uuid.UUID, hashedPassword string) error
	Status(ctx context.Context, userID uuid.UUID, status string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
