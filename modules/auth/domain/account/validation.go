package account

import (
	"nfxidentity/constants"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
)

func validateAccountStatus(status enums.AuthAccountStatus) error {
	if !constants.AuthAccountStatus.Valid(status) {
		return authErr.ErrAccountStatusInvalid
	}
	return nil
}

func validateSignupPlatform(platform enums.AuthSignupPlatform) error {
	if !constants.AuthSignupPlatform.Valid(platform) {
		return authErr.ErrAccountSignupPlatformInvalid
	}
	return nil
}
