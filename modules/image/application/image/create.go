package image

import (
	"context"
	imageDomain "nfxid/modules/image/domain/image"

	"github.com/google/uuid"
)

type CreateImageCmd struct {
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
}

func (s *Service) CreateImage(ctx context.Context, cmd CreateImageCmd) (*imageDomain.Image, error) {
	editable := imageDomain.ImageEditable{
		TypeID:           cmd.TypeID,
		UserID:           cmd.UserID,
		SourceDomain:     cmd.SourceDomain,
		Filename:         cmd.Filename,
		OriginalFilename: cmd.OriginalFilename,
		MimeType:         cmd.MimeType,
		Size:             cmd.Size,
		Width:            cmd.Width,
		Height:           cmd.Height,
		StoragePath:      cmd.StoragePath,
		URL:              cmd.URL,
		IsPublic:         cmd.IsPublic,
		Metadata:         cmd.Metadata,
	}

	img := imageDomain.NewImage(editable)
	if err := img.Validate(); err != nil {
		return nil, err
	}

	if err := s.imageRepo.Create(ctx, img); err != nil {
		return nil, err
	}

	return img, nil
}
