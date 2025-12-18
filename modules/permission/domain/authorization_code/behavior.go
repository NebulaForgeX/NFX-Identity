package authorization_code

import (
	authorizationCodeErrors "nfxid/modules/permission/domain/authorization_code/errors"
	"time"
)

func (ac *AuthorizationCode) Use() error {
	if ac.IsDeleted() {
		return authorizationCodeErrors.ErrAuthorizationCodeNotFound
	}

	if !ac.IsActive() {
		return authorizationCodeErrors.ErrAuthorizationCodeInactive
	}

	if ac.state.Editable.UsedCount >= ac.state.Editable.MaxUses {
		return authorizationCodeErrors.ErrAuthorizationCodeAlreadyUsed
	}

	if ac.ExpiresAt() != nil && time.Now().UTC().After(*ac.ExpiresAt()) {
		return authorizationCodeErrors.ErrAuthorizationCodeExpired
	}

	ac.state.Editable.UsedCount++
	ac.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ac *AuthorizationCode) Activate() error {
	if ac.IsDeleted() {
		return authorizationCodeErrors.ErrAuthorizationCodeNotFound
	}

	ac.state.IsActive = true
	ac.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ac *AuthorizationCode) Deactivate() error {
	if ac.IsDeleted() {
		return authorizationCodeErrors.ErrAuthorizationCodeNotFound
	}

	ac.state.IsActive = false
	ac.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ac *AuthorizationCode) Delete() error {
	if ac.IsDeleted() {
		return nil // idempotent
	}

	now := time.Now().UTC()
	ac.state.DeletedAt = &now
	ac.state.UpdatedAt = now
	return nil
}
