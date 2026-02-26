package get

import (
	"context"
	"errors"

	imageErr "nfxidentity/errors/src/image"
	"nfxidentity/modules/image/domain/image_variants"
	"nfxidentity/modules/image/infrastructure/rdb/models"
	"nfxidentity/modules/image/infrastructure/repository/image_variants/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ImageVariant，实现 image_variants.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*image_variants.ImageVariant, error) {
	var m models.ImageVariant
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageErr.ErrImageVariantNotFound
		}
		return nil, err
	}
	return mapper.ImageVariantModelToDomain(&m), nil
}
