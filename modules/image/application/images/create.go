package images

import (
	"context"
	imageCommands "nfxid/modules/image/application/images/commands"
	imageDomain "nfxid/modules/image/domain/images"

	"github.com/google/uuid"
)

// CreateImage 创建图片
func (s *Service) CreateImage(ctx context.Context, cmd imageCommands.CreateImageCmd) (uuid.UUID, error) {
	// Create domain entity
	image, err := imageDomain.NewImage(imageDomain.NewImageParams{
		TypeID:          cmd.TypeID,
		UserID:          cmd.UserID,
		TenantID:        cmd.TenantID,
		AppID:           cmd.AppID,
		SourceDomain:    cmd.SourceDomain,
		Filename:        cmd.Filename,
		OriginalFilename: cmd.OriginalFilename,
		MimeType:        cmd.MimeType,
		Size:            cmd.Size,
		Width:           cmd.Width,
		Height:          cmd.Height,
		StoragePath:     cmd.StoragePath,
		URL:             cmd.URL,
		IsPublic:        cmd.IsPublic,
		Metadata:        cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.imageRepo.Create.New(ctx, image); err != nil {
		return uuid.Nil, err
	}

	return image.ID(), nil
}
