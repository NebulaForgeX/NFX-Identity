package get

import (
	"context"
	"errors"
	"nfxid/modules/image/domain/image_variants"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/image_variants/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByImageID 根据 ImageID 获取 ImageVariants，实现 image_variants.Get 接口
func (h *Handler) ByImageID(ctx context.Context, imageID uuid.UUID) ([]*image_variants.ImageVariant, error) {
	var ms []models.ImageVariant
	if err := h.db.WithContext(ctx).Where("image_id = ?", imageID).Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*image_variants.ImageVariant{}, nil
		}
		return nil, err
	}
	
	result := make([]*image_variants.ImageVariant, len(ms))
	for i := range ms {
		result[i] = mapper.ImageVariantModelToDomain(&ms[i])
	}
	return result, nil
}
