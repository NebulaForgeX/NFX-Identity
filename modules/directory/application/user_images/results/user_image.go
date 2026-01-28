package results

import (
	"time"

	"nfxid/modules/directory/domain/user_images"

	"github.com/google/uuid"
)

type UserImageRO struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ImageID      uuid.UUID
	DisplayOrder int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// UserImageMapper 将 Domain UserImage 转换为 Application UserImageRO
func UserImageMapper(ui *user_images.UserImage) UserImageRO {
	if ui == nil {
		return UserImageRO{}
	}

	return UserImageRO{
		ID:           ui.ID(),
		UserID:       ui.UserID(),
		ImageID:      ui.ImageID(),
		DisplayOrder: ui.DisplayOrder(),
		CreatedAt:    ui.CreatedAt(),
		UpdatedAt:    ui.UpdatedAt(),
		DeletedAt:    ui.DeletedAt(),
	}
}
