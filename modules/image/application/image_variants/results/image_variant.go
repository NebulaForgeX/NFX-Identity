package results

import (
	"time"

	"nfxid/modules/image/domain/image_variants"

	"github.com/google/uuid"
)

type ImageVariantRO struct {
	ID          uuid.UUID
	ImageID     uuid.UUID
	VariantKey  string
	Width       *int
	Height      *int
	Size        *int64
	MimeType    *string
	StoragePath string
	URL         *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ImageVariantMapper 将 Domain ImageVariant 转换为 Application ImageVariantRO
func ImageVariantMapper(iv *image_variants.ImageVariant) ImageVariantRO {
	if iv == nil {
		return ImageVariantRO{}
	}

	return ImageVariantRO{
		ID:          iv.ID(),
		ImageID:     iv.ImageID(),
		VariantKey:  iv.VariantKey(),
		Width:       iv.Width(),
		Height:      iv.Height(),
		Size:        iv.Size(),
		MimeType:    iv.MimeType(),
		StoragePath: iv.StoragePath(),
		URL:         iv.URL(),
		CreatedAt:   iv.CreatedAt(),
		UpdatedAt:   iv.UpdatedAt(),
	}
}
