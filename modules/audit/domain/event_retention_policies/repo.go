package event_retention_policies

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 EventRetentionPolicy 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, erp *EventRetentionPolicy) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*EventRetentionPolicy, error)
	ByPolicyName(ctx context.Context, policyName string) (*EventRetentionPolicy, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*EventRetentionPolicy, error)
	ByStatus(ctx context.Context, status string) ([]*EventRetentionPolicy, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByPolicyName(ctx context.Context, policyName string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, erp *EventRetentionPolicy) error
	Status(ctx context.Context, id uuid.UUID, status string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
