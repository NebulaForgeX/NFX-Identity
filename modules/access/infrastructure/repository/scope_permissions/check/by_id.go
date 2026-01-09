package check

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 检查 ScopePermission 是否存在，实现 scope_permissions.Check 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ScopePermission{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}
