package audios

import (
	asseterrs "nfxidentity/errors/src/asset"

	"github.com/google/uuid"
)

func validateNewAudioParams(p NewAudioParams) error {
	if p.UploaderID == uuid.Nil {
		return asseterrs.ErrAudioUploaderIDRequired
	}
	if p.FilePath == "" {
		return asseterrs.ErrAudioFilePathRequired
	}
	if p.FileName == "" {
		return asseterrs.ErrAudioFileNameRequired
	}
	if p.FileSize <= 0 {
		return asseterrs.ErrAudioFileSizeInvalid
	}
	if p.MimeType == "" {
		return asseterrs.ErrAudioMimeTypeRequired
	}
	return nil
}
