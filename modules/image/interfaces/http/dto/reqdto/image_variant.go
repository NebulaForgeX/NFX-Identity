package reqdto

import (
	imageVariantAppCommands "nfxid/modules/image/application/image_variants/commands"

	"github.com/google/uuid"
)

type ImageVariantCreateRequestDTO struct {
	ImageID     uuid.UUID `json:"image_id" validate:"required,uuid"`
	VariantKey  string    `json:"variant_key" validate:"required"`
	Width       *int      `json:"width,omitempty"`
	Height      *int      `json:"height,omitempty"`
	Size        *int64    `json:"size,omitempty"`
	MimeType    *string   `json:"mime_type,omitempty"`
	StoragePath string    `json:"storage_path" validate:"required"`
	URL         *string   `json:"url,omitempty"`
}

type ImageVariantByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type ImageVariantUpdateRequestDTO struct {
	Width       *int    `json:"width,omitempty"`
	Height      *int    `json:"height,omitempty"`
	Size        *int64  `json:"size,omitempty"`
	MimeType    *string `json:"mime_type,omitempty"`
	StoragePath *string `json:"storage_path,omitempty"`
	URL         *string `json:"url,omitempty"`
}

type ImageVariantUpdateURLRequestDTO struct {
	URL string `json:"url" validate:"required"`
}

func (r *ImageVariantCreateRequestDTO) ToCreateCmd() imageVariantAppCommands.CreateImageVariantCmd {
	return imageVariantAppCommands.CreateImageVariantCmd{
		ImageID:     r.ImageID,
		VariantKey:  r.VariantKey,
		Width:       r.Width,
		Height:      r.Height,
		Size:        r.Size,
		MimeType:    r.MimeType,
		StoragePath: r.StoragePath,
		URL:         r.URL,
	}
}

func (r *ImageVariantUpdateRequestDTO) ToUpdateCmd(imageVariantID uuid.UUID) imageVariantAppCommands.UpdateImageVariantCmd {
	return imageVariantAppCommands.UpdateImageVariantCmd{
		ImageVariantID: imageVariantID,
		Width:          r.Width,
		Height:         r.Height,
		Size:           r.Size,
		MimeType:       r.MimeType,
		StoragePath:    r.StoragePath,
		URL:            r.URL,
	}
}
