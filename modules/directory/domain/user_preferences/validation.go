package user_preferences

import (
	dirErr "nfxid/errors/src/directory"

	"github.com/google/uuid"
)

func (up *UserPreference) Validate() error {
	if up.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	return nil
}
