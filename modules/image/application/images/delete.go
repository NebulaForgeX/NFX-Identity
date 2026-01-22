package images

import (
	"context"
	imageCommands "nfxid/modules/image/application/images/commands"
)

// DeleteImage 删除图片
func (s *Service) DeleteImage(ctx context.Context, cmd imageCommands.DeleteImageCmd) error {
	// Delete from repository (hard delete)
	return s.imageRepo.Delete.ByID(ctx, cmd.ImageID)
}
