package get

import (
	"context"
	"nfxid/modules/directory/domain/user_occupations"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_occupations/mapper"

	"github.com/google/uuid"
)

// CurrentByUserID 根据 UserID 获取当前职业，实现 user_occupations.Get 接口
func (h *Handler) CurrentByUserID(ctx context.Context, userID uuid.UUID) ([]*user_occupations.UserOccupation, error) {
	var ms []models.UserOccupation
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND is_current = ?", userID, true).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*user_occupations.UserOccupation, len(ms))
	for i, m := range ms {
		result[i] = mapper.UserOccupationModelToDomain(&m)
	}
	return result, nil
}
