package image_types

import (
	"context"
	imageTypeCommands "nfxid/modules/image/application/image_types/commands"
)

// DeleteImageType 删除图片类型
func (s *Service) DeleteImageType(ctx context.Context, cmd imageTypeCommands.DeleteImageTypeCmd) error {
	// Delete from repository (hard delete)
	return s.imageTypeRepo.Delete.ByID(ctx, cmd.ImageTypeID)
}
