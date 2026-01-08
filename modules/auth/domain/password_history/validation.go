package password_history

import "github.com/google/uuid"

func (ph *PasswordHistory) Validate() error {
	if ph.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if ph.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if ph.PasswordHash() == "" {
		return ErrPasswordHashRequired
	}
	return nil
}
