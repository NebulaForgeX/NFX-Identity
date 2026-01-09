package check

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"
)

// ByScope 根据 Scope 检查 Scope 是否存在，实现 scopes.Check 接口
func (h *Handler) ByScope(ctx context.Context, scope string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Scope{}).
		Where("scope = ?", scope).
		Count(&count).Error
	return count > 0, err
}
