package reqdto

import (
	userCredentialAppCommands "nfxid/modules/auth/application/user_credentials/commands"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"

	"github.com/google/uuid"
)

type UserCredentialCreateRequestDTO struct {
	UserID            uuid.UUID                `json:"user_id" validate:"required"`
	TenantID          uuid.UUID                `json:"tenant_id" validate:"required"`
	CredentialType    string                    `json:"credential_type" validate:"required"`
	PasswordHash      *string                   `json:"password_hash,omitempty"`
	HashAlg           *string                   `json:"hash_alg,omitempty"`
	HashParams        map[string]interface{}     `json:"hash_params,omitempty"`
	Status            string                    `json:"status,omitempty"`
	MustChangePassword bool                     `json:"must_change_password,omitempty"`
}

type UserCredentialUpdateRequestDTO struct {
	ID                uuid.UUID `params:"id" validate:"required,uuid"`
	Status            string    `json:"status,omitempty"`
	MustChangePassword bool     `json:"must_change_password,omitempty"`
}

type UserCredentialByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

func (r *UserCredentialCreateRequestDTO) ToCreateCmd() userCredentialAppCommands.CreateUserCredentialCmd {
	cmd := userCredentialAppCommands.CreateUserCredentialCmd{
		UserID:            r.UserID,
		TenantID:          r.TenantID,
		PasswordHash:      r.PasswordHash,
		HashAlg:           r.HashAlg,
		HashParams:        r.HashParams,
		MustChangePassword: r.MustChangePassword,
	}
	
	if r.CredentialType != "" {
		cmd.CredentialType = userCredentialDomain.CredentialType(r.CredentialType)
	}
	if r.Status != "" {
		cmd.Status = userCredentialDomain.CredentialStatus(r.Status)
	}
	
	return cmd
}

func (r *UserCredentialUpdateRequestDTO) ToUpdateCmd() userCredentialAppCommands.UpdateUserCredentialCmd {
	cmd := userCredentialAppCommands.UpdateUserCredentialCmd{
		UserCredentialID: r.ID,
		MustChangePassword: r.MustChangePassword,
	}
	
	if r.Status != "" {
		cmd.Status = userCredentialDomain.CredentialStatus(r.Status)
	}
	
	return cmd
}
