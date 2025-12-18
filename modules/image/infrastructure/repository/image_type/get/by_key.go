package get

import (
	"context"
	"errors"
	imageTypeDomain "nfxid/modules/image/domain/image_type"
	imageTypeDomainErrors "nfxid/modules/image/domain/image_type/errors"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/mapper"

	"gorm.io/gorm"
)

// ByKey 根据 Key 获取 ImageType，实现 imageTypeDomain.Get 接口
func (h *Handler) ByKey(ctx context.Context, key string) (*imageTypeDomain.ImageType, error) {
	var m models.ImageType
	if err := h.db.WithContext(ctx).Where("key = ?", key).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageTypeDomainErrors.ErrImageTypeNotFound
		}
		return nil, err
	}
	return mapper.ImageTypeModelToDomain(&m), nil
}
