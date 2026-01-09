package create

import (
	"context"
	"nfxid/modules/image/domain/image_types"
	"nfxid/modules/image/infrastructure/repository/image_types/mapper"
)

// New 创建新的 ImageType，实现 image_types.Create 接口
func (h *Handler) New(ctx context.Context, it *image_types.ImageType) error {
	m := mapper.ImageTypeDomainToModel(it)
	return h.db.WithContext(ctx).Create(&m).Error
}
