package get

import (
	"context"
	"errors"
	"nfxid/modules/image/domain/images"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/images/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByTypeID 根据 TypeID 获取 Images，实现 images.Get 接口
func (h *Handler) ByTypeID(ctx context.Context, typeID uuid.UUID) ([]*images.Image, error) {
	var ms []models.Image
	if err := h.db.WithContext(ctx).Where("type_id = ?", typeID).Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*images.Image{}, nil
		}
		return nil, err
	}

	result := make([]*images.Image, len(ms))
	for i := range ms {
		result[i] = mapper.ImageModelToDomain(&ms[i])
	}
	return result, nil
}
