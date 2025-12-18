package get

import (
	"context"
	"errors"
	imageDomain "nfxid/modules/image/domain/image"
	imageDomainErrors "nfxid/modules/image/domain/image/errors"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/mapper"

	"gorm.io/gorm"
)

// ByStoragePath 根据 StoragePath 获取 Image，实现 imageDomain.Get 接口
func (h *Handler) ByStoragePath(ctx context.Context, storagePath string) (*imageDomain.Image, error) {
	var m models.Image
	if err := h.db.WithContext(ctx).Where("storage_path = ?", storagePath).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageDomainErrors.ErrImageNotFound
		}
		return nil, err
	}
	return mapper.ImageModelToDomain(&m), nil
}
