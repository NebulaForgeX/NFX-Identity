package get

import (
	"context"
	"nfxid/modules/auth/domain/login_attempts"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/mapper"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 获取 LoginAttempt 列表，实现 login_attempts.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) ([]*login_attempts.LoginAttempt, error) {
	var ms []models.LoginAttempt
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*login_attempts.LoginAttempt, len(ms))
	for i, m := range ms {
		result[i] = mapper.LoginAttemptModelToDomain(&m)
	}
	return result, nil
}
