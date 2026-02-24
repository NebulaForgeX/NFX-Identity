package user_avatars

import (
	dirErr "nfxid/errors/src/directory"
	"time"

	"github.com/google/uuid"
)

type NewUserAvatarParams struct {
	UserID  uuid.UUID
	ImageID uuid.UUID
}

func NewUserAvatar(p NewUserAvatarParams) (*UserAvatar, error) {
	if err := validateUserAvatarParams(p); err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserAvatarFromState(UserAvatarState{
		UserID:    p.UserID,
		ImageID:   p.ImageID,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

func NewUserAvatarFromState(st UserAvatarState) *UserAvatar {
	return &UserAvatar{state: st}
}

func validateUserAvatarParams(p NewUserAvatarParams) error {
	if p.UserID == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if p.ImageID == uuid.Nil {
		return dirErr.ErrImageIDRequired
	}
	return nil
}
