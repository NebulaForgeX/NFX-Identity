package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/trusted_devices"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserIDAndDeviceIDAndTenantID 根据 UserID、DeviceID 和 TenantID 获取 TrustedDevice，实现 trusted_devices.Get 接口
func (h *Handler) ByUserIDAndDeviceIDAndTenantID(ctx context.Context, userID uuid.UUID, deviceID string, tenantID uuid.UUID) (*trusted_devices.TrustedDevice, error) {
	var m models.TrustedDevice
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND device_id = ? AND tenant_id = ?", userID, deviceID, tenantID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, trusted_devices.ErrTrustedDeviceNotFound
		}
		return nil, err
	}
	return mapper.TrustedDeviceModelToDomain(&m), nil
}
