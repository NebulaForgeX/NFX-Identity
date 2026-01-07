package create

import (
	"context"
	imageDomain "nfxid/modules/image/domain/image"
	"nfxid/modules/image/infrastructure/repository/mapper"
)

// New 创建新的 Image，实现 imageDomain.Create 接口
func (h *Handler) New(ctx context.Context, img *imageDomain.Image) error {
	m := mapper.ImageDomainToModel(img)
	return h.db.WithContext(ctx).Create(&m).Error
}
