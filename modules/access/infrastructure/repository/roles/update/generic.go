package update

import (
	"context"
	"nfxid/modules/access/domain/roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/roles/mapper"
)

// Generic 通用更新 Role，实现 roles.Update 接口
func (h *Handler) Generic(ctx context.Context, r *roles.Role) error {
	m := mapper.RoleDomainToModel(r)
	updates := mapper.RoleModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Role{}).
		Where("id = ?", r.ID()).
		Updates(updates).Error
}
