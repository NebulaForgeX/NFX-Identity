package images

import (
	"time"

	"github.com/google/uuid"
)

// ImageVO 读模型（来自 asset.ImagesActiveView）。
type ImageVO struct {
	ID         uuid.UUID `json:"id"`
	FilePath   string    `json:"file_path"`
	FileName   string    `json:"file_name"`
	FileSize   int64     `json:"file_size"`
	MimeType   string    `json:"mime_type"`
	Width      *int      `json:"width,omitempty"`
	Height     *int      `json:"height,omitempty"`
	AltText    *string   `json:"alt_text,omitempty"`
	UploaderID uuid.UUID `json:"uploader_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
