package create

import (
	"context"
	"nfxid/modules/auth/domain/user"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// New 创建新的 User，实现 user.Create 接口
func (h *Handler) New(ctx context.Context, u *user.User) error {
	m := mapper.UserDomainToModel(u)
	return h.db.WithContext(ctx).Create(&m).Error
}
