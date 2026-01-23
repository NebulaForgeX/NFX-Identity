package respdto

import (
	"time"

	mfaFactorAppResult "nfxid/modules/auth/application/mfa_factors/results"

	"github.com/google/uuid"
)

type MFAFactorDTO struct {
	ID                uuid.UUID  `json:"id"`
	FactorID          string     `json:"factor_id"`
	TenantID          uuid.UUID  `json:"tenant_id"`
	UserID            uuid.UUID  `json:"user_id"`
	Type              string     `json:"type"`
	SecretEncrypted   *string    `json:"secret_encrypted,omitempty"`
	Phone             *string    `json:"phone,omitempty"`
	Email             *string    `json:"email,omitempty"`
	Name              *string    `json:"name,omitempty"`
	Enabled           bool       `json:"enabled"`
	CreatedAt         time.Time  `json:"created_at"`
	LastUsedAt        *time.Time `json:"last_used_at,omitempty"`
	RecoveryCodesHash *string    `json:"recovery_codes_hash,omitempty"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
}

// MFAFactorROToDTO converts application MFAFactorRO to response DTO
func MFAFactorROToDTO(v *mfaFactorAppResult.MFAFactorRO) *MFAFactorDTO {
	if v == nil {
		return nil
	}

	return &MFAFactorDTO{
		ID:                v.ID,
		FactorID:          v.FactorID,
		TenantID:          v.TenantID,
		UserID:            v.UserID,
		Type:              string(v.Type),
		SecretEncrypted:   v.SecretEncrypted,
		Phone:             v.Phone,
		Email:             v.Email,
		Name:              v.Name,
		Enabled:           v.Enabled,
		CreatedAt:         v.CreatedAt,
		LastUsedAt:        v.LastUsedAt,
		RecoveryCodesHash: v.RecoveryCodesHash,
		UpdatedAt:         v.UpdatedAt,
		DeletedAt:         v.DeletedAt,
	}
}
