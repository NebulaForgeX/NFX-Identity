package files

import (
	"time"

	"github.com/google/uuid"
)

type NewFileParams struct {
	ID         *uuid.UUID
	FilePath   string
	FileName   string
	FileSize   int64
	MimeType   string
	UploaderID uuid.UUID
}

func NewFile(p NewFileParams) (*File, error) {
	if err := validateNewFileParams(p); err != nil {
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
	return NewFileFromState(FileState{
		ID:         id,
		FilePath:   p.FilePath,
		FileName:   p.FileName,
		FileSize:   p.FileSize,
		MimeType:   p.MimeType,
		UploaderID: p.UploaderID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}), nil
}
