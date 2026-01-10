package results

import (
	"time"

	"nfxid/modules/auth/domain/user_credentials"

	"github.com/google/uuid"
)

type UserCredentialRO struct {
	ID                 uuid.UUID
	UserID             uuid.UUID
	TenantID           uuid.UUID
	CredentialType     user_credentials.CredentialType
	PasswordHash       *string
	HashAlg            *string
	HashParams         map[string]interface{}
	PasswordUpdatedAt  *time.Time
	LastSuccessLoginAt *time.Time
	Status             user_credentials.CredentialStatus
	MustChangePassword bool
	Version            int
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}

// UserCredentialMapper 将 Domain UserCredential 转换为 Application UserCredentialRO
func UserCredentialMapper(uc *user_credentials.UserCredential) UserCredentialRO {
	if uc == nil {
		return UserCredentialRO{}
	}

	return UserCredentialRO{
		ID:                 uc.ID(),
		UserID:             uc.UserID(),
		TenantID:           uc.TenantID(),
		CredentialType:     uc.CredentialType(),
		PasswordHash:       uc.PasswordHash(),
		HashAlg:            uc.HashAlg(),
		HashParams:         uc.HashParams(),
		PasswordUpdatedAt:  uc.PasswordUpdatedAt(),
		LastSuccessLoginAt: uc.LastSuccessLoginAt(),
		Status:             uc.Status(),
		MustChangePassword: uc.MustChangePassword(),
		Version:            uc.Version(),
		CreatedAt:          uc.CreatedAt(),
		UpdatedAt:          uc.UpdatedAt(),
		DeletedAt:          uc.DeletedAt(),
	}
}
