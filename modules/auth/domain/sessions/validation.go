package sessions

import "github.com/google/uuid"

func (s *Session) Validate() error {
	if s.SessionID() == "" {
		return ErrSessionIDRequired
	}
	if s.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if s.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if s.ExpiresAt().IsZero() {
		return ErrExpiresAtRequired
	}
	return nil
}
