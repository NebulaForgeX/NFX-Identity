package create

import (
	"context"
	"nfxid/modules/directory/domain/badges"
	"nfxid/modules/directory/infrastructure/repository/badges/mapper"
)

// New 创建新的 Badge，实现 badges.Create 接口
func (h *Handler) New(ctx context.Context, b *badges.Badge) error {
	m := mapper.BadgeDomainToModel(b)
	return h.db.WithContext(ctx).Create(&m).Error
}
