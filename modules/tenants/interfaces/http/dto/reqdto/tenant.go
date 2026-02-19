package reqdto

import (
	tenantAppCommands "nfxid/modules/tenants/application/tenants/commands"
	tenantDomain "nfxid/modules/tenants/domain/tenants"

	"github.com/google/uuid"
)

type TenantCreateRequestDTO struct {
	TenantID      string                 `json:"tenant_id" validate:"required"`
	Name          string                 `json:"name" validate:"required"`
	DisplayName   *string                `json:"display_name,omitempty"`
	Status        string                 `json:"status,omitempty"`
	PrimaryDomain *string                `json:"primary_domain,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

type TenantUpdateRequestDTO struct {
	ID            uuid.UUID              `uri:"id" validate:"required,uuid"`
	Name          string                 `json:"name" validate:"required"`
	DisplayName   *string                `json:"display_name,omitempty"`
	PrimaryDomain *string                `json:"primary_domain,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

type TenantUpdateStatusRequestDTO struct {
	ID     uuid.UUID `uri:"id" validate:"required,uuid"`
	Status string    `json:"status" validate:"required"`
}

type TenantByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type TenantByTenantIDRequestDTO struct {
	TenantID string `uri:"tenant_id" validate:"required"`
}

func (r *TenantCreateRequestDTO) ToCreateCmd() tenantAppCommands.CreateTenantCmd {
	cmd := tenantAppCommands.CreateTenantCmd{
		TenantID:      r.TenantID,
		Name:          r.Name,
		DisplayName:   r.DisplayName,
		PrimaryDomain: r.PrimaryDomain,
		Metadata:      r.Metadata,
	}

	// Parse status
	if r.Status != "" {
		cmd.Status = tenantDomain.TenantStatus(r.Status)
	} else {
		cmd.Status = tenantDomain.TenantStatusActive
	}

	return cmd
}

func (r *TenantUpdateRequestDTO) ToUpdateCmd() tenantAppCommands.UpdateTenantCmd {
	return tenantAppCommands.UpdateTenantCmd{
		TenantID:      r.ID,
		Name:          r.Name,
		DisplayName:   r.DisplayName,
		PrimaryDomain: r.PrimaryDomain,
		Metadata:      r.Metadata,
	}
}

func (r *TenantUpdateStatusRequestDTO) ToUpdateStatusCmd() tenantAppCommands.UpdateTenantStatusCmd {
	return tenantAppCommands.UpdateTenantStatusCmd{
		TenantID: r.ID,
		Status:   tenantDomain.TenantStatus(r.Status),
	}
}
