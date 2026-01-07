package reqdto

import (
	authorizationCodeCommands "nfxid/modules/permission/application/authorization_code/commands"
	"time"

	"github.com/google/uuid"
)

type AuthorizationCodeCreateRequestDTO struct {
	Code      string     `json:"code" validate:"required,min=1,max=255"`
	MaxUses   int        `json:"max_uses" validate:"required,min=1"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	IsActive  bool       `json:"is_active,omitempty"`
}

func (dto *AuthorizationCodeCreateRequestDTO) ToCreateCmd() authorizationCodeCommands.CreateAuthorizationCodeCmd {
	return authorizationCodeCommands.CreateAuthorizationCodeCmd{
		Code:      dto.Code,
		MaxUses:   dto.MaxUses,
		CreatedBy: dto.CreatedBy,
		ExpiresAt: dto.ExpiresAt,
		IsActive:  dto.IsActive,
	}
}

type AuthorizationCodeUseRequestDTO struct {
	Code string `json:"code" validate:"required"`
}

func (dto *AuthorizationCodeUseRequestDTO) ToUseCmd() authorizationCodeCommands.UseAuthorizationCodeCmd {
	return authorizationCodeCommands.UseAuthorizationCodeCmd{
		Code: dto.Code,
	}
}

type AuthorizationCodeByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type AuthorizationCodeByCodeRequestDTO struct {
	Code string `params:"code" validate:"required"`
}
