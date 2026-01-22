package image_types

import (
	"context"
	imageTypeCommands "nfxid/modules/image/application/image_types/commands"
)

// UpdateImageType 更新图片类型
func (s *Service) UpdateImageType(ctx context.Context, cmd imageTypeCommands.UpdateImageTypeCmd) error {
	// Get domain entity
	imageType, err := s.imageTypeRepo.Get.ByID(ctx, cmd.ImageTypeID)
	if err != nil {
		return err
	}

	// Update domain entity using behavior
	if err := imageType.Update(cmd.Key, cmd.Description, cmd.MaxWidth, cmd.MaxHeight, cmd.AspectRatio); err != nil {
		return err
	}

	// Save to repository
	return s.imageTypeRepo.Update.Generic(ctx, imageType)
}
