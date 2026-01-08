package images

import (
	"time"

	"github.com/google/uuid"
)

type NewImageParams struct {
	TypeID          *uuid.UUID
	UserID          *uuid.UUID
	TenantID        *uuid.UUID
	AppID           *uuid.UUID
	SourceDomain    *string
	Filename        string
	OriginalFilename string
	MimeType        string
	Size            int64
	Width           *int
	Height          *int
	StoragePath     string
	URL             *string
	IsPublic        bool
	Metadata        map[string]interface{}
}

func NewImage(p NewImageParams) (*Image, error) {
	if err := validateImageParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewImageFromState(ImageState{
		ID:              id,
		TypeID:          p.TypeID,
		UserID:          p.UserID,
		TenantID:        p.TenantID,
		AppID:           p.AppID,
		SourceDomain:    p.SourceDomain,
		Filename:        p.Filename,
		OriginalFilename: p.OriginalFilename,
		MimeType:        p.MimeType,
		Size:            p.Size,
		Width:           p.Width,
		Height:          p.Height,
		StoragePath:     p.StoragePath,
		URL:             p.URL,
		IsPublic:        p.IsPublic,
		Metadata:        p.Metadata,
		CreatedAt:       now,
		UpdatedAt:       now,
	}), nil
}

func NewImageFromState(st ImageState) *Image {
	return &Image{state: st}
}

func validateImageParams(p NewImageParams) error {
	if p.Filename == "" {
		return ErrFilenameRequired
	}
	if p.OriginalFilename == "" {
		return ErrOriginalFilenameRequired
	}
	if p.MimeType == "" {
		return ErrMimeTypeRequired
	}
	if p.StoragePath == "" {
		return ErrStoragePathRequired
	}
	if p.Size <= 0 {
		return ErrSizeRequired
	}
	return nil
}
