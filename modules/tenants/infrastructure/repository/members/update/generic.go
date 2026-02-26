package update

import (
	"context"
	"nfxidentity/modules/tenants/domain/members"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/members/mapper"
)

// Generic 通用更新 Member，实现 members.Update 接口
func (h *Handler) Generic(ctx context.Context, m *members.Member) error {
	model := mapper.MemberDomainToModel(m)
	updates := mapper.MemberModelToUpdates(model)
	return h.db.WithContext(ctx).
		Model(&models.Member{}).
		Where("id = ?", m.ID()).
		Updates(updates).Error
}
