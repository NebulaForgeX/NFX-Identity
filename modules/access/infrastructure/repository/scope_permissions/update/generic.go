package update

import (
	"context"
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/mapper"
)

// Generic 通用更新 ScopePermission，实现 scope_permissions.Update 接口
func (h *Handler) Generic(ctx context.Context, sp *scope_permissions.ScopePermission) error {
	m := mapper.ScopePermissionDomainToModel(sp)
	updates := mapper.ScopePermissionModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ScopePermission{}).
		Where("id = ?", sp.ID()).
		Updates(updates).Error
}
