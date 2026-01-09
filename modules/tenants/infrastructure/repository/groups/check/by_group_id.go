package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"
)

// ByGroupID 根据 GroupID 检查 Group 是否存在，实现 groups.Check 接口
func (h *Handler) ByGroupID(ctx context.Context, groupID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Group{}).
		Where("group_id = ?", groupID).
		Count(&count).Error
	return count > 0, err
}
