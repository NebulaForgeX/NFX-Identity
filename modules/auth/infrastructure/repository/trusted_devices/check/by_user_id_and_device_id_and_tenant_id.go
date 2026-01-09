package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndDeviceIDAndTenantID 根据 UserID、DeviceID 和 TenantID 检查 TrustedDevice 是否存在，实现 trusted_devices.Check 接口
func (h *Handler) ByUserIDAndDeviceIDAndTenantID(ctx context.Context, userID uuid.UUID, deviceID string, tenantID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.TrustedDevice{}).
		Where("user_id = ? AND device_id = ? AND tenant_id = ?", userID, deviceID, tenantID).
		Count(&count).Error
	return count > 0, err
}
