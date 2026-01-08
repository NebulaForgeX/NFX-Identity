package client_credentials

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 ClientCredential 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, cc *ClientCredential) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ClientCredential, error)
	ByClientID(ctx context.Context, clientID string) (*ClientCredential, error)
	ByAppID(ctx context.Context, appID uuid.UUID) ([]*ClientCredential, error)
	ActiveByAppID(ctx context.Context, appID uuid.UUID) (*ClientCredential, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByClientID(ctx context.Context, clientID string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, cc *ClientCredential) error
	UpdateLastUsed(ctx context.Context, clientID string) error
	Revoke(ctx context.Context, clientID string, revokedBy uuid.UUID, reason string) error
	Rotate(ctx context.Context, clientID string, newSecretHash, newHashAlg string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByClientID(ctx context.Context, clientID string) error
}
