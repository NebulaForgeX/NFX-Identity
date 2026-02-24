package user_images

import (
	dirErr "nfxid/errors/src/directory"

	"github.com/google/uuid"
)

func (ui *UserImage) Validate() error {
	if ui.UserID() == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if ui.ImageID() == uuid.Nil {
		return dirErr.ErrImageIDRequired
	}
	return nil
}
