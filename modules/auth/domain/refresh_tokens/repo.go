package refresh_tokens

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 RefreshToken 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, rt *RefreshToken) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*RefreshToken, error)
	ByTokenID(ctx context.Context, tokenID string) (*RefreshToken, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*RefreshToken, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByTokenID(ctx context.Context, tokenID string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, rt *RefreshToken) error
	Revoke(ctx context.Context, tokenID string, reason RevokeReason) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByTokenID(ctx context.Context, tokenID string) error
}
