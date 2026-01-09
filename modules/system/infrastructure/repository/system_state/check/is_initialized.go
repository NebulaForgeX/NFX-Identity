package check

import (
	"context"
	"nfxid/modules/system/infrastructure/rdb/models"
)

// IsInitialized 检查系统是否已初始化，实现 system_state.Check 接口
func (h *Handler) IsInitialized(ctx context.Context) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.SystemState{}).
		Where("initialized = ?", true).
		Count(&count).Error
	return count > 0, err
}
