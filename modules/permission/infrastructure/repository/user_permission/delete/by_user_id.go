package delete

import (
	"context"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 删除 UserPermission 列表，实现 userPermissionDomain.Delete 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&models.UserPermission{})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
