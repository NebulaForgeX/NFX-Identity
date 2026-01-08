package image_variants

import "github.com/google/uuid"

func (iv *ImageVariant) Validate() error {
	if iv.ImageID() == uuid.Nil {
		return ErrImageIDRequired
	}
	if iv.VariantKey() == "" {
		return ErrVariantKeyRequired
	}
	if iv.StoragePath() == "" {
		return ErrStoragePathRequired
	}
	return nil
}
