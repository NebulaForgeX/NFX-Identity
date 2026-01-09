package delete

import (
	"context"
	"nfxid/modules/auth/domain/user_credentials"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 删除 UserCredential，实现 user_credentials.Delete 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&models.UserCredential{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return user_credentials.ErrUserCredentialNotFound
	}
	return nil
}
