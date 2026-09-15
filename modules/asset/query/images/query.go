package images

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ImageVO struct {
	ID         uuid.UUID `json:"id"`
	FilePath   string    `json:"file_path"`
	FileName   string    `json:"file_name"`
	FileSize   int64     `json:"file_size"`
	MimeType   string    `json:"mime_type"`
	UploaderID uuid.UUID `json:"uploader_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Query struct{ Single Single }

type Single interface {
	ByID(ctx context.Context, id uuid.UUID) (*ImageVO, error)
	ListByUploader(ctx context.Context, uploaderID uuid.UUID) ([]ImageVO, error)
}
