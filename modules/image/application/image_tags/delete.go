package image_tags

import (
	"context"
	imageTagCommands "nfxid/modules/image/application/image_tags/commands"
)

// DeleteImageTag 删除图片标签
func (s *Service) DeleteImageTag(ctx context.Context, cmd imageTagCommands.DeleteImageTagCmd) error {
	// Delete from repository (hard delete)
	return s.imageTagRepo.Delete.ByID(ctx, cmd.ImageTagID)
}

// DeleteImageTagByImageIDAndTag 根据图片ID和标签删除图片标签
func (s *Service) DeleteImageTagByImageIDAndTag(ctx context.Context, cmd imageTagCommands.DeleteImageTagByImageIDAndTagCmd) error {
	// Delete from repository (hard delete)
	return s.imageTagRepo.Delete.ByImageIDAndTag(ctx, cmd.ImageID, cmd.Tag)
}
