package videos

import (
	"time"

	asseterrs "nfxidentity/errors/src/asset"
)

func (e *Video) EnsureNotDeleted() error {
	if e.DeletedAt() != nil {
		return asseterrs.ErrVideoNotFound
	}
	return nil
}

func (e *Video) UpdateFilePath(newPath string) error {
	if err := e.EnsureNotDeleted(); err != nil {
		return err
	}
	if newPath == "" {
		return asseterrs.ErrVideoFilePathRequired
	}
	e.state.FilePath = newPath
	e.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (e *Video) UpdateFileMeta(fileName string, fileSize int64, mimeType string) error {
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

func (e *Video) Delete() error {
	if e.state.DeletedAt != nil {
		return nil
	}
	now := time.Now().UTC()
	e.state.DeletedAt = &now
	e.state.UpdatedAt = now
	return nil
}
