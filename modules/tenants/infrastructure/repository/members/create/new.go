package create

import (
	"context"
	"nfxid/modules/tenants/domain/members"
	"nfxid/modules/tenants/infrastructure/repository/members/mapper"
)

// New 创建新的 Member，实现 members.Create 接口
func (h *Handler) New(ctx context.Context, m *members.Member) error {
	model := mapper.MemberDomainToModel(m)
	return h.db.WithContext(ctx).Create(&model).Error
}
