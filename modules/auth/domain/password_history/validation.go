package password_history

import (
	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

func (ph *PasswordHistory) Validate() error {
	if ph.UserID() == uuid.Nil {
		return authErr.ErrUserIDRequired
	}
	if ph.TenantID() == uuid.Nil {
		return authErr.ErrTenantIDRequired
	}
	if ph.PasswordHash() == "" {
		return authErr.ErrPasswordHashRequired
	}
	return nil
}
