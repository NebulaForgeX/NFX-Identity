package single

import (
	"context"
	userDomainViews "nfxid/modules/auth/domain/user/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/views"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
)

// UserBadges 根据 UserID 获取用户徽章，实现 profileBadgeDomain.Single 接口
func (h *Handler) UserBadges(ctx context.Context, userID uuid.UUID) ([]*userDomainViews.UserBadgesView, error) {
	var items []views.UserBadgesView
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&items).Error; err != nil {
		return nil, err
	}
	result := slice.MapP(items, mapper.UserBadgesViewToDomain)
	// Convert to pointers
	pointerResult := make([]*userDomainViews.UserBadgesView, len(result))
	for i := range result {
		pointerResult[i] = &result[i]
	}
	return pointerResult, nil
}
