package image_tags

import "github.com/google/uuid"

func (it *ImageTag) Validate() error {
	if it.ImageID() == uuid.Nil {
		return ErrImageIDRequired
	}
	if it.Tag() == "" {
		return ErrTagRequired
	}
	return nil
}
