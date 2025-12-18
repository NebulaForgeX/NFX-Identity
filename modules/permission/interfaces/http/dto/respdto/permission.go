package respdto

import (
	permissionViews "nfxid/modules/permission/application/permission/views"
	"time"

	"github.com/google/uuid"
)

type PermissionDTO struct {
	ID          uuid.UUID `json:"id"`
	Tag         string    `json:"tag"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Category    string    `json:"category,omitempty"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func PermissionViewToDTO(v *permissionViews.PermissionView) *PermissionDTO {
	if v == nil {
		return nil
	}

	return &PermissionDTO{
		ID:          v.ID,
		Tag:         v.Tag,
		Name:        v.Name,
		Description: v.Description,
		Category:    string(v.Category), // Convert enum to string for JSON
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
