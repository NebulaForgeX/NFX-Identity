package image_variants

import (
	"context"
	imageVariantCommands "nfxid/modules/image/application/image_variants/commands"
	imageVariantDomain "nfxid/modules/image/domain/image_variants"

	"github.com/google/uuid"
)

// CreateImageVariant 创建图片变体
func (s *Service) CreateImageVariant(ctx context.Context, cmd imageVariantCommands.CreateImageVariantCmd) (uuid.UUID, error) {
	// Create domain entity
	imageVariant, err := imageVariantDomain.NewImageVariant(imageVariantDomain.NewImageVariantParams{
		ImageID:     cmd.ImageID,
		VariantKey:  cmd.VariantKey,
		Width:       cmd.Width,
		Height:      cmd.Height,
		Size:        cmd.Size,
		MimeType:    cmd.MimeType,
		StoragePath: cmd.StoragePath,
		URL:         cmd.URL,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.imageVariantRepo.Create.New(ctx, imageVariant); err != nil {
		return uuid.Nil, err
	}

	return imageVariant.ID(), nil
}
