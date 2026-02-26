package sessions

import (
	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

func (s *Session) Validate() error {
	if s.SessionID() == "" {
		return authErr.ErrSessionIDRequired
	}
	if s.UserID() == uuid.Nil {
		return authErr.ErrUserIDRequired
	}
	if s.TenantID() == uuid.Nil {
		return authErr.ErrTenantIDRequired
	}
	if s.ExpiresAt().IsZero() {
		return authErr.ErrExpiresAtRequired
	}
	return nil
}
