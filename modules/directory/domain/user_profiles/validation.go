package user_profiles

import (
	dirErr "nfxid/errors/src/directory"

	"github.com/google/uuid"
)

func (up *UserProfile) Validate() error {
	if up.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	return nil
}
