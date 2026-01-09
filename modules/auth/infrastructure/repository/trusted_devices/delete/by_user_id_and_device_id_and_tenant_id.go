package delete

import (
	"context"
	"nfxid/modules/auth/domain/trusted_devices"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndDeviceIDAndTenantID 根据 UserID、DeviceID 和 TenantID 删除 TrustedDevice，实现 trusted_devices.Delete 接口
func (h *Handler) ByUserIDAndDeviceIDAndTenantID(ctx context.Context, userID uuid.UUID, deviceID string, tenantID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("user_id = ? AND device_id = ? AND tenant_id = ?", userID, deviceID, tenantID).
		Delete(&models.TrustedDevice{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return trusted_devices.ErrTrustedDeviceNotFound
	}
	return nil
}
