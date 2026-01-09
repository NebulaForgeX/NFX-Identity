package get

import (
	"context"
	"nfxid/modules/directory/domain/user_educations"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_educations/mapper"

	"github.com/google/uuid"
)

// CurrentByUserID 根据 UserID 获取当前教育，实现 user_educations.Get 接口
func (h *Handler) CurrentByUserID(ctx context.Context, userID uuid.UUID) ([]*user_educations.UserEducation, error) {
	var ms []models.UserEducation
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND is_current = ?", userID, true).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*user_educations.UserEducation, len(ms))
	for i, m := range ms {
		result[i] = mapper.UserEducationModelToDomain(&m)
	}
	return result, nil
}
