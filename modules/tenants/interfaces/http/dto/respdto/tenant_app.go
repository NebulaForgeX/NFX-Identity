package respdto

import (
	"time"

	tenantAppAppResult "nfxid/modules/tenants/application/tenant_apps/results"

	"github.com/google/uuid"
)

type TenantAppDTO struct {
	ID        uuid.UUID              `json:"id"`
	TenantID  uuid.UUID              `json:"tenant_id"`
	AppID     uuid.UUID              `json:"app_id"`
	Status    string                 `json:"status"`
	CreatedBy *uuid.UUID             `json:"created_by,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

// TenantAppROToDTO converts application TenantAppRO to response DTO
func TenantAppROToDTO(v *tenantAppAppResult.TenantAppRO) *TenantAppDTO {
	if v == nil {
		return nil
	}

	return &TenantAppDTO{
		ID:        v.ID,
		TenantID:  v.TenantID,
		AppID:     v.AppID,
		Status:    string(v.Status),
		CreatedBy: v.CreatedBy,
		Settings:  v.Settings,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}

// TenantAppListROToDTO converts list of TenantAppRO to DTOs
func TenantAppListROToDTO(results []tenantAppAppResult.TenantAppRO) []TenantAppDTO {
	dtos := make([]TenantAppDTO, len(results))
	for i, v := range results {
		if dto := TenantAppROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
