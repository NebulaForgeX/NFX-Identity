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

// ByID 根据 ID 获取 ImageVariant，实现 image_variants.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*image_variants.ImageVariant, error) {
	var m models.ImageVariant
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, image_variants.ErrImageVariantNotFound
		}
		return nil, err
	}
	return mapper.ImageVariantModelToDomain(&m), nil
}
