package reqdto

import (
	mfaFactorAppCommands "nfxid/modules/auth/application/mfa_factors/commands"
	mfaFactorDomain "nfxid/modules/auth/domain/mfa_factors"

	"github.com/google/uuid"
)

type MFAFactorCreateRequestDTO struct {
	FactorID         string  `json:"factor_id" validate:"required"`
	TenantID         uuid.UUID `json:"tenant_id" validate:"required"`
	UserID           uuid.UUID `json:"user_id" validate:"required"`
	Type             string  `json:"type" validate:"required"`
	SecretEncrypted  *string `json:"secret_encrypted,omitempty"`
	Phone            *string `json:"phone,omitempty"`
	Email            *string `json:"email,omitempty"`
	Name             *string `json:"name,omitempty"`
	Enabled          bool    `json:"enabled"`
	RecoveryCodesHash *string `json:"recovery_codes_hash,omitempty"`
}

type MFAFactorUpdateRequestDTO struct {
	Name    *string `json:"name,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

type MFAFactorByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

func (r *MFAFactorCreateRequestDTO) ToCreateCmd() mfaFactorAppCommands.CreateMFAFactorCmd {
	return mfaFactorAppCommands.CreateMFAFactorCmd{
		FactorID:         r.FactorID,
		TenantID:         r.TenantID,
		UserID:           r.UserID,
		Type:             mfaFactorDomain.MFAType(r.Type),
		SecretEncrypted:  r.SecretEncrypted,
		Phone:            r.Phone,
		Email:            r.Email,
		Name:             r.Name,
		Enabled:          r.Enabled,
		RecoveryCodesHash: r.RecoveryCodesHash,
	}
}

func (r *MFAFactorUpdateRequestDTO) ToUpdateCmd(factorID string) mfaFactorAppCommands.UpdateMFAFactorCmd {
	cmd := mfaFactorAppCommands.UpdateMFAFactorCmd{
		FactorID: factorID,
	}
	if r.Name != nil {
		cmd.Name = r.Name
	}
	if r.Enabled != nil {
		cmd.Enabled = *r.Enabled
	}
	return cmd
}
