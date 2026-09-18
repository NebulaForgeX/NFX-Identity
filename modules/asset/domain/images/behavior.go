package images

import (
	"time"

	asseterrs "nfxidentity/errors/src/asset"
)

func (e *Image) EnsureNotDeleted() error {
	if e.DeletedAt() != nil {
		return asseterrs.ErrImageNotFound
	}
	return nil
}

func (e *Image) UpdateFilePath(newPath string) error {
	if err := e.EnsureNotDeleted(); err != nil {
		return err
	}
	if newPath == "" {
		return asseterrs.ErrImageFilePathRequired
	}
	e.state.FilePath = newPath
	e.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (e *Image) UpdateFileMeta(fileName string, fileSize int64, mimeType string) error {
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

// Delete 软删：仅改领域状态；持久化由 Repo.Update 实现。
func (e *Image) Delete() error {
	if e.state.DeletedAt != nil {
		return nil
	}
	now := time.Now().UTC()
	e.state.DeletedAt = &now
	e.state.UpdatedAt = now
	return nil
}
