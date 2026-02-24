package image_variants

import (
	imageErr "nfxid/errors/src/image"
	"time"

	"github.com/google/uuid"
)

type NewImageVariantParams struct {
	ImageID     uuid.UUID
	VariantKey  string
	Width       *int
	Height      *int
	Size        *int64
	MimeType    *string
	StoragePath string
	URL         *string
}

func NewImageVariant(p NewImageVariantParams) (*ImageVariant, error) {
	if err := validateImageVariantParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewImageVariantFromState(ImageVariantState{
		ID:          id,
		ImageID:     p.ImageID,
		VariantKey:  p.VariantKey,
		Width:       p.Width,
		Height:      p.Height,
		Size:        p.Size,
		MimeType:    p.MimeType,
		StoragePath: p.StoragePath,
		URL:         p.URL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewImageVariantFromState(st ImageVariantState) *ImageVariant {
	return &ImageVariant{state: st}
}

func validateImageVariantParams(p NewImageVariantParams) error {
	if p.ImageID == uuid.Nil {
		return imageErr.ErrImageIDRequired
	}
	if p.VariantKey == "" {
		return imageErr.ErrVariantKeyRequired
	}
	if p.StoragePath == "" {
		return imageErr.ErrStoragePathRequired
	}
	return nil
}
