package refreshtoken

import (
	"time"

	authErr "nfxidentity/errors/src/auth"
)

func (t *RefreshToken) EnsureNotDeleted() error {
	if t.DeletedAt() != nil {
		return authErr.ErrRefreshTokenNotFound
	}
	return nil
}

func (t *RefreshToken) Revoke(at time.Time) error {
	if err := t.EnsureNotDeleted(); err != nil {
		return err
	}
	if t.state.RevokedAt != nil {
		return authErr.ErrRefreshTokenAlreadyRevoked
	}
	x := at.UTC()
	t.state.RevokedAt = &x
	return nil
}

func (t *RefreshToken) Delete() error {
	if t.DeletedAt() != nil {
		return nil
	}
	now := time.Now().UTC()
	t.state.DeletedAt = &now
	return nil
}
