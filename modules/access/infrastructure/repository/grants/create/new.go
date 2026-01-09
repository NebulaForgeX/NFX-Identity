package create

import (
	"context"
	"nfxid/modules/access/domain/grants"
	"nfxid/modules/access/infrastructure/repository/grants/mapper"
)

// New 创建新的 Grant，实现 grants.Create 接口
func (h *Handler) New(ctx context.Context, g *grants.Grant) error {
	m := mapper.GrantDomainToModel(g)
	return h.db.WithContext(ctx).Create(&m).Error
}
