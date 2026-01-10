package respdto

import (
	"time"

	tenantAppResult "nfxid/modules/tenants/application/tenants/results"

	"github.com/google/uuid"
)

type TenantDTO struct {
	ID            uuid.UUID              `json:"id"`
	TenantID      string                 `json:"tenant_id"`
	Name          string                 `json:"name"`
	DisplayName   *string                `json:"display_name,omitempty"`
	Status        string                 `json:"status"`
	PrimaryDomain *string                `json:"primary_domain,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	DeletedAt     *time.Time             `json:"deleted_at,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// TenantROToDTO converts application TenantRO to response DTO
func TenantROToDTO(v *tenantAppResult.TenantRO) *TenantDTO {
	if v == nil {
		return nil
	}

	return &TenantDTO{
		ID:            v.ID,
		TenantID:      v.TenantID,
		Name:          v.Name,
		DisplayName:   v.DisplayName,
		Status:        string(v.Status),
		PrimaryDomain: v.PrimaryDomain,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		DeletedAt:     v.DeletedAt,
		Metadata:      v.Metadata,
	}
}

// TenantListROToDTO converts list of TenantRO to DTOs
func TenantListROToDTO(results []tenantAppResult.TenantRO) []TenantDTO {
	dtos := make([]TenantDTO, len(results))
	for i, v := range results {
		if dto := TenantROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
