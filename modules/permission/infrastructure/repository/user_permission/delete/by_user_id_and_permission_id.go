package delete

import (
	"context"
	userPermissionDomainErrors "nfxid/modules/permission/domain/user_permission/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndPermissionID 根据 UserID 和 PermissionID 删除 UserPermission，实现 userPermissionDomain.Delete 接口
func (h *Handler) ByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("user_id = ? AND permission_id = ?", userID, permissionID).
		Delete(&models.UserPermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userPermissionDomainErrors.ErrUserPermissionNotFound
	}
	return nil
}
