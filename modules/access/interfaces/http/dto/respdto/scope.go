package respdto

import (
	"time"

	scopeAppResult "nfxid/modules/access/application/scopes/results"
)

type ScopeDTO struct {
	Scope       string     `json:"scope"`
	Description *string    `json:"description,omitempty"`
	IsSystem    bool       `json:"is_system"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// ScopeROToDTO converts application ScopeRO to response DTO
func ScopeROToDTO(v *scopeAppResult.ScopeRO) *ScopeDTO {
	if v == nil {
		return nil
	}

	return &ScopeDTO{
		Scope:       v.Scope,
		Description: v.Description,
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}
