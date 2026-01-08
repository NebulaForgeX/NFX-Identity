package password_history

import (
	"time"

	"github.com/google/uuid"
)

type NewPasswordHistoryParams struct {
	UserID      uuid.UUID
	TenantID    uuid.UUID
	PasswordHash string
	HashAlg     *string
}

func NewPasswordHistory(p NewPasswordHistoryParams) (*PasswordHistory, error) {
	if err := validatePasswordHistoryParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewPasswordHistoryFromState(PasswordHistoryState{
		ID:           id,
		UserID:       p.UserID,
		TenantID:     p.TenantID,
		PasswordHash: p.PasswordHash,
		HashAlg:      p.HashAlg,
		CreatedAt:    now,
	}), nil
}

func NewPasswordHistoryFromState(st PasswordHistoryState) *PasswordHistory {
	return &PasswordHistory{state: st}
}

func validatePasswordHistoryParams(p NewPasswordHistoryParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.PasswordHash == "" {
		return ErrPasswordHashRequired
	}
	return nil
}
