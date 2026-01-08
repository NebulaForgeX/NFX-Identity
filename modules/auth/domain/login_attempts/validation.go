package login_attempts

import "github.com/google/uuid"

func (la *LoginAttempt) Validate() error {
	if la.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if la.Identifier() == "" {
		return ErrIdentifierRequired
	}
	return nil
}
