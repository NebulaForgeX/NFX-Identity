package create

import (
	"context"
	"nfxid/modules/image/domain/images"
	"nfxid/modules/image/infrastructure/repository/images/mapper"
)

// New 创建新的 Image，实现 images.Create 接口
func (h *Handler) New(ctx context.Context, i *images.Image) error {
	m := mapper.ImageDomainToModel(i)
	return h.db.WithContext(ctx).Create(&m).Error
}
