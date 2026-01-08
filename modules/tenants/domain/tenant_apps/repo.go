package tenant_apps

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 TenantApp 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, ta *TenantApp) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*TenantApp, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*TenantApp, error)
	ByAppID(ctx context.Context, appID uuid.UUID) ([]*TenantApp, error)
	ByTenantIDAndAppID(ctx context.Context, tenantID, appID uuid.UUID) (*TenantApp, error)
	ByStatus(ctx context.Context, status TenantAppStatus) ([]*TenantApp, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByTenantIDAndAppID(ctx context.Context, tenantID, appID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, ta *TenantApp) error
	Status(ctx context.Context, id uuid.UUID, status TenantAppStatus) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByTenantIDAndAppID(ctx context.Context, tenantID, appID uuid.UUID) error
}
