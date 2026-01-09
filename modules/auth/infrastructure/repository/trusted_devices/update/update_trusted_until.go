package update

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// UpdateTrustedUntil 更新信任到期时间，实现 trusted_devices.Update 接口
func (h *Handler) UpdateTrustedUntil(ctx context.Context, deviceID string, trustedUntil time.Time) error {
	updates := map[string]any{
		models.TrustedDeviceCols.TrustedUntil: trustedUntil,
		models.TrustedDeviceCols.UpdatedAt:   time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.TrustedDevice{}).
		Where("device_id = ?", deviceID).
		Updates(updates).Error
}
