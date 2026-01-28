package user_avatars

import (
	"time"

	"github.com/google/uuid"
)

func (ua *UserAvatar) UpdateImageID(imageID uuid.UUID) error {
	if imageID == uuid.Nil {
		return ErrImageIDRequired
	}
	ua.state.ImageID = imageID
	ua.state.UpdatedAt = time.Now().UTC()
	return nil
}
