package image_types

import (
	"time"
)

func (it *ImageType) Update(key string, description *string, maxWidth, maxHeight *int, aspectRatio *string) error {
	if key == "" {
		return ErrKeyRequired
	}

	it.state.Key = key
	if description != nil {
		it.state.Description = description
	}
	if maxWidth != nil {
		it.state.MaxWidth = maxWidth
	}
	if maxHeight != nil {
		it.state.MaxHeight = maxHeight
	}
	if aspectRatio != nil {
		it.state.AspectRatio = aspectRatio
	}

	it.state.UpdatedAt = time.Now().UTC()
	return nil
}
