package user_avatars

import (
	dirErr "nfxidentity/errors/src/directory"
	"time"

	"github.com/google/uuid"
)

func (ua *UserAvatar) UpdateImageID(imageID uuid.UUID) error {
	if imageID == uuid.Nil {
		return dirErr.ErrImageIDRequired
	}
	ua.state.ImageID = imageID
	ua.state.UpdatedAt = time.Now().UTC()
	return nil
}
