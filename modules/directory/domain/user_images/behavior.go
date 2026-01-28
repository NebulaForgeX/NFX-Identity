package user_images

import (
	"time"

	"github.com/google/uuid"
)

func (ui *UserImage) UpdateDisplayOrder(displayOrder int) error {
	if ui.DeletedAt() != nil {
		return ErrUserImageNotFound
	}
	ui.state.DisplayOrder = displayOrder
	ui.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ui *UserImage) UpdateImageID(imageID uuid.UUID) error {
	if ui.DeletedAt() != nil {
		return ErrUserImageNotFound
	}
	if imageID == uuid.Nil {
		return ErrImageIDRequired
	}
	ui.state.ImageID = imageID
	ui.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ui *UserImage) Delete() error {
	if ui.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	ui.state.DeletedAt = &now
	ui.state.UpdatedAt = now
	return nil
}
