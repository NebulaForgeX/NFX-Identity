package delete

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByImageIDAndVariantKey 根据 ImageID 和 VariantKey 删除 ImageVariant，实现 image_variants.Delete 接口
func (h *Handler) ByImageIDAndVariantKey(ctx context.Context, imageID uuid.UUID, variantKey string) error {
	return h.db.WithContext(ctx).
		Where("image_id = ? AND variant_key = ?", imageID, variantKey).
		Delete(&models.ImageVariant{}).Error
}
