package commands

import (
	"nfxid/modules/tenants/domain/tenant_apps"

	"github.com/google/uuid"
)

// CreateTenantAppCmd 创建租户应用命令
type CreateTenantAppCmd struct {
	TenantID  uuid.UUID
	AppID     uuid.UUID
	Status    tenant_apps.TenantAppStatus
	CreatedBy *uuid.UUID
	Settings  map[string]interface{}
}

// UpdateTenantAppStatusCmd 更新租户应用状态命令
type UpdateTenantAppStatusCmd struct {
	TenantAppID uuid.UUID
	Status      tenant_apps.TenantAppStatus
}

// UpdateTenantAppSettingsCmd 更新租户应用设置命令
type UpdateTenantAppSettingsCmd struct {
	TenantAppID uuid.UUID
	Settings    map[string]interface{}
}

// DeleteTenantAppCmd 删除租户应用命令
type DeleteTenantAppCmd struct {
	TenantAppID uuid.UUID
}
