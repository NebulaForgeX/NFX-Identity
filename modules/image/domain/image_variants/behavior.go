package image_variants

import (
	"time"
)

func (iv *ImageVariant) UpdateURL(url string) error {
	iv.state.URL = &url
	iv.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (iv *ImageVariant) Update(width, height *int, size *int64, mimeType, storagePath, url *string) error {
	if width != nil {
		iv.state.Width = width
	}
	if height != nil {
		iv.state.Height = height
	}
	if size != nil {
		iv.state.Size = size
	}
	if mimeType != nil {
		iv.state.MimeType = mimeType
	}
	if storagePath != nil {
		if *storagePath == "" {
			return ErrStoragePathRequired
		}
		iv.state.StoragePath = *storagePath
	}
	if url != nil {
		iv.state.URL = url
	}

	iv.state.UpdatedAt = time.Now().UTC()
	return nil
}
