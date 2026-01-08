package api_keys

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 APIKey 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ak *APIKey) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*APIKey, error)
	ByKeyID(ctx context.Context, keyID string) (*APIKey, error)
	ByAppID(ctx context.Context, appID uuid.UUID) ([]*APIKey, error)
	ActiveByAppID(ctx context.Context, appID uuid.UUID) ([]*APIKey, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByKeyID(ctx context.Context, keyID string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ak *APIKey) error
	UpdateLastUsed(ctx context.Context, keyID string) error
	Revoke(ctx context.Context, keyID string, revokedBy uuid.UUID, reason string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByKeyID(ctx context.Context, keyID string) error
}
