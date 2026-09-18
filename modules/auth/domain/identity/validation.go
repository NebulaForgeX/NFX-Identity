package identity

import (
	"strings"

	"nfxidentity/constants"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
)

func validateIdentityProvider(p enums.AuthIdentityProvider) error {
	if !constants.AuthIdentityProvider.Valid(p) {
		return authErr.ErrIdentityProviderInvalid
	}
	return nil
}

func validateProviderSubject(subject string) error {
	if strings.TrimSpace(subject) == "" {
		return authErr.ErrIdentityProviderSubjectRequired
	}
	return nil
}

func validatePasswordHashForProvider(p enums.AuthIdentityProvider, hash *string) error {
	if p != enums.AuthIdentityProviderPassword {
		return nil
	}
	if hash == nil || strings.TrimSpace(*hash) == "" {
		return authErr.ErrIdentityPasswordHashRequired
	}
	return nil
}

func (e *IdentityEditable) Validate() error {
	if err := validateIdentityProvider(e.IdentityProvider); err != nil {
		return err
	}
	if err := validateProviderSubject(e.ProviderSubject); err != nil {
		return err
	}
	return validatePasswordHashForProvider(e.IdentityProvider, e.PasswordHash)
}
