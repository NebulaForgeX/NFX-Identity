package results

import (
	"time"

	"nfxid/modules/directory/domain/user_avatars"

	"github.com/google/uuid"
)

type UserAvatarRO struct {
	UserID    uuid.UUID
	ImageID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserAvatarMapper 将 Domain UserAvatar 转换为 Application UserAvatarRO
func UserAvatarMapper(ua *user_avatars.UserAvatar) UserAvatarRO {
	if ua == nil {
		return UserAvatarRO{}
	}

	return UserAvatarRO{
		UserID:    ua.UserID(),
		ImageID:   ua.ImageID(),
		CreatedAt: ua.CreatedAt(),
		UpdatedAt: ua.UpdatedAt(),
	}
}
