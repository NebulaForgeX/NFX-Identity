package videos

import (
	asseterrs "nfxidentity/errors/src/asset"

	"github.com/google/uuid"
)

func validateNewVideoParams(p NewVideoParams) error {
	if p.UploaderID == uuid.Nil {
		return asseterrs.ErrVideoUploaderIDRequired
	}
	if p.FilePath == "" {
		return asseterrs.ErrVideoFilePathRequired
	}
	if p.FileName == "" {
		return asseterrs.ErrVideoFileNameRequired
	}
	if p.FileSize <= 0 {
		return asseterrs.ErrVideoFileSizeInvalid
	}
	if p.MimeType == "" {
		return asseterrs.ErrVideoMimeTypeRequired
	}
	return nil
}
