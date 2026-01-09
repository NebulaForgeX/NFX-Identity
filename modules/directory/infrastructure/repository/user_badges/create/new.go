package create

import (
	"context"
	"nfxid/modules/directory/domain/user_badges"
	"nfxid/modules/directory/infrastructure/repository/user_badges/mapper"
)

// New 创建新的 UserBadge，实现 user_badges.Create 接口
func (h *Handler) New(ctx context.Context, ub *user_badges.UserBadge) error {
	m := mapper.UserBadgeDomainToModel(ub)
	return h.db.WithContext(ctx).Create(&m).Error
}
