package image_tags

import (
	imageErr "nfxid/errors/src/image"

	"github.com/google/uuid"
)

func (it *ImageTag) Validate() error {
	if it.ImageID() == uuid.Nil {
		return imageErr.ErrImageIDRequired
	}
	if it.Tag() == "" {
		return imageErr.ErrTagRequired
	}
	return nil
}
