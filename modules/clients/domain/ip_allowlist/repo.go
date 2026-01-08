package ip_allowlist

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 IPAllowlist 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ip *IPAllowlist) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*IPAllowlist, error)
	ByRuleID(ctx context.Context, ruleID string) (*IPAllowlist, error)
	ByAppID(ctx context.Context, appID uuid.UUID) ([]*IPAllowlist, error)
	ActiveByAppID(ctx context.Context, appID uuid.UUID) ([]*IPAllowlist, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByRuleID(ctx context.Context, ruleID string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ip *IPAllowlist) error
	Status(ctx context.Context, id uuid.UUID, status AllowlistStatus) error
	Revoke(ctx context.Context, ruleID string, revokedBy uuid.UUID, reason string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByRuleID(ctx context.Context, ruleID string) error
}
