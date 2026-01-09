package get

import (
	"context"
	"errors"
	"nfxid/modules/image/domain/image_types"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/image_types/mapper"

	"gorm.io/gorm"
)

// All 获取所有 ImageTypes，实现 image_types.Get 接口
func (h *Handler) All(ctx context.Context) ([]*image_types.ImageType, error) {
	var ms []models.ImageType
	if err := h.db.WithContext(ctx).Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*image_types.ImageType{}, nil
		}
		return nil, err
	}
	
	result := make([]*image_types.ImageType, len(ms))
	for i := range ms {
		result[i] = mapper.ImageTypeModelToDomain(&ms[i])
	}
	return result, nil
}
