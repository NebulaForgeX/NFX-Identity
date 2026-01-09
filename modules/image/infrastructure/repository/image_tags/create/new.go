package create

import (
	"context"
	"nfxid/modules/image/domain/image_tags"
	"nfxid/modules/image/infrastructure/repository/image_tags/mapper"
)

// New 创建新的 ImageTag，实现 image_tags.Create 接口
func (h *Handler) New(ctx context.Context, it *image_tags.ImageTag) error {
	m := mapper.ImageTagDomainToModel(it)
	return h.db.WithContext(ctx).Create(&m).Error
}
