package respdto

import (
	"time"

	userAvatarAppResult "nfxid/modules/directory/application/user_avatars/results"

	"github.com/google/uuid"
)

type UserAvatarDTO struct {
	UserID    uuid.UUID `json:"user_id"`
	ImageID   uuid.UUID `json:"image_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserAvatarROToDTO converts application UserAvatarRO to response DTO
func UserAvatarROToDTO(v *userAvatarAppResult.UserAvatarRO) *UserAvatarDTO {
	if v == nil {
		return nil
	}

	return &UserAvatarDTO{
		UserID:    v.UserID,
		ImageID:   v.ImageID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
