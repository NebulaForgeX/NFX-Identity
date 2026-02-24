package images

import (
	imageErr "nfxid/errors/src/image"
)

func (i *Image) Validate() error {
	if i.Filename() == "" {
		return imageErr.ErrFilenameRequired
	}
	if i.OriginalFilename() == "" {
		return imageErr.ErrOriginalFilenameRequired
	}
	if i.MimeType() == "" {
		return imageErr.ErrMimeTypeRequired
	}
	if i.StoragePath() == "" {
		return imageErr.ErrStoragePathRequired
	}
	if i.Size() <= 0 {
		return imageErr.ErrSizeRequired
	}
	return nil
}
