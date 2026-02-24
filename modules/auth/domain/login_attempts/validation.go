package login_attempts

import (
	authErr "nfxid/errors/src/auth"
)

func (la *LoginAttempt) Validate() error {
	if la.Identifier() == "" {
		return authErr.ErrIdentifierRequired
	}
	return nil
}
