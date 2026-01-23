package respdto

import (
	"time"

	passwordHistoryAppResult "nfxid/modules/auth/application/password_history/results"

	"github.com/google/uuid"
)

type PasswordHistoryDTO struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	TenantID    uuid.UUID `json:"tenant_id"`
	PasswordHash string   `json:"password_hash"`
	HashAlg     *string   `json:"hash_alg,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// PasswordHistoryROToDTO converts application PasswordHistoryRO to response DTO
func PasswordHistoryROToDTO(v *passwordHistoryAppResult.PasswordHistoryRO) *PasswordHistoryDTO {
	if v == nil {
		return nil
	}

	return &PasswordHistoryDTO{
		ID:          v.ID,
		UserID:      v.UserID,
		TenantID:    v.TenantID,
		PasswordHash: v.PasswordHash,
		HashAlg:     v.HashAlg,
		CreatedAt:   v.CreatedAt,
	}
}
