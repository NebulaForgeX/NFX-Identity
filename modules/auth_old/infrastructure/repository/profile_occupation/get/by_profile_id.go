package get

import (
	"context"
	occupation "nfxid/modules/auth/domain/profile_occupation"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
)

// ByProfileID 根据 ProfileID 获取 Occupation 列表，实现 occupation.Get 接口
func (h *Handler) ByProfileID(ctx context.Context, profileID uuid.UUID) ([]*occupation.Occupation, error) {
	var ms []models.Occupation
	if err := h.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Where("deleted_at IS NULL").
		Order("start_date DESC").
		Find(&ms).Error; err != nil {
		return nil, err
	}

	entities := make([]*occupation.Occupation, len(ms))
	for i := range ms {
		entities[i] = mapper.OccupationModelToDomain(&ms[i])
	}
	return entities, nil
}
