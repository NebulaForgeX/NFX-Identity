package views

import (
	"time"

	"github.com/google/uuid"
)

type ImageView struct {
	ID               uuid.UUID
	TypeID           *uuid.UUID
	UserID           *uuid.UUID
	SourceDomain     *string
	Filename         string
	OriginalFilename string
	MimeType         string
	Size             int64
	Width            *int
	Height           *int
	StoragePath      string
	URL              *string
	IsPublic         bool
	Metadata         map[string]interface{}
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

