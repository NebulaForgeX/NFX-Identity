package respdto

import (
	"time"

	appAppResult "nfxid/modules/clients/application/apps/results"

	"github.com/google/uuid"
)

type AppDTO struct {
	ID          uuid.UUID              `json:"id"`
	AppID       string                 `json:"app_id"`
	TenantID    uuid.UUID              `json:"tenant_id"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	Environment string                 `json:"environment"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	CreatedBy   *uuid.UUID             `json:"created_by,omitempty"`
	UpdatedBy   *uuid.UUID             `json:"updated_by,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	DeletedAt   *time.Time             `json:"deleted_at,omitempty"`
}

// AppROToDTO converts application AppRO to response DTO
func AppROToDTO(v *appAppResult.AppRO) *AppDTO {
	if v == nil {
		return nil
	}

	return &AppDTO{
		ID:          v.ID,
		AppID:       v.AppID,
		TenantID:    v.TenantID,
		Name:        v.Name,
		Description: v.Description,
		Type:        string(v.Type),
		Status:      string(v.Status),
		Environment: string(v.Environment),
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		CreatedBy:   v.CreatedBy,
		UpdatedBy:   v.UpdatedBy,
		Metadata:    v.Metadata,
		DeletedAt:   v.DeletedAt,
	}
}
