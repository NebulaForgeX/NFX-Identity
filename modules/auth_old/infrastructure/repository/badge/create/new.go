package create

import (
	"context"
	"nfxid/modules/auth/domain/badge"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// New 创建新的 Badge，实现 badge.Create 接口
func (h *Handler) New(ctx context.Context, b *badge.Badge) error {
	m := mapper.BadgeDomainToModel(b)
	return h.db.WithContext(ctx).Create(&m).Error
}
