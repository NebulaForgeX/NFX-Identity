package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 检查 UserCredential 是否存在，实现 user_credentials.Check 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserCredential{}).
		Where("id = ?", userID).
		Count(&count).Error
	return count > 0, err
}
