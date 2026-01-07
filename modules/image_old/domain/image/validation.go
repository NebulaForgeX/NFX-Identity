package image

import (
	"errors"
	"strings"
)

var (
	ErrInvalidFilename     = errors.New("filename cannot be empty")
	ErrInvalidStoragePath  = errors.New("storage path cannot be empty")
	ErrInvalidMimeType     = errors.New("mime type cannot be empty")
	ErrInvalidSize         = errors.New("size must be greater than 0")
)

func (i *Image) Validate() error {
	if strings.TrimSpace(i.state.Filename) == "" {
		return ErrInvalidFilename
	}
	if strings.TrimSpace(i.state.StoragePath) == "" {
		return ErrInvalidStoragePath
	}
	if strings.TrimSpace(i.state.MimeType) == "" {
		return ErrInvalidMimeType
	}
	if i.state.Size <= 0 {
		return ErrInvalidSize
	}
	return nil
}

