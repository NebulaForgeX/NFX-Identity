package images

func (i *Image) Validate() error {
	if i.Filename() == "" {
		return ErrFilenameRequired
	}
	if i.OriginalFilename() == "" {
		return ErrOriginalFilenameRequired
	}
	if i.MimeType() == "" {
		return ErrMimeTypeRequired
	}
	if i.StoragePath() == "" {
		return ErrStoragePathRequired
	}
	if i.Size() <= 0 {
		return ErrSizeRequired
	}
	return nil
}
