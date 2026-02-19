package reqdto

import (
	passwordHistoryAppCommands "nfxid/modules/auth/application/password_history/commands"

	"github.com/google/uuid"
)

type PasswordHistoryCreateRequestDTO struct {
	UserID       uuid.UUID `json:"user_id" validate:"required"`
	TenantID     uuid.UUID `json:"tenant_id" validate:"required"`
	PasswordHash string    `json:"password_hash" validate:"required"`
	HashAlg      *string   `json:"hash_alg,omitempty"`
}

type PasswordHistoryByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *PasswordHistoryCreateRequestDTO) ToCreateCmd() passwordHistoryAppCommands.CreatePasswordHistoryCmd {
	return passwordHistoryAppCommands.CreatePasswordHistoryCmd{
		UserID:       r.UserID,
		TenantID:     r.TenantID,
		PasswordHash: r.PasswordHash,
		HashAlg:      r.HashAlg,
	}
}
