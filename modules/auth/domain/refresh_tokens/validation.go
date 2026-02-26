package refresh_tokens

import (
	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

func (rt *RefreshToken) Validate() error {
	if rt.TokenID() == "" {
		return authErr.ErrTokenIDRequired
	}
	if rt.UserID() == uuid.Nil {
		return authErr.ErrUserIDRequired
	}
	if rt.ExpiresAt().IsZero() {
		return authErr.ErrExpiresAtRequired
	}
	return nil
}
