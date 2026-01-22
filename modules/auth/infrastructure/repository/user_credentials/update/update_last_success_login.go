package update

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// UpdateLastSuccessLogin 更新最后成功登录时间，实现 user_credentials.Update 接口
func (h *Handler) UpdateLastSuccessLogin(ctx context.Context, userID uuid.UUID) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.UserCredentialCols.LastSuccessLoginAt: &now,
		models.UserCredentialCols.UpdatedAt:          now,
	}

	return h.db.WithContext(ctx).
		Model(&models.UserCredential{}).
		Where("id = ?", userID).
		Updates(updates).Error
}
