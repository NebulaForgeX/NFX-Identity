package image_variants

import (
	imageErr "nfxid/errors/src/image"

	"github.com/google/uuid"
)

func (iv *ImageVariant) Validate() error {
	if iv.ImageID() == uuid.Nil {
		return imageErr.ErrImageIDRequired
	}
	if iv.VariantKey() == "" {
		return imageErr.ErrVariantKeyRequired
	}
	if iv.StoragePath() == "" {
		return imageErr.ErrStoragePathRequired
	}
	return nil
}
