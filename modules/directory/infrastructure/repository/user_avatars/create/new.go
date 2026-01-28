package create

import (
	"context"
	"nfxid/modules/directory/domain/user_avatars"
	"nfxid/modules/directory/infrastructure/repository/user_avatars/mapper"
)

// New 创建新的 UserAvatar，实现 user_avatars.Create 接口
func (h *Handler) New(ctx context.Context, ua *user_avatars.UserAvatar) error {
	m := mapper.UserAvatarDomainToModel(ua)
	return h.db.WithContext(ctx).Create(&m).Error
}
