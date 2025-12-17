package update

import (
	"context"
	"nfxid/modules/auth/domain/role"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// Generic 通用更新 Role，实现 role.Update 接口
func (h *Handler) Generic(ctx context.Context, r *role.Role) error {
	m := mapper.RoleDomainToModel(r)
	updates := mapper.RoleModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Role{}).
		Where("id = ?", r.ID()).
		Updates(updates).Error
}
