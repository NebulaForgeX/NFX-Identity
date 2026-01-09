package update

import (
	"context"
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/groups/mapper"
)

// Generic 通用更新 Group，实现 groups.Update 接口
func (h *Handler) Generic(ctx context.Context, g *groups.Group) error {
	m := mapper.GroupDomainToModel(g)
	updates := mapper.GroupModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Group{}).
		Where("id = ?", g.ID()).
		Updates(updates).Error
}
