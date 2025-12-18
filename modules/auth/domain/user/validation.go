package user

import (
	"nfxid/modules/auth/domain/user/errors"
	"regexp"
	"strings"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
)

func (e *UserEditable) Validate() error {
	if strings.TrimSpace(e.Username) == "" {
		return errors.ErrUserUsernameRequired
	}
	if len(e.Username) < 3 {
		return errors.ErrUserUsernameRequired
	}
	if len(e.Username) > 50 {
		return errors.ErrUserUsernameRequired
	}

	if strings.TrimSpace(e.Email) == "" {
		return errors.ErrUserEmailRequired
	}
	if !emailRegex.MatchString(e.Email) {
		return errors.ErrUserEmailRequired
	}

	if e.Phone != nil {
		if strings.TrimSpace(*e.Phone) == "" {
			return errors.ErrUserPhoneRequired
		}
		if !phoneRegex.MatchString(*e.Phone) {
			return errors.ErrUserPhoneRequired
		}
	}

	if e.Password == "" {
		return errors.ErrUserPasswordRequired
	}
	if len(e.Password) < 6 {
		return errors.ErrUserPasswordRequired
	}

	return nil
}
