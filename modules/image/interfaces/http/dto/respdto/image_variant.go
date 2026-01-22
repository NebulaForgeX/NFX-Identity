package respdto

import (
	"time"

	imageVariantAppResult "nfxid/modules/image/application/image_variants/results"

	"github.com/google/uuid"
)

type ImageVariantDTO struct {
	ID          uuid.UUID `json:"id"`
	ImageID     uuid.UUID `json:"image_id"`
	VariantKey  string    `json:"variant_key"`
	Width       *int      `json:"width,omitempty"`
	Height      *int      `json:"height,omitempty"`
	Size        *int64    `json:"size,omitempty"`
	MimeType    *string   `json:"mime_type,omitempty"`
	StoragePath string    `json:"storage_path"`
	URL         *string   `json:"url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ImageVariantROToDTO converts application ImageVariantRO to response DTO
func ImageVariantROToDTO(v *imageVariantAppResult.ImageVariantRO) *ImageVariantDTO {
	if v == nil {
		return nil
	}

	return &ImageVariantDTO{
		ID:          v.ID,
		ImageID:     v.ImageID,
		VariantKey:  v.VariantKey,
		Width:       v.Width,
		Height:      v.Height,
		Size:        v.Size,
		MimeType:    v.MimeType,
		StoragePath: v.StoragePath,
		URL:         v.URL,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
