package get

import (
	"context"
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/password_resets/mapper"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 获取 PasswordReset 列表，实现 password_resets.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) ([]*password_resets.PasswordReset, error) {
	var ms []models.PasswordReset
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*password_resets.PasswordReset, len(ms))
	for i, m := range ms {
		result[i] = mapper.PasswordResetModelToDomain(&m)
	}
	return result, nil
}
