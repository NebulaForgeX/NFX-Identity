package images

import (
	asseterrs "nfxidentity/errors/src/asset"

	"github.com/google/uuid"
)

func validateNewImageParams(p NewImageParams) error {
	if p.UploaderID == uuid.Nil {
		return asseterrs.ErrImageUploaderIDRequired
	}
	if p.FilePath == "" {
		return asseterrs.ErrImageFilePathRequired
	}
	if p.FileName == "" {
		return asseterrs.ErrImageFileNameRequired
	}
	if p.FileSize <= 0 {
		return asseterrs.ErrImageFileSizeInvalid
	}
	if p.MimeType == "" {
		return asseterrs.ErrImageMimeTypeRequired
	}
	return nil
}
