package user_images

import (
	"time"

	"github.com/google/uuid"
)

type NewUserImageParams struct {
	UserID       uuid.UUID
	ImageID      uuid.UUID
	DisplayOrder int
}

func NewUserImage(p NewUserImageParams) (*UserImage, error) {
	if err := validateUserImageParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserImageFromState(UserImageState{
		ID:           id,
		UserID:       p.UserID,
		ImageID:      p.ImageID,
		DisplayOrder: p.DisplayOrder,
		CreatedAt:    now,
		UpdatedAt:    now,
	}), nil
}

func NewUserImageFromState(st UserImageState) *UserImage {
	return &UserImage{state: st}
}

func validateUserImageParams(p NewUserImageParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.ImageID == uuid.Nil {
		return ErrImageIDRequired
	}
	return nil
}
