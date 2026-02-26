package get

import (
	"context"
	"errors"

	imageErr "nfxidentity/errors/src/image"
	"nfxidentity/modules/image/domain/image_types"
	"nfxidentity/modules/image/infrastructure/rdb/models"
	"nfxidentity/modules/image/infrastructure/repository/image_types/mapper"

	"gorm.io/gorm"
)

// ByKey 根据 Key 获取 ImageType，实现 image_types.Get 接口
func (h *Handler) ByKey(ctx context.Context, key string) (*image_types.ImageType, error) {
	var m models.ImageType
	if err := h.db.WithContext(ctx).Where("key = ?", key).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageErr.ErrImageTypeNotFound
		}
		return nil, err
	}
	return mapper.ImageTypeModelToDomain(&m), nil
}
