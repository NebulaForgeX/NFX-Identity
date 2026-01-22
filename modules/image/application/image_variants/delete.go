package image_variants

import (
	"context"
	imageVariantCommands "nfxid/modules/image/application/image_variants/commands"
)

// DeleteImageVariant 删除图片变体
func (s *Service) DeleteImageVariant(ctx context.Context, cmd imageVariantCommands.DeleteImageVariantCmd) error {
	// Delete from repository (hard delete)
	return s.imageVariantRepo.Delete.ByID(ctx, cmd.ImageVariantID)
}

// DeleteImageVariantByImageIDAndVariantKey 根据图片ID和变体Key删除图片变体
func (s *Service) DeleteImageVariantByImageIDAndVariantKey(ctx context.Context, cmd imageVariantCommands.DeleteImageVariantByImageIDAndVariantKeyCmd) error {
	// Delete from repository (hard delete)
	return s.imageVariantRepo.Delete.ByImageIDAndVariantKey(ctx, cmd.ImageID, cmd.VariantKey)
}
