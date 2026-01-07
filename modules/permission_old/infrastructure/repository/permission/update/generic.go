package update

import (
	"context"
	permissionDomain "nfxid/modules/permission/domain/permission"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"
)

// Generic 更新 Permission，实现 permissionDomain.Update 接口
func (h *Handler) Generic(ctx context.Context, p *permissionDomain.Permission) error {
	m := mapper.PermissionDomainToModel(p)
	updates := mapper.PermissionModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Permission{}).
		Where("id = ?", p.ID()).
		Updates(updates).Error
}
