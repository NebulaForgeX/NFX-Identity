package create

import (
	"context"
	"nfxid/modules/auth/domain/role"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// New 创建新的 Role，实现 role.Create 接口
func (h *Handler) New(ctx context.Context, r *role.Role) error {
	m := mapper.RoleDomainToModel(r)
	return h.db.WithContext(ctx).Create(&m).Error
}
