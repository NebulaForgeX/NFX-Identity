package respdto

import (
	"time"

	permissionAppResult "nfxid/modules/access/application/permissions/results"

	"github.com/google/uuid"
)

type PermissionDTO struct {
	ID          uuid.UUID  `json:"id"`
	Key         string     `json:"key"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	IsSystem    bool       `json:"is_system"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// PermissionROToDTO converts application PermissionRO to response DTO
func PermissionROToDTO(v *permissionAppResult.PermissionRO) *PermissionDTO {
	if v == nil {
		return nil
	}

	return &PermissionDTO{
		ID:          v.ID,
		Key:         v.Key,
		Name:        v.Name,
		Description: v.Description,
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}

// PermissionListROToDTO converts list of PermissionRO to DTOs
func PermissionListROToDTO(results []permissionAppResult.PermissionRO) []PermissionDTO {
	dtos := make([]PermissionDTO, len(results))
	for i, v := range results {
		if dto := PermissionROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
