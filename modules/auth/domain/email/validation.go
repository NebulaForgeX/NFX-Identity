package email

import (
	"regexp"
	"strings"

	authErr "nfxidentity/errors/src/auth"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func validateEmailAddress(addr string) error {
	if strings.TrimSpace(addr) == "" {
		return authErr.ErrEmailAddressRequired
	}
	if !emailRegex.MatchString(strings.TrimSpace(addr)) {
		return authErr.ErrEmailAddressFormatInvalid
	}
	return nil
}

func (e *EmailEditable) Validate() error {
	return validateEmailAddress(e.Email)
}
