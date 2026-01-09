package update

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// UpdateLastUsed 更新最后使用时间，实现 trusted_devices.Update 接口
func (h *Handler) UpdateLastUsed(ctx context.Context, deviceID string) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.TrustedDeviceCols.LastUsedAt: now,
		models.TrustedDeviceCols.UpdatedAt:  now,
	}

	return h.db.WithContext(ctx).
		Model(&models.TrustedDevice{}).
		Where("device_id = ?", deviceID).
		Updates(updates).Error
}
