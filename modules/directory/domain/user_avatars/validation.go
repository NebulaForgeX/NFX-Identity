package user_avatars

import "github.com/google/uuid"

func (ua *UserAvatar) Validate() error {
	if ua.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if ua.ImageID() == uuid.Nil {
		return ErrImageIDRequired
	}
	return nil
}
