package authorization_code

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 AuthorizationCode 的仓库结构体，包含增删改查等子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ac *AuthorizationCode) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*AuthorizationCode, error)
	ByCode(ctx context.Context, code string) (*AuthorizationCode, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByCode(ctx context.Context, code string) (bool, error)
	ByCodeAndIncrement(ctx context.Context, code string) (*AuthorizationCode, error) // 检查有效则 used_count +1
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ac *AuthorizationCode) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
