package groups

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 Group 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, g *Group) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Group, error)
	ByGroupID(ctx context.Context, groupID string) (*Group, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*Group, error)
	ByParentGroupID(ctx context.Context, parentGroupID uuid.UUID) ([]*Group, error)
	ByType(ctx context.Context, groupType GroupType) ([]*Group, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByGroupID(ctx context.Context, groupID string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, g *Group) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
