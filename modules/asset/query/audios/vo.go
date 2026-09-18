package audios

import (
	"time"

	"github.com/google/uuid"
)

type AudioVO struct {
	ID              uuid.UUID `json:"id"`
	FilePath        string    `json:"file_path"`
	FileName        string    `json:"file_name"`
	FileSize        int64     `json:"file_size"`
	MimeType        string    `json:"mime_type"`
	DurationSeconds *float64  `json:"duration_seconds,omitempty"`
	UploaderID      uuid.UUID `json:"uploader_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
