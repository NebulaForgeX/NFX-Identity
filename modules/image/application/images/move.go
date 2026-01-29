package images

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"nfxid/constants"
	imageAppCommands "nfxid/modules/image/application/images/commands"
	imageAppResult "nfxid/modules/image/application/images/results"

	"github.com/google/uuid"
)

// MoveImage 移动图片（从 tmp 移动到目标目录）
// 流程：
// 1. 获取图片信息
// 2. 验证图片是否在 tmp 目录
// 3. 复制文件到目标目录
// 4. 更新数据库中的存储路径
// 5. 删除原 tmp 文件
func (s *Service) MoveImage(ctx context.Context, imageID uuid.UUID, targetType string) (imageAppResult.ImageRO, error) {
	// 获取图片信息
	image, err := s.imageRepo.Get.ByID(ctx, imageID)
	if err != nil {
		return imageAppResult.ImageRO{}, fmt.Errorf("image not found: %w", err)
	}

	// 验证图片是否在 tmp 目录
	if !strings.Contains(image.StoragePath(), constants.StoragePathTmp) {
		// 图片不在 tmp 目录，直接返回当前状态
		return imageAppResult.ImageRO{
			ID:               image.ID(),
			TypeID:           image.TypeID(),
			UserID:           image.UserID(),
			TenantID:         image.TenantID(),
			AppID:            image.AppID(),
			SourceDomain:     image.SourceDomain(),
			Filename:         image.Filename(),
			OriginalFilename: image.OriginalFilename(),
			MimeType:         image.MimeType(),
			Size:             image.Size(),
			Width:            image.Width(),
			Height:           image.Height(),
			StoragePath:      image.StoragePath(),
			URL:              image.URL(),
			IsPublic:         image.IsPublic(),
			Metadata:         image.Metadata(),
			CreatedAt:        image.CreatedAt(),
			UpdatedAt:        image.UpdatedAt(),
			DeletedAt:        image.DeletedAt(),
		}, nil
	}

	// 构建新路径
	var targetPath string
	switch targetType {
	case "avatar":
		targetPath = constants.StoragePathAvatar
	case "background":
		targetPath = constants.StoragePathBackground
	default:
		return imageAppResult.ImageRO{}, fmt.Errorf("invalid target type: %s", targetType)
	}

	// 获取文件名
	filename := filepath.Base(image.StoragePath())
	newStoragePath := filepath.Join(targetPath, filename)

	oldFullPath := filepath.Join(s.storageBasePath, image.StoragePath())
	newFullPath := filepath.Join(s.storageBasePath, newStoragePath)

	// 确保目标目录存在
	if err := os.MkdirAll(filepath.Dir(newFullPath), 0755); err != nil {
		return imageAppResult.ImageRO{}, fmt.Errorf("failed to create target directory: %w", err)
	}

	// 复制文件
	if err := copyFile(oldFullPath, newFullPath); err != nil {
		return imageAppResult.ImageRO{}, fmt.Errorf("failed to copy file: %w", err)
	}

	// 更新数据库中的存储路径
	updateCmd := imageAppCommands.UpdateImageCmd{
		ImageID:     imageID,
		StoragePath: &newStoragePath,
	}
	if err := s.UpdateImage(ctx, updateCmd); err != nil {
		// 回滚：删除已复制的文件
		os.Remove(newFullPath)
		return imageAppResult.ImageRO{}, fmt.Errorf("failed to update storage path: %w", err)
	}

	// 删除原 tmp 文件
	os.Remove(oldFullPath)

	// 获取更新后的图片信息
	return s.GetImage(ctx, imageID)
}

// copyFile 复制文件
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
