package videos

import (
	"time"

	"github.com/google/uuid"
)

type NewVideoParams struct {
	ID              *uuid.UUID
	FilePath        string
	FileName        string
	FileSize        int64
	MimeType        string
	DurationSeconds *float64
	Width           *int
	Height          *int
	UploaderID      uuid.UUID
}

func NewVideo(p NewVideoParams) (*Video, error) {
	if err := validateNewVideoParams(p); err != nil {
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
	return NewVideoFromState(VideoState{
		ID:              id,
		FilePath:        p.FilePath,
		FileName:        p.FileName,
		FileSize:        p.FileSize,
		MimeType:        p.MimeType,
		DurationSeconds: p.DurationSeconds,
		Width:           p.Width,
		Height:          p.Height,
		UploaderID:      p.UploaderID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}), nil
}
