package check

import (
	"context"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// ByAppID 根据 AppID 检查 App 是否存在，实现 apps.Check 接口
func (h *Handler) ByAppID(ctx context.Context, appID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Application{}).
		Where("app_id = ?", appID).
		Count(&count).Error
	return count > 0, err
}
