package check

import (
	"context"
	"nfxid/modules/permission/infrastructure/rdb/models"
)

// ByTag 根据 Tag 检查 Permission 是否存在，实现 permissionDomain.Check 接口
func (h *Handler) ByTag(ctx context.Context, tag string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Permission{}).
		Where("tag = ?", tag).
		Count(&count).Error
	return count > 0, err
}
