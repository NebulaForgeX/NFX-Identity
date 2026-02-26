package update

import (
	"context"
	"nfxidentity/modules/tenants/domain/member_roles"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/member_roles/mapper"
)

// Generic 通用更新 MemberRole，实现 member_roles.Update 接口
func (h *Handler) Generic(ctx context.Context, mr *member_roles.MemberRole) error {
	m := mapper.MemberRoleDomainToModel(mr)
	updates := mapper.MemberRoleModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.MemberRole{}).
		Where("id = ?", mr.ID()).
		Updates(updates).Error
}
