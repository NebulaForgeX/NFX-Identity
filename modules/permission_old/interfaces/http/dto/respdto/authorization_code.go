package respdto

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
	"time"

	"github.com/google/uuid"
)

type AuthorizationCodeDTO struct {
	ID          uuid.UUID  `json:"id"`
	Code        string     `json:"code"`
	MaxUses     int        `json:"max_uses"`
	UsedCount   int        `json:"used_count"`
	CreatedBy   *uuid.UUID `json:"created_by,omitempty"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	IsActive    bool       `json:"is_active"`
	IsAvailable bool       `json:"is_available"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func AuthorizationCodeDomainToDTO(ac *authorizationCodeDomain.AuthorizationCode) *AuthorizationCodeDTO {
	if ac == nil {
		return nil
	}

	editable := ac.Editable()

	return &AuthorizationCodeDTO{
		ID:          ac.ID(),
		Code:        editable.Code,
		MaxUses:     editable.MaxUses,
		UsedCount:   editable.UsedCount,
		CreatedBy:   ac.CreatedBy(),
		ExpiresAt:   ac.ExpiresAt(),
		IsActive:    ac.IsActive(),
		IsAvailable: ac.IsAvailable(),
		CreatedAt:   ac.CreatedAt(),
		UpdatedAt:   ac.UpdatedAt(),
	}
}
