package tenant_roles

import (
	"context"

	"github.com/google/uuid"
)

// Repo 租户角色仓库接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 创建
type Create interface {
	New(ctx context.Context, r *TenantRole) error
}

// Get 查询
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*TenantRole, error)
	ByTenantIDAndRoleKey(ctx context.Context, tenantID uuid.UUID, roleKey string) (*TenantRole, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*TenantRole, error)
}

// Check 存在性检查
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByTenantIDAndRoleKey(ctx context.Context, tenantID uuid.UUID, roleKey string) (bool, error)
}

// Update 更新
type Update interface {
	Generic(ctx context.Context, r *TenantRole) error
}

// Delete 删除
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
