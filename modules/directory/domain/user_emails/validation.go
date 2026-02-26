package user_emails

import (
	dirErr "nfxidentity/errors/src/directory"

	"github.com/google/uuid"
)

func (ue *UserEmail) Validate() error {
	if ue.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if ue.Email() == "" {
		return dirErr.ErrEmailRequired
	}
	return nil
}
