package create

import (
	"context"
	"nfxid/modules/directory/domain/user_profiles"
	"nfxid/modules/directory/infrastructure/repository/user_profiles/mapper"
)

// New 创建新的 UserProfile，实现 user_profiles.Create 接口
func (h *Handler) New(ctx context.Context, up *user_profiles.UserProfile) error {
	m := mapper.UserProfileDomainToModel(up)
	return h.db.WithContext(ctx).Create(&m).Error
}
