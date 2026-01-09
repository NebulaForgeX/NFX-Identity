package delete

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByDeviceID 根据 DeviceID 删除 TrustedDevice，实现 trusted_devices.Delete 接口
func (h *Handler) ByDeviceID(ctx context.Context, deviceID string) error {
	return h.db.WithContext(ctx).
		Where("device_id = ?", deviceID).
		Delete(&models.TrustedDevice{}).Error
}
