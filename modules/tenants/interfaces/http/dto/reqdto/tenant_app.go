package reqdto

import (
	tenantAppAppCommands "nfxid/modules/tenants/application/tenant_apps/commands"
	tenantAppDomain "nfxid/modules/tenants/domain/tenant_apps"

	"github.com/google/uuid"
)

type TenantAppCreateRequestDTO struct {
	TenantID  uuid.UUID              `json:"tenant_id" validate:"required,uuid"`
	AppID     uuid.UUID              `json:"app_id" validate:"required,uuid"`
	Status    string                 `json:"status,omitempty"`
	CreatedBy *uuid.UUID             `json:"created_by,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
}

type TenantAppUpdateStatusRequestDTO struct {
	ID     uuid.UUID `uri:"id" validate:"required,uuid"`
	Status string    `json:"status" validate:"required"`
}

type TenantAppUpdateSettingsRequestDTO struct {
	ID       uuid.UUID              `uri:"id" validate:"required,uuid"`
	Settings map[string]interface{} `json:"settings" validate:"required"`
}

func (r *TenantAppCreateRequestDTO) ToCreateCmd() tenantAppAppCommands.CreateTenantAppCmd {
	cmd := tenantAppAppCommands.CreateTenantAppCmd{
		TenantID:  r.TenantID,
		AppID:     r.AppID,
		CreatedBy: r.CreatedBy,
		Settings:  r.Settings,
	}

	// Parse status
	if r.Status != "" {
		cmd.Status = tenantAppDomain.TenantAppStatus(r.Status)
	} else {
		cmd.Status = tenantAppDomain.TenantAppStatusActive
	}

	return cmd
}

func (r *TenantAppUpdateStatusRequestDTO) ToUpdateStatusCmd() tenantAppAppCommands.UpdateTenantAppStatusCmd {
	return tenantAppAppCommands.UpdateTenantAppStatusCmd{
		TenantAppID: r.ID,
		Status:      tenantAppDomain.TenantAppStatus(r.Status),
	}
}

func (r *TenantAppUpdateSettingsRequestDTO) ToUpdateSettingsCmd() tenantAppAppCommands.UpdateTenantAppSettingsCmd {
	return tenantAppAppCommands.UpdateTenantAppSettingsCmd{
		TenantAppID: r.ID,
		Settings:    r.Settings,
	}
}
