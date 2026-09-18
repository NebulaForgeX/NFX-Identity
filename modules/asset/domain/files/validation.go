package files

import (
	asseterrs "nfxidentity/errors/src/asset"

	"github.com/google/uuid"
)

func validateNewFileParams(p NewFileParams) error {
	if p.UploaderID == uuid.Nil {
		return asseterrs.ErrFileUploaderIDRequired
	}
	if p.FilePath == "" {
		return asseterrs.ErrFileFilePathRequired
	}
	if p.FileName == "" {
		return asseterrs.ErrFileFileNameRequired
	}
	if p.FileSize <= 0 {
		return asseterrs.ErrFileFileSizeInvalid
	}
	if p.MimeType == "" {
		return asseterrs.ErrFileMimeTypeRequired
	}
	return nil
}
