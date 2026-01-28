package respdto

import (
	"time"

	userImageAppResult "nfxid/modules/directory/application/user_images/results"

	"github.com/google/uuid"
)

type UserImageDTO struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	ImageID      uuid.UUID  `json:"image_id"`
	DisplayOrder int         `json:"display_order"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

// UserImageROToDTO converts application UserImageRO to response DTO
func UserImageROToDTO(v *userImageAppResult.UserImageRO) *UserImageDTO {
	if v == nil {
		return nil
	}

	return &UserImageDTO{
		ID:           v.ID,
		UserID:       v.UserID,
		ImageID:      v.ImageID,
		DisplayOrder: v.DisplayOrder,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		DeletedAt:    v.DeletedAt,
	}
}

// UserImageListROToDTO converts list of UserImageRO to DTOs
func UserImageListROToDTO(results []userImageAppResult.UserImageRO) []UserImageDTO {
	dtos := make([]UserImageDTO, len(results))
	for i, v := range results {
		if dto := UserImageROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
