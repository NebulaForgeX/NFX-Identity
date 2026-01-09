package get

import (
	"context"
	"nfxid/modules/auth/domain/trusted_devices"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/mapper"

	"github.com/google/uuid"
)

// ByUserIDAndTenantID 根据 UserID 和 TenantID 获取 TrustedDevice 列表，实现 trusted_devices.Get 接口
func (h *Handler) ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) ([]*trusted_devices.TrustedDevice, error) {
	var ms []models.TrustedDevice
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND tenant_id = ?", userID, tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*trusted_devices.TrustedDevice, len(ms))
	for i, m := range ms {
		result[i] = mapper.TrustedDeviceModelToDomain(&m)
	}
	return result, nil
}
