package get

import (
	"context"
	education "nfxid/modules/auth/domain/profile_education"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
)

// ByProfileID 根据 ProfileID 获取 Education 列表，实现 education.Get 接口
func (h *Handler) ByProfileID(ctx context.Context, profileID uuid.UUID) ([]*education.Education, error) {
	var ms []models.Education
	if err := h.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Where("deleted_at IS NULL").
		Order("start_date DESC").
		Find(&ms).Error; err != nil {
		return nil, err
	}

	entities := make([]*education.Education, len(ms))
	for i := range ms {
		entities[i] = mapper.EducationModelToDomain(&ms[i])
	}
	return entities, nil
}
