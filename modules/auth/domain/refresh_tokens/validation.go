package refresh_tokens

import "github.com/google/uuid"

func (rt *RefreshToken) Validate() error {
	if rt.TokenID() == "" {
		return ErrTokenIDRequired
	}
	if rt.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if rt.ExpiresAt().IsZero() {
		return ErrExpiresAtRequired
	}
	return nil
}
