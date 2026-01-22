package respdto

import (
	"time"

	userCredentialAppResult "nfxid/modules/auth/application/user_credentials/results"

	"github.com/google/uuid"
)

type UserCredentialDTO struct {
	ID                 uuid.UUID              `json:"id"`
	UserID             uuid.UUID              `json:"user_id"`
	CredentialType     string                 `json:"credential_type"`
	PasswordHash       *string                `json:"password_hash,omitempty"`
	HashAlg            *string                `json:"hash_alg,omitempty"`
	HashParams         map[string]interface{} `json:"hash_params,omitempty"`
	PasswordUpdatedAt  *time.Time             `json:"password_updated_at,omitempty"`
	LastSuccessLoginAt *time.Time             `json:"last_success_login_at,omitempty"`
	Status             string                 `json:"status"`
	MustChangePassword bool                   `json:"must_change_password"`
	Version            int                    `json:"version"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
	DeletedAt          *time.Time             `json:"deleted_at,omitempty"`
}

// UserCredentialROToDTO converts application UserCredentialRO to response DTO
func UserCredentialROToDTO(v *userCredentialAppResult.UserCredentialRO) *UserCredentialDTO {
	if v == nil {
		return nil
	}

	return &UserCredentialDTO{
		ID:                 v.ID,
		UserID:             v.UserID,
		CredentialType:     string(v.CredentialType),
		PasswordHash:       v.PasswordHash,
		HashAlg:            v.HashAlg,
		HashParams:         v.HashParams,
		PasswordUpdatedAt:  v.PasswordUpdatedAt,
		LastSuccessLoginAt: v.LastSuccessLoginAt,
		Status:             string(v.Status),
		MustChangePassword: v.MustChangePassword,
		Version:            v.Version,
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
		DeletedAt:          v.DeletedAt,
	}
}
