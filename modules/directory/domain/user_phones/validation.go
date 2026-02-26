package user_phones

import (
	dirErr "nfxidentity/errors/src/directory"

	"github.com/google/uuid"
)

func (up *UserPhone) Validate() error {
	if up.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if up.Phone() == "" {
		return dirErr.ErrPhoneRequired
	}
	return nil
}
