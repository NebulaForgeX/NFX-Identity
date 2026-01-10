package results

import (
	"time"

	"nfxid/modules/auth/domain/password_history"

	"github.com/google/uuid"
)

type PasswordHistoryRO struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	TenantID    uuid.UUID
	PasswordHash string
	HashAlg     *string
	CreatedAt   time.Time
}

// PasswordHistoryMapper 将 Domain PasswordHistory 转换为 Application PasswordHistoryRO
func PasswordHistoryMapper(ph *password_history.PasswordHistory) PasswordHistoryRO {
	if ph == nil {
		return PasswordHistoryRO{}
	}

	return PasswordHistoryRO{
		ID:          ph.ID(),
		UserID:      ph.UserID(),
		TenantID:    ph.TenantID(),
		PasswordHash: ph.PasswordHash(),
		HashAlg:     ph.HashAlg(),
		CreatedAt:   ph.CreatedAt(),
	}
}
