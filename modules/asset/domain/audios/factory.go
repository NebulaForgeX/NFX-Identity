package audios

import (
	"time"

	"github.com/google/uuid"
)

type NewAudioParams struct {
	ID              *uuid.UUID
	FilePath        string
	FileName        string
	FileSize        int64
	MimeType        string
	DurationSeconds *float64
	UploaderID      uuid.UUID
}

func NewAudio(p NewAudioParams) (*Audio, error) {
	if err := validateNewAudioParams(p); err != nil {
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
	return NewAudioFromState(AudioState{
		ID:              id,
		FilePath:        p.FilePath,
		FileName:        p.FileName,
		FileSize:        p.FileSize,
		MimeType:        p.MimeType,
		DurationSeconds: p.DurationSeconds,
		UploaderID:      p.UploaderID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}), nil
}
