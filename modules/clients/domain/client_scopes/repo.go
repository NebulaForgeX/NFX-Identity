package client_scopes

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 ClientScope 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, cs *ClientScope) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ClientScope, error)
	ByAppID(ctx context.Context, appID uuid.UUID) ([]*ClientScope, error)
	ByAppIDAndScope(ctx context.Context, appID uuid.UUID, scope string) (*ClientScope, error)
	ActiveByAppID(ctx context.Context, appID uuid.UUID) ([]*ClientScope, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByAppIDAndScope(ctx context.Context, appID uuid.UUID, scope string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, cs *ClientScope) error
	Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID, reason string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByAppIDAndScope(ctx context.Context, appID uuid.UUID, scope string) error
}
