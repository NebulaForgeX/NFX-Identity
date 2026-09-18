package phone

import (
	"regexp"
	"strings"

	authErr "nfxidentity/errors/src/auth"
)

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

func validatePhoneNumber(num string) error {
	s := strings.TrimSpace(num)
	if s == "" {
		return authErr.ErrPhoneNumberRequired
	}
	if !phoneRegex.MatchString(s) {
		return authErr.ErrPhoneNumberFormatInvalid
	}
	return nil
}

func (e *PhoneEditable) Validate() error {
	return validatePhoneNumber(e.Phone)
}
