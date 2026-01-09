package create

import (
	"context"
	"nfxid/modules/directory/domain/users"
	"nfxid/modules/directory/infrastructure/repository/users/mapper"
)

// New 创建新的 User，实现 users.Create 接口
func (h *Handler) New(ctx context.Context, u *users.User) error {
	m := mapper.UserDomainToModel(u)
	return h.db.WithContext(ctx).Create(&m).Error
}
