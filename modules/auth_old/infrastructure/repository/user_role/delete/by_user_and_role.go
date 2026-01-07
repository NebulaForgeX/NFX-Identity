package delete

import (
	"context"
	userRoleErrors "nfxid/modules/auth/domain/user_role/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserAndRole 根据 UserID 和 RoleID 删除 UserRole，实现 user_role.Delete 接口
func (h *Handler) ByUserAndRole(ctx context.Context, userID, roleID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&models.UserRole{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userRoleErrors.ErrUserRoleNotFound
	}
	return nil
}
