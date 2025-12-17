package update

import (
	"context"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// Password 更新 User 的密码，实现 user.Update 接口
func (h *Handler) Password(ctx context.Context, userID uuid.UUID, hashedPassword string) error {
	result := h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", userID).
		Update(models.UserCols.PasswordHash, hashedPassword)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userDomainErrors.ErrUserNotFound
	}
	return nil
}
