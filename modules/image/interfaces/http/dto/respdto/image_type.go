package respdto

import (
	"time"

	imageTypeAppResult "nfxid/modules/image/application/image_types/results"

	"github.com/google/uuid"
)

type ImageTypeDTO struct {
	ID          uuid.UUID `json:"id"`
	Key         string    `json:"key"`
	Description *string   `json:"description,omitempty"`
	MaxWidth    *int      `json:"max_width,omitempty"`
	MaxHeight   *int      `json:"max_height,omitempty"`
	AspectRatio *string   `json:"aspect_ratio,omitempty"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ImageTypeROToDTO converts application ImageTypeRO to response DTO
func ImageTypeROToDTO(v *imageTypeAppResult.ImageTypeRO) *ImageTypeDTO {
	if v == nil {
		return nil
	}

	return &ImageTypeDTO{
		ID:          v.ID,
		Key:         v.Key,
		Description: v.Description,
		MaxWidth:    v.MaxWidth,
		MaxHeight:   v.MaxHeight,
		AspectRatio: v.AspectRatio,
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
