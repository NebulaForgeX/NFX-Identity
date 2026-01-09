package update

import (
	"context"
	"nfxid/modules/auth/domain/trusted_devices"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/mapper"
)

// Generic 通用更新 TrustedDevice，实现 trusted_devices.Update 接口
func (h *Handler) Generic(ctx context.Context, td *trusted_devices.TrustedDevice) error {
	m := mapper.TrustedDeviceDomainToModel(td)
	updates := mapper.TrustedDeviceModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.TrustedDevice{}).
		Where("id = ?", td.ID()).
		Updates(updates).Error
}
