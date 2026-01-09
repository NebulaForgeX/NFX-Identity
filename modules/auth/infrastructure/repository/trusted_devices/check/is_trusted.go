package check

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// IsTrusted 检查设备是否被信任，实现 trusted_devices.Check 接口
func (h *Handler) IsTrusted(ctx context.Context, userID uuid.UUID, deviceID string, tenantID uuid.UUID) (bool, error) {
	var m models.TrustedDevice
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND device_id = ? AND tenant_id = ?", userID, deviceID, tenantID).
		First(&m).Error; err != nil {
		return false, err
	}

	// 检查信任时间是否过期
	return m.TrustedUntil.After(time.Now()), nil
}
