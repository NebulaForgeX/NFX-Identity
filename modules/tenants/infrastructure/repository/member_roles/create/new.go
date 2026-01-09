package create

import (
	"context"
	"nfxid/modules/tenants/domain/member_roles"
	"nfxid/modules/tenants/infrastructure/repository/member_roles/mapper"
)

// New 创建新的 MemberRole，实现 member_roles.Create 接口
func (h *Handler) New(ctx context.Context, mr *member_roles.MemberRole) error {
	m := mapper.MemberRoleDomainToModel(mr)
	return h.db.WithContext(ctx).Create(&m).Error
}
