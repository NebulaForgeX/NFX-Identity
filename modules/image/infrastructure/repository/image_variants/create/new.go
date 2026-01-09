package create

import (
	"context"
	"nfxid/modules/image/domain/image_variants"
	"nfxid/modules/image/infrastructure/repository/image_variants/mapper"
)

// New 创建新的 ImageVariant，实现 image_variants.Create 接口
func (h *Handler) New(ctx context.Context, iv *image_variants.ImageVariant) error {
	m := mapper.ImageVariantDomainToModel(iv)
	return h.db.WithContext(ctx).Create(&m).Error
}
