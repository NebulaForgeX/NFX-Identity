package create

import (
	"context"
	"nfxid/modules/auth/domain/user_role"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// New 创建新的 UserRole，实现 user_role.Create 接口
func (h *Handler) New(ctx context.Context, ur *user_role.UserRole) error {
	m := mapper.UserRoleDomainToModel(ur)
	return h.db.WithContext(ctx).Create(&m).Error
}
