package respdto

import (
	"time"

	rolePermissionAppResult "nfxid/modules/access/application/role_permissions/results"

	"github.com/google/uuid"
)

type RolePermissionDTO struct {
	ID           uuid.UUID  `json:"id"`
	RoleID       uuid.UUID  `json:"role_id"`
	PermissionID uuid.UUID  `json:"permission_id"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    *uuid.UUID `json:"created_by,omitempty"`
}

// RolePermissionROToDTO converts application RolePermissionRO to response DTO
func RolePermissionROToDTO(v *rolePermissionAppResult.RolePermissionRO) *RolePermissionDTO {
	if v == nil {
		return nil
	}

	return &RolePermissionDTO{
		ID:           v.ID,
		RoleID:       v.RoleID,
		PermissionID: v.PermissionID,
		CreatedAt:    v.CreatedAt,
		CreatedBy:    v.CreatedBy,
	}
}
