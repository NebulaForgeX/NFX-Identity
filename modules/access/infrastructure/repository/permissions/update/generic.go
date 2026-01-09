package update

import (
	"context"
	"nfxid/modules/access/domain/permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/permissions/mapper"
)

// Generic 通用更新 Permission，实现 permissions.Update 接口
func (h *Handler) Generic(ctx context.Context, p *permissions.Permission) error {
	m := mapper.PermissionDomainToModel(p)
	updates := mapper.PermissionModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Permission{}).
		Where("id = ?", p.ID()).
		Updates(updates).Error
}
