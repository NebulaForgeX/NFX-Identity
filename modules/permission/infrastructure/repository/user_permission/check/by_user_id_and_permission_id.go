package check

import (
	"context"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndPermissionID 根据 UserID 和 PermissionID 检查 UserPermission 是否存在，实现 userPermissionDomain.Check 接口
func (h *Handler) ByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserPermission{}).
		Where("user_id = ? AND permission_id = ?", userID, permissionID).
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count > 0, err
}
