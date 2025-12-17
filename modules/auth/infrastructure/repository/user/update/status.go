package update

import (
	"context"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// Status 更新 User 的状态，实现 user.Update 接口
func (h *Handler) Status(ctx context.Context, userID uuid.UUID, status string) error {
	result := h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", userID).
		Update(models.UserCols.Status, status)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userDomainErrors.ErrUserNotFound
	}
	return nil
}
