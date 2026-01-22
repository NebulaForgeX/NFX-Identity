package respdto

import (
	"time"

	imageTagAppResult "nfxid/modules/image/application/image_tags/results"

	"github.com/google/uuid"
)

type ImageTagDTO struct {
	ID         uuid.UUID `json:"id"`
	ImageID    uuid.UUID `json:"image_id"`
	Tag        string    `json:"tag"`
	Confidence *float64  `json:"confidence,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// ImageTagROToDTO converts application ImageTagRO to response DTO
func ImageTagROToDTO(v *imageTagAppResult.ImageTagRO) *ImageTagDTO {
	if v == nil {
		return nil
	}

	return &ImageTagDTO{
		ID:         v.ID,
		ImageID:    v.ImageID,
		Tag:        v.Tag,
		Confidence: v.Confidence,
		CreatedAt:  v.CreatedAt,
	}
}
