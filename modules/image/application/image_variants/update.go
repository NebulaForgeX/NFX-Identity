package image_variants

import (
	"context"
	imageVariantCommands "nfxid/modules/image/application/image_variants/commands"
)

// UpdateImageVariant 更新图片变体
func (s *Service) UpdateImageVariant(ctx context.Context, cmd imageVariantCommands.UpdateImageVariantCmd) error {
	// Get domain entity
	imageVariant, err := s.imageVariantRepo.Get.ByID(ctx, cmd.ImageVariantID)
	if err != nil {
		return err
	}

	// Update domain entity using behavior
	if err := imageVariant.Update(cmd.Width, cmd.Height, cmd.Size, cmd.MimeType, cmd.StoragePath, cmd.URL); err != nil {
		return err
	}

	// Save to repository
	return s.imageVariantRepo.Update.Generic(ctx, imageVariant)
}

// UpdateImageVariantURL 更新图片变体URL
func (s *Service) UpdateImageVariantURL(ctx context.Context, cmd imageVariantCommands.UpdateImageVariantURLCmd) error {
	// Get domain entity
	imageVariant, err := s.imageVariantRepo.Get.ByID(ctx, cmd.ImageVariantID)
	if err != nil {
		return err
	}

	// Update URL using domain behavior
	if err := imageVariant.UpdateURL(cmd.URL); err != nil {
		return err
	}

	// Save to repository
	return s.imageVariantRepo.Update.UpdateURL(ctx, cmd.ImageVariantID, cmd.URL)
}
