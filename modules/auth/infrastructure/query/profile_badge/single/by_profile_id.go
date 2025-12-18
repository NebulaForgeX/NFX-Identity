package single

import (
	"context"
	profileBadgeDomainViews "nfxid/modules/auth/domain/profile_badge/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
)

// ByProfileID 根据 ProfileID 获取 ProfileBadge 列表，实现 profileBadgeDomain.Single 接口
func (h *Handler) ByProfileID(ctx context.Context, profileID uuid.UUID) ([]*profileBadgeDomainViews.ProfileBadgeView, error) {
	var items []models.ProfileBadge
	if err := h.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Order("earned_at DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	result := slice.MapP(items, mapper.ProfileBadgeModelToDomain)
	// Convert to pointers
	pointerResult := make([]*profileBadgeDomainViews.ProfileBadgeView, len(result))
	for i := range result {
		pointerResult[i] = &result[i]
	}
	return pointerResult, nil
}
