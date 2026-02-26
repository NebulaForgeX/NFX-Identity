package user_avatars

import (
	dirErr "nfxidentity/errors/src/directory"

	"github.com/google/uuid"
)

func (ua *UserAvatar) Validate() error {
	if ua.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if ua.ImageID() == uuid.Nil {
		return dirErr.ErrImageIDRequired
	}
	return nil
}
