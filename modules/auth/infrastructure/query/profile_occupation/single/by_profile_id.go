package single

import (
	"context"
	occupationDomainViews "nfxid/modules/auth/domain/profile_occupation/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
)

// ByProfileID 根据 ProfileID 获取 Occupation 列表，实现 occupationDomain.Single 接口
func (h *Handler) ByProfileID(ctx context.Context, profileID uuid.UUID) ([]*occupationDomainViews.OccupationView, error) {
	var items []models.Occupation
	if err := h.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Where("deleted_at IS NULL").
		Order("start_date DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	result := slice.MapP(items, mapper.OccupationModelToDomain)
	// Convert to pointers
	pointerResult := make([]*occupationDomainViews.OccupationView, len(result))
	for i := range result {
		pointerResult[i] = &result[i]
	}
	return pointerResult, nil
}
