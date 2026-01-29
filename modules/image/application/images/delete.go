package images

import (
	"context"
	"os"
	"path/filepath"

	imageCommands "nfxid/modules/image/application/images/commands"
)

// DeleteImage 删除图片（先删磁盘文件，再删数据库记录）
func (s *Service) DeleteImage(ctx context.Context, cmd imageCommands.DeleteImageCmd) error {
	image, err := s.imageRepo.Get.ByID(ctx, cmd.ImageID)
	if err != nil {
		return err
	}
	fullPath := filepath.Join(s.storageBasePath, image.StoragePath())
	_ = os.Remove(fullPath)
	return s.imageRepo.Delete.ByID(ctx, cmd.ImageID)
}
