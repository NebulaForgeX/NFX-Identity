package get

import (
	"context"
	"nfxid/modules/auth/domain/password_history"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/password_history/mapper"

	"github.com/google/uuid"
)

// ByUserIDLimit 根据 UserID 获取 PasswordHistory 列表（限制数量），实现 password_history.Get 接口
func (h *Handler) ByUserIDLimit(ctx context.Context, userID uuid.UUID, limit int) ([]*password_history.PasswordHistory, error) {
	var ms []models.PasswordHistory
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*password_history.PasswordHistory, len(ms))
	for i, m := range ms {
		result[i] = mapper.PasswordHistoryModelToDomain(&m)
	}
	return result, nil
}
