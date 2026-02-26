package image_types

import (
	imageErr "nfxidentity/errors/src/image"
)

func (it *ImageType) Validate() error {
	if it.Key() == "" {
		return imageErr.ErrKeyRequired
	}
	return nil
}
