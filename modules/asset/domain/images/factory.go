package images

import (
	"time"

	"github.com/google/uuid"
)

type NewImageParams struct {
	ID         *uuid.UUID
	FilePath   string
	FileName   string
	FileSize   int64
	MimeType   string
	Width      *int
	Height     *int
	AltText    *string
	UploaderID uuid.UUID
}

func NewImage(p NewImageParams) (*Image, error) {
	if err := validateNewImageParams(p); err != nil {
		return nil, err
	}
	id := uuid.Nil
	if p.ID != nil {
		id = *p.ID
	}
	if id == uuid.Nil {
		var err error
		id, err = uuid.NewV7()
		if err != nil {
			return nil, err
		}
	}
	now := time.Now().UTC()
	return NewImageFromState(ImageState{
		ID:         id,
		FilePath:   p.FilePath,
		FileName:   p.FileName,
		FileSize:   p.FileSize,
		MimeType:   p.MimeType,
		Width:      p.Width,
		Height:     p.Height,
		AltText:    p.AltText,
		UploaderID: p.UploaderID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}), nil
}
