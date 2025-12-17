package delete

import (
	"context"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 User，实现 user.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.User{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userDomainErrors.ErrUserNotFound
	}
	return nil
}
