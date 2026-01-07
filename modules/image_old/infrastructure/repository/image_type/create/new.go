package create

import (
	"context"
	imageTypeDomain "nfxid/modules/image/domain/image_type"
	"nfxid/modules/image/infrastructure/repository/mapper"
)

// New 创建新的 ImageType，实现 imageTypeDomain.Create 接口
func (h *Handler) New(ctx context.Context, it *imageTypeDomain.ImageType) error {
	m := mapper.ImageTypeDomainToModel(it)
	return h.db.WithContext(ctx).Create(&m).Error
}
