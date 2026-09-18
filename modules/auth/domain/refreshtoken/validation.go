package refreshtoken

import (
	"strings"
	"time"

	authErr "nfxidentity/errors/src/auth"
)

func validateRefreshTokenIssue(now, expiresAt time.Time, tokenHash string) error {
	if strings.TrimSpace(tokenHash) == "" {
		return authErr.ErrRefreshTokenTokenHashRequired
	}
	if expiresAt.IsZero() {
		return authErr.ErrRefreshTokenExpiresAtRequired
	}
	if !expiresAt.After(now) {
		return authErr.ErrRefreshTokenExpiresAtInvalid
	}
	return nil
}
