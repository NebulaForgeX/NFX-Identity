package image

import (
	"time"

	"github.com/google/uuid"
)

func NewImage(editable ImageEditable) *Image {
	now := time.Now()
	return &Image{
		state: ImageState{
			ID:               uuid.New(),
			TypeID:           editable.TypeID,
			UserID:           editable.UserID,
			SourceDomain:     editable.SourceDomain,
			Filename:         editable.Filename,
			OriginalFilename: editable.OriginalFilename,
			MimeType:         editable.MimeType,
			Size:             editable.Size,
			Width:            editable.Width,
			Height:           editable.Height,
			StoragePath:      editable.StoragePath,
			URL:              editable.URL,
			IsPublic:         editable.IsPublic,
			Metadata:         editable.Metadata,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}
}

func NewImageWithID(id uuid.UUID, editable ImageEditable) *Image {
	now := time.Now()
	return &Image{
		state: ImageState{
			ID:               id,
			TypeID:           editable.TypeID,
			UserID:           editable.UserID,
			SourceDomain:     editable.SourceDomain,
			Filename:         editable.Filename,
			OriginalFilename: editable.OriginalFilename,
			MimeType:         editable.MimeType,
			Size:             editable.Size,
			Width:            editable.Width,
			Height:           editable.Height,
			StoragePath:      editable.StoragePath,
			URL:              editable.URL,
			IsPublic:         editable.IsPublic,
			Metadata:         editable.Metadata,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}
}

func NewImageFromState(state ImageState) *Image {
	return &Image{
		state: state,
	}
}

