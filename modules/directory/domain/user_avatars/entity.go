package user_avatars

import (
	"time"

	"github.com/google/uuid"
)

type UserAvatar struct {
	state UserAvatarState
}

type UserAvatarState struct {
	UserID    uuid.UUID
	ImageID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ua *UserAvatar) UserID() uuid.UUID    { return ua.state.UserID }
func (ua *UserAvatar) ImageID() uuid.UUID   { return ua.state.ImageID }
func (ua *UserAvatar) CreatedAt() time.Time { return ua.state.CreatedAt }
func (ua *UserAvatar) UpdatedAt() time.Time { return ua.state.UpdatedAt }
