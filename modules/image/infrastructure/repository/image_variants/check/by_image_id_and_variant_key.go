package check

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByImageIDAndVariantKey 根据 ImageID 和 VariantKey 检查 ImageVariant 是否存在，实现 image_variants.Check 接口
func (h *Handler) ByImageIDAndVariantKey(ctx context.Context, imageID uuid.UUID, variantKey string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ImageVariant{}).
		Where("image_id = ? AND variant_key = ?", imageID, variantKey).
		Count(&count).Error
	return count > 0, err
}
