package commands

import (
	"nfxid/modules/tenants/domain/tenants"

	"github.com/google/uuid"
)

// CreateTenantCmd 创建租户命令
type CreateTenantCmd struct {
	TenantID      string
	Name          string
	DisplayName   *string
	Status        tenants.TenantStatus
	PrimaryDomain *string
	Metadata      map[string]interface{}
}

// UpdateTenantCmd 更新租户命令
type UpdateTenantCmd struct {
	TenantID      uuid.UUID
	Name          string
	DisplayName   *string
	PrimaryDomain *string
	Metadata      map[string]interface{}
}

// UpdateTenantStatusCmd 更新租户状态命令
type UpdateTenantStatusCmd struct {
	TenantID uuid.UUID
	Status   tenants.TenantStatus
}

// DeleteTenantCmd 删除租户命令
type DeleteTenantCmd struct {
	TenantID uuid.UUID
}
