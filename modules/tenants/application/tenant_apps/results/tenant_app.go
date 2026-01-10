package results

import (
	"time"

	"nfxid/modules/tenants/domain/tenant_apps"

	"github.com/google/uuid"
)

type TenantAppRO struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	AppID     uuid.UUID
	Status    tenant_apps.TenantAppStatus
	CreatedAt time.Time
	CreatedBy *uuid.UUID
	UpdatedAt time.Time
	Settings  map[string]interface{}
}

// TenantAppMapper 将 Domain TenantApp 转换为 Application TenantAppRO
func TenantAppMapper(ta *tenant_apps.TenantApp) TenantAppRO {
	if ta == nil {
		return TenantAppRO{}
	}

	return TenantAppRO{
		ID:        ta.ID(),
		TenantID:  ta.TenantID(),
		AppID:     ta.AppID(),
		Status:    ta.Status(),
		CreatedAt: ta.CreatedAt(),
		CreatedBy: ta.CreatedBy(),
		UpdatedAt: ta.UpdatedAt(),
		Settings:  ta.Settings(),
	}
}
