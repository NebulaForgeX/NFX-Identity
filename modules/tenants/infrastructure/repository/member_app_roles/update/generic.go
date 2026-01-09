package update

import (
	"context"
	"nfxid/modules/tenants/domain/member_app_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_app_roles/mapper"
)

// Generic 通用更新 MemberAppRole，实现 member_app_roles.Update 接口
func (h *Handler) Generic(ctx context.Context, mar *member_app_roles.MemberAppRole) error {
	m := mapper.MemberAppRoleDomainToModel(mar)
	updates := mapper.MemberAppRoleModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.MemberAppRole{}).
		Where("id = ?", mar.ID()).
		Updates(updates).Error
}
