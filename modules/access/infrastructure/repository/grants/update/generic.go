package update

import (
	"context"
	"nfxid/modules/access/domain/grants"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/grants/mapper"
)

// Generic 通用更新 Grant，实现 grants.Update 接口
func (h *Handler) Generic(ctx context.Context, g *grants.Grant) error {
	m := mapper.GrantDomainToModel(g)
	updates := mapper.GrantModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Grant{}).
		Where("id = ?", g.ID()).
		Updates(updates).Error
}
