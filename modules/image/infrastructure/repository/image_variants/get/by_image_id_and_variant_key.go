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

// ByImageIDAndVariantKey 根据 ImageID 和 VariantKey 获取 ImageVariant，实现 image_variants.Get 接口
func (h *Handler) ByImageIDAndVariantKey(ctx context.Context, imageID uuid.UUID, variantKey string) (*image_variants.ImageVariant, error) {
	var m models.ImageVariant
	if err := h.db.WithContext(ctx).
		Where("image_id = ? AND variant_key = ?", imageID, variantKey).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, image_variants.ErrImageVariantNotFound
		}
		return nil, err
	}
	return mapper.ImageVariantModelToDomain(&m), nil
}
