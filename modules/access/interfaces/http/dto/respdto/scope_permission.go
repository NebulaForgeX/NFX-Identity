package respdto

import (
	"time"

	scopePermissionAppResult "nfxid/modules/access/application/scope_permissions/results"

	"github.com/google/uuid"
)

type ScopePermissionDTO struct {
	ID           uuid.UUID `json:"id"`
	Scope        string    `json:"scope"`
	PermissionID uuid.UUID `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
}

// ScopePermissionROToDTO converts application ScopePermissionRO to response DTO
func ScopePermissionROToDTO(v *scopePermissionAppResult.ScopePermissionRO) *ScopePermissionDTO {
	if v == nil {
		return nil
	}

	return &ScopePermissionDTO{
		ID:           v.ID,
		Scope:        v.Scope,
		PermissionID: v.PermissionID,
		CreatedAt:    v.CreatedAt,
	}
}
