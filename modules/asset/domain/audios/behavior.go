package audios

import (
	"time"

	asseterrs "nfxidentity/errors/src/asset"
)

func (e *Audio) EnsureNotDeleted() error {
	if e.DeletedAt() != nil {
		return asseterrs.ErrAudioNotFound
	}
	return nil
}

func (e *Audio) UpdateFilePath(newPath string) error {
	if err := e.EnsureNotDeleted(); err != nil {
		return err
	}
	if newPath == "" {
		return asseterrs.ErrAudioFilePathRequired
	}
	e.state.FilePath = newPath
	e.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (e *Audio) UpdateFileMeta(fileName string, fileSize int64, mimeType string) error {
	if err := e.EnsureNotDeleted(); err != nil {
		return err
	}
	if fileName != "" {
		e.state.FileName = fileName
	}
	if fileSize > 0 {
		e.state.FileSize = fileSize
	}
	if mimeType != "" {
		e.state.MimeType = mimeType
	}
	e.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (e *Audio) Delete() error {
	if e.state.DeletedAt != nil {
		return nil
	}
	now := time.Now().UTC()
	e.state.DeletedAt = &now
	e.state.UpdatedAt = now
	return nil
}
