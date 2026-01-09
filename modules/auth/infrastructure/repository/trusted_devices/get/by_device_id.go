package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/trusted_devices"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/mapper"

	"gorm.io/gorm"
)

// ByDeviceID 根据 DeviceID 获取 TrustedDevice，实现 trusted_devices.Get 接口
func (h *Handler) ByDeviceID(ctx context.Context, deviceID string) (*trusted_devices.TrustedDevice, error) {
	var m models.TrustedDevice
	if err := h.db.WithContext(ctx).Where("device_id = ?", deviceID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, trusted_devices.ErrTrustedDeviceNotFound
		}
		return nil, err
	}
	return mapper.TrustedDeviceModelToDomain(&m), nil
}
