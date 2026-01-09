package update

import (
	"context"
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/mapper"
)

// Generic 通用更新 MemberGroup，实现 member_groups.Update 接口
func (h *Handler) Generic(ctx context.Context, mg *member_groups.MemberGroup) error {
	m := mapper.MemberGroupDomainToModel(mg)
	updates := mapper.MemberGroupModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.MemberGroup{}).
		Where("id = ?", mg.ID()).
		Updates(updates).Error
}
