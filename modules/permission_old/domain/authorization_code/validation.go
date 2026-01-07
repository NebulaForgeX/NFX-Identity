package authorization_code

import (
	authorizationCodeErrors "nfxid/modules/permission/domain/authorization_code/errors"
	"strings"
)

func (e *AuthorizationCodeEditable) Validate() error {
	if strings.TrimSpace(e.Code) == "" {
		return authorizationCodeErrors.ErrAuthorizationCodeCodeRequired
	}
	if len(e.Code) < 1 {
		return authorizationCodeErrors.ErrAuthorizationCodeCodeInvalid
	}
	if len(e.Code) > 255 {
		return authorizationCodeErrors.ErrAuthorizationCodeCodeInvalid
	}

	if e.MaxUses < 1 {
		return authorizationCodeErrors.ErrAuthorizationCodeMaxUsesInvalid
	}

	if e.UsedCount < 0 {
		return authorizationCodeErrors.ErrAuthorizationCodeMaxUsesInvalid
	}

	if e.UsedCount > e.MaxUses {
		return authorizationCodeErrors.ErrAuthorizationCodeMaxUsesInvalid
	}

	return nil
}
