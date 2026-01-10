package respdto

import (
	"time"

	roleAppResult "nfxid/modules/access/application/roles/results"

	"github.com/google/uuid"
)

type RoleDTO struct {
	ID          uuid.UUID  `json:"id"`
	Key         string     `json:"key"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	ScopeType   string     `json:"scope_type"`
	IsSystem    bool       `json:"is_system"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// RoleROToDTO converts application RoleRO to response DTO
func RoleROToDTO(v *roleAppResult.RoleRO) *RoleDTO {
	if v == nil {
		return nil
	}

	return &RoleDTO{
		ID:          v.ID,
		Key:         v.Key,
		Name:        v.Name,
		Description: v.Description,
		ScopeType:   string(v.ScopeType),
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}

// RoleListROToDTO converts list of RoleRO to DTOs
func RoleListROToDTO(results []roleAppResult.RoleRO) []RoleDTO {
	dtos := make([]RoleDTO, len(results))
	for i, v := range results {
		if dto := RoleROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
