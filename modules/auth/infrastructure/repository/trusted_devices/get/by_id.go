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

// ByID 根据 ID 获取 TrustedDevice，实现 trusted_devices.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*trusted_devices.TrustedDevice, error) {
	var m models.TrustedDevice
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, trusted_devices.ErrTrustedDeviceNotFound
		}
		return nil, err
	}
	return mapper.TrustedDeviceModelToDomain(&m), nil
}
