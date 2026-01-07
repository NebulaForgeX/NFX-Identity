package create

import (
	"context"
	profileBadge "nfxid/modules/auth/domain/profile_badge"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// New 创建新的 ProfileBadge，实现 profile_badge.Create 接口
func (h *Handler) New(ctx context.Context, pb *profileBadge.ProfileBadge) error {
	m := mapper.ProfileBadgeDomainToModel(pb)
	return h.db.WithContext(ctx).Create(&m).Error
}
