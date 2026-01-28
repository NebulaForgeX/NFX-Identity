package user_images

import "github.com/google/uuid"

func (ui *UserImage) Validate() error {
	if ui.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if ui.ImageID() == uuid.Nil {
		return ErrImageIDRequired
	}
	return nil
}
