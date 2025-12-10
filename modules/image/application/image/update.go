package image

import (
	"context"
	imageDomain "nebulaid/modules/image/domain/image"

	"github.com/google/uuid"
)

type UpdateImageCmd struct {
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
}

func (s *Service) UpdateImage(ctx context.Context, cmd UpdateImageCmd) error {
	img, err := s.imageRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

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

	img.Update(editable)
	if err := img.Validate(); err != nil {
		return err
	}

	return s.imageRepo.Update(ctx, img)
}
