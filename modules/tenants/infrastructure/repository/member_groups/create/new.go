package create

import (
	"context"
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/mapper"
)

// New 创建新的 MemberGroup，实现 member_groups.Create 接口
func (h *Handler) New(ctx context.Context, mg *member_groups.MemberGroup) error {
	m := mapper.MemberGroupDomainToModel(mg)
	return h.db.WithContext(ctx).Create(&m).Error
}
