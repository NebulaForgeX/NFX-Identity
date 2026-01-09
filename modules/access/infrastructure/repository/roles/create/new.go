package create

import (
	"context"
	"nfxid/modules/access/domain/roles"
	"nfxid/modules/access/infrastructure/repository/roles/mapper"
)

// New 创建新的 Role，实现 roles.Create 接口
func (h *Handler) New(ctx context.Context, r *roles.Role) error {
	m := mapper.RoleDomainToModel(r)
	return h.db.WithContext(ctx).Create(&m).Error
}
