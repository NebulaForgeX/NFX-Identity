package image_type

import (
	"errors"
	"strings"
)

var (
	ErrInvalidKey = errors.New("key cannot be empty")
)

func (it *ImageType) Validate() error {
	if strings.TrimSpace(it.state.Key) == "" {
		return ErrInvalidKey
	}
	return nil
}

