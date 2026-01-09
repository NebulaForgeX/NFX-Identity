package create

import (
	"context"
	"nfxid/modules/tenants/domain/member_app_roles"
	"nfxid/modules/tenants/infrastructure/repository/member_app_roles/mapper"
)

// New 创建新的 MemberAppRole，实现 member_app_roles.Create 接口
func (h *Handler) New(ctx context.Context, mar *member_app_roles.MemberAppRole) error {
	m := mapper.MemberAppRoleDomainToModel(mar)
	return h.db.WithContext(ctx).Create(&m).Error
}
