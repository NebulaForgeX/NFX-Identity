package update

import (
	"context"
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/role_permissions/mapper"
)

// Generic 通用更新 RolePermission，实现 role_permissions.Update 接口
func (h *Handler) Generic(ctx context.Context, rp *role_permissions.RolePermission) error {
	m := mapper.RolePermissionDomainToModel(rp)
	updates := mapper.RolePermissionModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.RolePermission{}).
		Where("id = ?", rp.ID()).
		Updates(updates).Error
}
